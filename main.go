package main

import (
	"log"
	"os"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/handlers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.PostgresConnection()

	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		content := []byte("<html><head><title>XAMA</title></head><body><h1>JUSTIN CASE</h1></body></html>")
		c.Set("content-type", "text/html; charset=utf-8")

		c.Status(200).Set("success", "true")
		return c.Send(content)
	})

	app.Post("/pessoas", handlers.CreatePessoa)
	app.Get("/pessoas/:id", handlers.BuscaPessoa)
	app.Get("/pessoas", handlers.BuscaPessoaPorTermo)
	app.Get("/contagem-pessoas", handlers.ContaPessoas)

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "4200")
	}
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
