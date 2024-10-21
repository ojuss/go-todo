package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main () {
	fmt.Println("Hello Ojus")
	
	app := fiber.New()

	app.Get("/book", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello dumb"})
	})

	log.Fatal(app.Listen(":4000"))
}