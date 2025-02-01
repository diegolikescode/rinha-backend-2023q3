package main

import (
	"os"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	Database := config.PostgresConnection()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		htmlContent := "<html><head><title>just testin</title></head><body><h1>JUSTIN CASE</h1></body></html>"
		c.Data(200, "text/html; charset=utf-8", []byte(htmlContent))
	})

	r.POST("/pessoas", func(c *gin.Context) {
		handlers.CreatePessoa(c, Database)
	})

	r.GET("/pessoas/:id", func(c *gin.Context) {
		handlers.BuscaPessoa(c, Database)
	})

	r.GET("/pessoas", func(c *gin.Context) {
		handlers.BuscaPessoaPorTermo(c, Database)
	})

	r.GET("/contagem-pessoas", func(c *gin.Context) {
		handlers.ContaPessoas(c, Database)
	})

	r.Run(":" + os.Getenv("PORT"))
}
