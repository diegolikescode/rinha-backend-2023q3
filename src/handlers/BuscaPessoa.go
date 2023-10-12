package handlers

import (
	"net/http"
	"rinha-backend-2023q3/src/entities"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuscaPessoa (c *gin.Context, db *gorm.DB) {
	userID := c.Param("id")

	var user entities.Pessoa
	if db.Where("id = ?", userID).First(&user).RowsAffected == 0 {	
		c.IndentedJSON(http.StatusOK, entities.HttpResponse{
			Message: "usuario nao encontrado",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
	return 
}

func BuscaPessoaPorTermo (c *gin.Context, db *gorm.DB) {
	searchTerm := c.Query("t")
	userTerm := "%" + searchTerm + "%"


	var users []entities.Pessoa
	if db.Where(
		"apelido LIKE ? OR nome LIKE ? OR stack LIKE ?", 
		userTerm, userTerm, userTerm).Find(&users).RowsAffected == 0 {
		c.IndentedJSON(http.StatusOK, entities.HttpResponse{
			Message: "usuario nao encontrado",
		})
		return
	}

	c.IndentedJSON(http.StatusOK,  users)
	return
}

func ContaPessoas (c *gin.Context, db *gorm.DB) {
	var count int64
	db.Model(&entities.Pessoa{}).Count(&count)
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, strconv.FormatInt(count, 10))
	return 
}

