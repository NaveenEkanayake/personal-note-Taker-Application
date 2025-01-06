package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv" // Make sure to import godotenv
)

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Slice to hold notes
var personalNotes = []*Note{}

func main() {

	app := fiber.New()

	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the PORT from the .env file
	PORT := os.Getenv("PORT")

	// If PORT is not set, use the default port
	if PORT == "" {
		PORT = "3000" // default to 3000 if not found
	}

	// Define a route for GET request
	app.Get("/api/getall", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"mesg":          "Personal notes Retrieval Successful",
			"personalNotes": personalNotes,
		})
	})

	app.Post("/api/addnote", func(c *fiber.Ctx) error {
		personalNote := &Note{} // Create a pointer to the Note struct
		if err := c.BodyParser(personalNote); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		if personalNote.Title == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Personal Note Title is required",
			})
		}
		if personalNote.Content == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Personal Note Content is required",
			})
		}
		personalNote.ID = len(personalNotes) + 1
		personalNotes = append(personalNotes, personalNote)
		return c.Status(201).JSON(fiber.Map{
			"msg":          "Personal Note Added Successfully",
			"personalNote": personalNote,
		})
	})

	// Update an existing note
	app.Put("/api/updateNote/:id", func(c *fiber.Ctx) error {
		// Get the ID from the URL parameters
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid note ID",
			})
		}

		// Find the note by ID
		for i, personalNote := range personalNotes {
			if personalNote.ID == id {
				// Parse the updated data
				updatedNote := &Note{}
				if err := c.BodyParser(updatedNote); err != nil {
					return c.Status(400).JSON(fiber.Map{
						"error": "Invalid request body",
					})
				}

				// Validate updated data
				if updatedNote.Title == "" {
					return c.Status(400).JSON(fiber.Map{
						"error": "Title is required",
					})
				}
				if updatedNote.Content == "" {
					return c.Status(400).JSON(fiber.Map{
						"error": "Content is required",
					})
				}

				// Update the note
				personalNotes[i].Title = updatedNote.Title
				personalNotes[i].Content = updatedNote.Content

				return c.Status(200).JSON(fiber.Map{
					"msg":          "Personal Note Updated Successfully",
					"personalNote": personalNotes[i],
				})
			}
		}

		// If note not found
		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	})

	// Delete an existing note
	app.Delete("/api/DeleteNote/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		// Loop through the notes to find the matching ID
		for i, personalNote := range personalNotes {
			// Convert the string id to an integer for proper comparison
			if fmt.Sprintf("%d", personalNote.ID) == id {
				// Save the deleted note for the response
				deletedNote := personalNotes[i]

				// Remove the note by appending slices
				personalNotes = append(personalNotes[:i], personalNotes[i+1:]...)

				// Return a success message with the deleted note
				return c.Status(200).JSON(fiber.Map{
					"msg":         "Personal Note Deleted Successfully",
					"deletedNote": deletedNote,
				})
			}
		}

		// If the note is not found, return an error
		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	})

	// Start the app on the port specified in the .env file
	log.Fatal(app.Listen(":" + PORT))
}
