package main

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	uuid.EnableRandPool()

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	app := fx.New(
		config.Module
	)

	app.Run()
}

// func main() {
// config.PostgresConnection()
// cache.InstanciateCacheSingleton()

// http.Handle("/metrics", promhttp.Handler())
//
// app := fiber.New(fiber.Config{ReadTimeout: 20 * time.Second, WriteTimeout: 20 * time.Second})
// app.Get("/", func(c fiber.Ctx) error {
// 	content := []byte("<html><head><title>XAMA</title></head><body><h1>JUSTIN CASE</h1></body></html>")
// 	c.Set("content-type", "text/html; charset=utf-8")
//
// 	return c.Send(content)
// })
//
// app.Post("/pessoas", handlers.CreatePessoa)
// app.Get("/pessoas/:id", handlers.BuscaPessoa)
// app.Get("/pessoas", handlers.BuscaPessoaPorTermo)
// app.Get("/contagem-pessoas", handlers.ContaPessoas)
//
// if os.Getenv("PORT") == "" {
// 	os.Setenv("PORT", "4200")
// }
// log.Fatal(app.Listen(":" + os.Getenv("PORT")))
// }
