package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"personal-note-taker/controllers"
	"personal-note-taker/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	PORT := os.Getenv("PORT")
	MONGO_DB_URL := os.Getenv("MONGO_DB_URL")

	clientOptions := options.Client().ApplyURI(MONGO_DB_URL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	controllers.Collection = client.Database("PersonalNoteTaker").Collection("personalnote")
	fmt.Printf("Connected to MongoDB on port %s\n", PORT)

	app := fiber.New()
	routes.PersonalNoteRoutes(app)
	routes.PersonalauthRoutes(app)

	// Start the server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
