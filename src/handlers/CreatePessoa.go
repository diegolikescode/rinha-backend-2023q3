package handlers

import (
	"net/http"
	"reflect"
	"rinha-backend-2023q3/src/entities"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePessoa(c *gin.Context, db *gorm.DB) {
    var jsonEntrada entities.CreatePessoaDTO
    c.ShouldBindJSON(&jsonEntrada)

    if(len(jsonEntrada.Nome) > 100 || jsonEntrada.Nome == "" || reflect.TypeOf(jsonEntrada.Nome).Kind() != reflect.String) {
	c.IndentedJSON(http.StatusUnprocessableEntity, entities.HttpResponse{
	    Message: "nome eh obrigatorio e deve ser menor que 100 chars",
	})
	return
    }

    if(len(jsonEntrada.Apelido) > 32 || jsonEntrada.Apelido == "") {
	c.IndentedJSON(http.StatusUnprocessableEntity, entities.HttpResponse{
	    Message: "apelido eh obrigatorio e deve ser menor que 32 chars",
	})
	return
    }

    if(!entities.ValidaFormatoData(jsonEntrada.Nascimento)) {
	c.IndentedJSON(http.StatusUnprocessableEntity, entities.HttpResponse{
	    Message: "campo nascimento formatado incorretamente (esperado YYYY-MM-DD)",
	})
	return
    }

    var user entities.Pessoa
    if db.Where("apelido = ?", jsonEntrada.Apelido).First(&user).RowsAffected > 0 {
	c.IndentedJSON(http.StatusUnprocessableEntity, entities.HttpResponse{
	    Message: "Esse apelido já está sendo utilizado",
	})
	return
    }

    // TODO: devia verificar o tipo da Stack tambem?

    pessoaBody := entities.Pessoa{
	Apelido: jsonEntrada.Apelido,
	Nome: jsonEntrada.Nome,
	Nascimento: jsonEntrada.Nascimento,
	Stack: strings.Join(jsonEntrada.Stack, ";"),
    }

    db.Create(&pessoaBody)

    if db.Error != nil {
	print(db.Error)
    }

    c.IndentedJSON(http.StatusOK, entities.HttpResponse{
	Message: "Pessoa criada com sucesso",
    })
    return 
}

