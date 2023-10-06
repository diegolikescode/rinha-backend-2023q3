package main

import (
	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    var Database = config.PostgresConnection()

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "workin",
        })
    })
    r.POST("/pessoa", func (c *gin.Context) {
        handlers.CreatePessoa(c, Database)
    })
    r.GET("/pessoa/:id", handlers.BuscaPessoa)

    r.Run(":6969")
}

