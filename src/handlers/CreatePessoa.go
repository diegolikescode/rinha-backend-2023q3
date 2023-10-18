package handlers

import (
	"net/http"
	"rinha-backend-2023q3/src/entities"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func validateValues(nome string, apelido string, nascimento string) (bool) {
    if(len(nome) > 100 || nome == ""){
	return false
    }

    if(len(apelido) > 32 || apelido == "") {
	return false
    }

    if(!entities.ValidaFormatoData(nascimento)) {
	return false
    }

    return true
}

func CreatePessoa(c *gin.Context, db *gorm.DB) {
    var jsonEntrada entities.CreatePessoaDTO
    if err := c.ShouldBindJSON(&jsonEntrada); err != nil {
	c.Writer.WriteHeader(http.StatusBadRequest)
	return 
    }

    if(!validateValues(jsonEntrada.Nome, jsonEntrada.Apelido, jsonEntrada.Nascimento)) {
	c.Writer.WriteHeader(http.StatusUnprocessableEntity)
	return
    }

    var user entities.Pessoa
    if db.Where("apelido = ?", jsonEntrada.Apelido).First(&user).RowsAffected > 0 {
	c.IndentedJSON(http.StatusUnprocessableEntity, entities.HttpResponse{
	    Message: "Esse apelido já está sendo utilizado",
	})
	return
    }

    newUUID := entities.CreateUUID()
    stackStr := strings.Join(jsonEntrada.Stack, ";")
    pessoaBody := entities.Pessoa{
	Id: newUUID,
	Apelido: jsonEntrada.Apelido,
	Nome: jsonEntrada.Nome,
	Nascimento: jsonEntrada.Nascimento,
	Stack: stackStr,
	SearchString: jsonEntrada.Apelido+jsonEntrada.Nome+stackStr,
    }

    db.Create(&pessoaBody)

    if db.Error != nil {
	print(db.Error)
    }

    c.Writer.WriteHeader(http.StatusCreated)
    c.Writer.Header().Add("Location", "/pessoas/"+ newUUID)
    return 
}

