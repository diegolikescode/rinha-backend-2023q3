package handlers

import (
	"net/http"
	"rinha-backend-2023q3/src/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePessoa(c *gin.Context, db *gorm.DB) {
    var pessoaBody entities.Pessoa
    c.ShouldBindJSON(&pessoaBody)

    c.IndentedJSON(http.StatusOK, pessoaBody)
    // MAKE THE VALIDATIONS
    // CREATE THE DATA

    return 
}

