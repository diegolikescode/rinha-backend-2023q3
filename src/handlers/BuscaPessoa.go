package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"rinha-backend-2023q3/src/entities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuscaPessoa(c *gin.Context, db *gorm.DB) {
	userID := c.Param("id")
	fmt.Println(userID)
	var user entities.Pessoa
	res := db.First(&user, "id = ?", userID)
	fmt.Println(res)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, user)
}

func BuscaPessoaPorTermo(c *gin.Context, db *gorm.DB) {
	searchTerm := c.Query("t")
	fmt.Println(searchTerm)

	if searchTerm == "" {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userTerm := "%" + searchTerm + "%"

	var users []entities.Pessoa
	db.Where("search_string LIKE ? LIMIT 50", userTerm).Find(&users)

	if len(users) == 0 {
		c.JSON(http.StatusOK, []interface{}{})
		return
	}

	var usersReturn []entities.ReturnPessoa
	for _, user := range users {
		usersReturn = append(usersReturn, entities.ReturnPessoa{
			Id:         user.Id,
			Apelido:    user.Apelido,
			Nome:       user.Nome,
			Nascimento: user.Nascimento,
			Stack:      strings.Split(user.Stack, ";"),
		})
	}

	c.IndentedJSON(http.StatusOK, usersReturn)
}

func ContaPessoas(c *gin.Context, db *gorm.DB) {
	var count int64
	db.Model(&entities.Pessoa{}).Count(&count)
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, strconv.FormatInt(count, 10))
}
