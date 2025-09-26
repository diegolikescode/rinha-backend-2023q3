package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/handlers"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.PostgresConnection()

	http.Handle("/metrics", promhttp.Handler())

	app := fiber.New(fiber.Config{ReadTimeout: 20 * time.Second, WriteTimeout: 20 * time.Second})
	app.Get("/", func(c fiber.Ctx) error {
		content := []byte("<html><head><title>XAMA</title></head><body><h1>JUSTIN CASE</h1></body></html>")
		c.Set("content-type", "text/html; charset=utf-8")

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
