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

	var user entities.ReturnPessoa
	if db.Where("id = ?", userID).First(&user).RowsAffected == 0 {	
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
	return 
}

func BuscaPessoaPorTermo (c *gin.Context, db *gorm.DB) {
	searchTerm := c.Query("t")
	if searchTerm == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}

	userTerm := "%" + searchTerm + "%"

	var users []entities.Pessoa
	db.Where("admin LIKE ? LIMIT 50", userTerm).Find(&users)

	if users == nil {
		c.IndentedJSON(http.StatusOK, []entities.ReturnPessoa{})
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

