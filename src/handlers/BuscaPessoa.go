package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"rinha-backend-2023q3/src/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func BuscaPessoa(c fiber.Ctx, db *gorm.DB) error {
	userID := c.Params("id")
	var user entities.Pessoa
	res := db.First(&user, "id = ?", userID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return nil
		}
	}

	userBytes, err := json.Marshal(user)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString("Something weird happened while trying to Marshal the user on BuscaPessoa")
		return nil

	}

	c.Status(http.StatusOK).Send([]byte(userBytes))
	return nil
}

func BuscaPessoaPorTermo(c fiber.Ctx, db *gorm.DB) error {
	searchTerm := c.Query("t")

	if searchTerm == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}

	userTerm := "%" + searchTerm + "%"

	var users []entities.Pessoa
	db.Where("search_string LIKE ? LIMIT 50", userTerm).Find(&users)

	if len(users) == 0 {
		c.Status(http.StatusOK)
		c.Send([]byte("{}"))
		return nil
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

	userBytes, err := json.Marshal(usersReturn)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString("Something weird happened while trying to Marshal the user on BuscaPessoa")
		return nil

	}

	c.Status(http.StatusOK)
	c.Send(userBytes)
	return nil
}

func ContaPessoas(c fiber.Ctx, db *gorm.DB) error {
	var count int64
	db.Model(&entities.Pessoa{}).Count(&count)
	c.Set("content-type", "text/plain")
	c.Status(http.StatusOK).SendString(strconv.FormatInt(count, 10))

	return nil
}
