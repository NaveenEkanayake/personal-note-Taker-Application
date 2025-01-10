package controllers

import (
	"context"
	"personal-note-taker/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

func GetNotes(c *fiber.Ctx) error {
	var personalNotes []models.PersonalNote

	// Find all notes in the collection
	cursor, err := Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving notes",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor to decode the notes
	for cursor.Next(context.Background()) {
		var personalNote models.PersonalNote
		if err := cursor.Decode(&personalNote); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error decoding note",
				"error":   err.Error(),
			})
		}
		personalNotes = append(personalNotes, personalNote)
	}

	if len(personalNotes) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No personal notes exist",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":       "Notes retrieved successfully!",
		"retrievedData": personalNotes,
	})
}

func GetNoteByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	filter := bson.M{"_id": objectID}

	var personalNote models.PersonalNote

	err = Collection.FindOne(context.Background(), filter).Decode(&personalNote)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Note not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Error retrieving the note",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Note retrieved successfully!",
		"note":    personalNote,
	})
}

func AddNote(c *fiber.Ctx) error {
	personalNote := new(models.PersonalNote)
	if err := c.BodyParser(personalNote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if personalNote.Title == "" || personalNote.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "All fields are required",
		})
	}

	insertResult, err := Collection.InsertOne(context.Background(), personalNote)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to insert note into database",
			"message": err.Error(),
		})
	}

	personalNote.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Note added successfully",
		"note":    personalNote,
	})
}

func UpdateNoteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	updatedData := new(models.PersonalNote)
	if err := c.BodyParser(updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"title":   updatedData.Title,
			"content": updatedData.Content,
		},
	}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update the note",
			"message": err.Error(),
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No note found with the given ID",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Note updated successfully",
		"updatingdata": update,
	})
}

func DeleteNoteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid note ID",
		})
	}

	filter := bson.M{"_id": objectID}

	result, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete the note",
			"message": err.Error(),
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No note found with the given ID",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Note deleted successfully",
	})
}
