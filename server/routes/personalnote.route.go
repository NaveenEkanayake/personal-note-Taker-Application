package routes

import (
	"personal-note-taker/controllers"
	"github.com/gofiber/fiber/v2"
)

func PersonalNoteRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/getnotes", controllers.GetNotes)
	api.Get("/getnote/:id", controllers.GetNoteByID)
	api.Post("/addnote", controllers.AddNote)
	api.Put("/updatenote/:id", controllers.UpdateNoteByID)
	api.Delete("/deletenote/:id", controllers.DeleteNoteByID)
}
