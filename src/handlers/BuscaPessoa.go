package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"rinha-backend-2023q3/src/config"
	"rinha-backend-2023q3/src/entities"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func BuscaPessoa(c fiber.Ctx) error {
	uid := c.Params("id")
	var u entities.ReturnPessoa
	var strStack string
	err := config.Pool.QueryRow(context.Background(),
		"SELECT id, apelido, nome, nascimento, stack FROM pessoas WHERE id = $1 LIMIT 1", uid).Scan(
		&u.Id, &u.Apelido, &u.Nome, &u.Nascimento, &strStack)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.Status(http.StatusNotFound)
			return nil
		}
		c.Status(http.StatusInternalServerError)
		return nil
	}

	u.Stack = strings.Split(strStack, ";")
	uBytes, err := json.Marshal(u)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.SendString("Something weird happened while trying to Marshal the user on BuscaPessoa")
		return nil

	}

	c.Status(http.StatusOK).Send([]byte(uBytes))
	return nil
}

func BuscaPessoaPorTermo(c fiber.Ctx) error {
	searchTerm := c.Query("t")

	if searchTerm == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}

	userTerm := "%" + searchTerm + "%"

	rows, err := config.Pool.Query(context.Background(), `SELECT id, apelido, nome, nascimento, stack
		FROM pessoas
		WHERE search_string LIKE $1 LIMIT 50`, userTerm)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	defer rows.Close()

	var usersReturn []entities.ReturnPessoa
	for rows.Next() {
		var p entities.ReturnPessoa
		var stackStr string
		err := rows.Scan(&p.Id, &p.Apelido, &p.Nome, &p.Nascimento, &stackStr)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return nil
		}

		usersReturn = append(usersReturn, entities.ReturnPessoa{
			Id:         p.Id,
			Apelido:    p.Apelido,
			Nome:       p.Nome,
			Nascimento: p.Nascimento,
			Stack:      strings.Split(stackStr, ";"),
		})
	}
	if len(usersReturn) == 0 {
		c.Status(http.StatusOK)
		c.Send([]byte("[]"))
		return nil
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

func ContaPessoas(c fiber.Ctx) error {
	var count int64
	// db.Model(&entities.Pessoa{}).Count(&count)

	err := config.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM pessoas").Scan(&count)
	if err != nil {
		log.Fatal("FUCK THIS! ERROR:", err)
		c.Set("content-type", "text/plain")
		c.Status(http.StatusOK).SendString("WE FUCKED UP")
		return nil
	}

	c.Set("content-type", "text/plain")
	c.Status(http.StatusOK).SendString(strconv.FormatInt(count, 10))

	return nil
}
