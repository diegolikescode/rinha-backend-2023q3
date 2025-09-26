package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/entities"

	"github.com/gofiber/fiber/v3"
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

func CreatePessoa(c fiber.Ctx) error {
	var jsonEntrada entities.CreatePessoaDTO
	if err := c.Bind().JSON(&jsonEntrada); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if !validateValues(jsonEntrada.Nome, jsonEntrada.Apelido, jsonEntrada.Nascimento) {
		c.Status(http.StatusUnprocessableEntity)
		return nil
	}

	uid := entities.CreateUUID()
	stackStr := strings.Join(jsonEntrada.Stack, ";")
	searchStr := jsonEntrada.Apelido + ";" + jsonEntrada.Nome + ";" + stackStr

	res, err := config.Pool.Exec(context.Background(), `
		INSERT INTO pessoas (id, apelido, nome, nascimento, stack, search_string, api_name)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (apelido) DO NOTHING
	`, uid, jsonEntrada.Apelido, jsonEntrada.Nome, jsonEntrada.Nascimento, stackStr, searchStr, os.Getenv("API_NAME"))
	if err != nil {
		fmt.Println("MY ERROR:::", err)
		c.Status(http.StatusInternalServerError)
		return nil
	}

	if res.RowsAffected() == 0 {
		c.Status(http.StatusUnprocessableEntity)
		return nil
	}

	c.Status(http.StatusCreated)
	c.Set("Location", "/pessoas/"+uid)
	return nil
}
