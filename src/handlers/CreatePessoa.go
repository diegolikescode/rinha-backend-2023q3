package handlers

import (
	"net/http"
	"strings"

	"rinha-backend-2023q3/src/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func validateValues(nome string, apelido string, nascimento string) bool {
	if len(nome) > 100 || nome == "" {
		return false
	}

	if len(apelido) > 32 || apelido == "" {
		return false
	}

	if !entities.ValidaFormatoData(nascimento) {
		return false
	}

	return true
}

func CreatePessoa(c fiber.Ctx, db *gorm.DB) error {
	var jsonEntrada entities.CreatePessoaDTO
	if err := c.Bind().JSON(&jsonEntrada); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if !validateValues(jsonEntrada.Nome, jsonEntrada.Apelido, jsonEntrada.Nascimento) {
		c.Status(http.StatusUnprocessableEntity)
		return nil
	}

	var user entities.Pessoa
	if db.Where("apelido = ?", jsonEntrada.Apelido).First(&user).RowsAffected > 0 {
		c.Status(http.StatusUnprocessableEntity)
		return nil
	}

	newUUID := entities.CreateUUID()
	stackStr := strings.Join(jsonEntrada.Stack, ";")
	pessoaBody := entities.Pessoa{
		Id:           newUUID,
		Apelido:      jsonEntrada.Apelido,
		Nome:         jsonEntrada.Nome,
		Nascimento:   jsonEntrada.Nascimento,
		Stack:        stackStr,
		SearchString: jsonEntrada.Apelido + ";" + jsonEntrada.Nome + stackStr,
	}

	db.Create(&pessoaBody)
	if db.Error != nil {
		print(db.Error)
	}

	c.Status(http.StatusCreated)
	c.Set("Location", "/pessoas/"+newUUID)
	return nil
}
