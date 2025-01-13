package controllers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Fullname  string             `json:"fullname"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// Claims struct for JWT
type Claims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	UserRole string `json:"role"`
	jwt.StandardClaims
}

var (
	JWT_SECRET_KEY string
	userCollection *mongo.Collection
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	if JWT_SECRET_KEY == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the environment")
	}

	mongoURI := os.Getenv("MONGO_DB_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_DB_URL is not set in the environment")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	userCollection = client.Database("PersonalNoteTaker").Collection("users")
}

// Signup handler
func Signup(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please provide valid fullname, email, and password",
		})
	}

	var existingUser User
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists with this email",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error hashing password",
		})
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully!",
		"email":   user.Email,
	})
}

// Login handler// Login handler
func LoginUser(c *fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var user User
	err := userCollection.FindOne(context.Background(), bson.M{"email": credentials.Email}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := generateJWT(user.ID.Hex(), user.Email, user.Fullname, "user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set the token in a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,                           // Ensures the cookie is not accessible via JavaScript
		Expires:  time.Now().Add(24 * time.Hour), // 1 day expiration
		Path:     "/",                            // Available to the entire site
	})

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// Generate JWT function
func generateJWT(userID string, email string, fullname string, userRole string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Email:    email,
		Fullname: fullname,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET_KEY))
}

func VerifyJWT(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token") // Read the token from the cookie
	if tokenStr == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No Token Found",
		})
	}

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	// Now fetch user info from the DB based on the claims, but exclude password.
	var user User
	objID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	// Fetch the user but ensure the password is not included
	err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Ensure password is never included in the response
	user.Password = ""

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// GetUser handler
func GetUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Ensure password is never included in response
	user.Password = "" // This line can be omitted entirely since password isn't included in the JWT

	return c.JSON(fiber.Map{
		"user": user,
	})
}

// Logout handler
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).SendString("Logout successful!")
}
