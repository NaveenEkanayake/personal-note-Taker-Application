package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello there1 !!!!")
	// Create a new Fiber app
	app := fiber.New()
	// Start the app on port 3000 and log errors if any
	log.Fatal(app.Listen(":3000"))
}
