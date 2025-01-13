package routes

import (
	"personal-note-taker/controllers"

	"github.com/gofiber/fiber/v2"
)

func PersonalauthRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/signup", controllers.Signup)
	api.Post("/login", controllers.LoginUser)
	api.Get("/verify-token", controllers.VerifyJWT, controllers.GetUser)
	api.Post("/logout", controllers.Logout)
}
