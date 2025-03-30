package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	srv := fiber.New()
	srv.Post("api/v1/contact", NewContact)
	srv.Get("api/v1/contact", ContactRead)
	srv.Put("api/v1/contact", ContactChange)
	srv.Delete("api/v1/contact", ContactDelete)
	srv.Post("/api/v1/group", NewGroup)
	srv.Get("/api/v1/group", GroupRead)
	srv.Put("/api/v1/group", GroupChange)
	srv.Delete("/api/v1/group", GroupDelete)
	// Порт сервера
	if err := srv.Listen(":6080"); err != nil {
		log.Fatal(err)
	}
}
