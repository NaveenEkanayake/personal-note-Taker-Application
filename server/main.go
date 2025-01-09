package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type personalnote struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}

var collection *mongo.Collection

func main() {
	app := fiber.New()

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	PORT := os.Getenv("PORT")
	MONGO_DB_URL := os.Getenv("MONGO_DB_URL")

	// MongoDB client setup
	clientOptions := options.Client().ApplyURI(MONGO_DB_URL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to confirm connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	// Accessing the MongoDB collection
	collection = client.Database("Personal-note_Taker").Collection("personalnote")
	fmt.Printf("Connected to MongoDB on port %s\n", PORT)

	// Define routes
	app.Get("/api/getnote", GetNote)
	app.Get("/api/getnote/:id", GetNoteByID)
	app.Post("/api/addnote", AddNote)
	app.Put("/api/updatenote/:id", updateNote)
	app.Delete("/api/deletenote/:id", deleteNote)

	// Start the Fiber app
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}

func GetNote(c *fiber.Ctx) error {
	var personalnotes []personalnote

	// Find all notes in the collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving notes",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(context.Background()) // Ensure cursor is closed after iteration

	// Iterate over the cursor to decode the notes
	for cursor.Next(context.Background()) {
		var personalnote personalnote
		if err := cursor.Decode(&personalnote); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error decoding note",
				"error":   err.Error(),
			})
		}
		personalnotes = append(personalnotes, personalnote)
	}

	// Check if no personal notes exist
	if len(personalnotes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No personal notes exist",
		})
	}

	// Return the retrieved notes
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Notes retrieved successfully!",
		"retrievedData": personalnotes,
	})
}

func GetNoteByID(c *fiber.Ctx) error {
    // Get the ID from the request parameters
    id := c.Params("id")

    // Convert the ID string to a MongoDB ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid note ID",
        })
    }

    // Define the filter to search for the note
    filter := bson.M{"_id": objectID}

    // Create a variable to hold the result
    var personalnote personalnote

    // Query the MongoDB collection
    err = collection.FindOne(context.Background(), filter).Decode(&personalnote)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "message": "Note not found",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error retrieving the note",
            "details": err.Error(),
        })
    }

    // Return the found note
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Note retrieved successfully!",
        "note":    personalnote,
    })
}


func AddNote(c *fiber.Ctx) error {
	// Parse the request body into the personalnote struct
	personalnote := new(personalnote)
	if err := c.BodyParser(personalnote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Validate required fields
	if personalnote.Title == "" || personalnote.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "All fields are required",
		})
	}

	// Insert the personalnote into the MongoDB collection
	insertResult, err := collection.InsertOne(context.Background(), personalnote)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert note into database",
			"message": err.Error(),
		})
	}

	// Assign the inserted ID back to the personalnote struct
	personalnote.ID = insertResult.InsertedID.(primitive.ObjectID)

	// Return the newly created note
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note added successfully",
		"note":    personalnote,
	})
}

func updateNote(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid personal note ID",
		})
	}

	// Parse the request body for the update data
	updatedData := new(personalnote)
	if err := c.BodyParser(updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Create the filter and update document
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title":   updatedData.Title,
			"content": updatedData.Content,
		},
	}

	// Perform the update operation
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update the personal note",
			"message": err.Error(),
		})
	}

	// Check if any document was modified
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No note found with the given ID",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Note updated successfully",
		"updatingdata": update,
	})
}

func deleteNote(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid personal note ID",
		})
	}

	// Create the filter to match the document by ID
	filter := bson.M{"_id": objectID}

	// Attempt to delete the document
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete the personal note",
			"message": err.Error(),
		})
	}

	// Check if a document was deleted
	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No note found with the given ID",
		})
	}

	// Return success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Note deleted successfully",
	})
}
