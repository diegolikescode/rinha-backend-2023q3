package main

import (
	"log"
	"os"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/handlers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	Database := config.PostgresConnection()

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		content := []byte("<html><head><title>XAMA</title></head><body><h1>JUSTIN CASE</h1></body></html>")
		c.Set("content-type", "text/html; charset=utf-8")

		c.Status(200).Set("success", "true")
		return c.Send(content)
	})

	app.Post("/pessoas", func(c fiber.Ctx) error {
		return handlers.CreatePessoa(c, Database)
	})

	app.Get("/pessoas/:id", func(c fiber.Ctx) error {
		return handlers.BuscaPessoa(c, Database)
	})

	app.Get("/pessoas", func(c fiber.Ctx) error {
		return handlers.BuscaPessoaPorTermo(c, Database)
	})

	app.Get("/contagem-pessoas", func(c fiber.Ctx) error {
		return handlers.ContaPessoas(c, Database)
	})

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "4200")
	}
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
