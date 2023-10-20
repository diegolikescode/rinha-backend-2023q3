package entities

import (
	"regexp"

	"github.com/satori/go.uuid"
)

type CreatePessoaDTO struct {
    Apelido string `json:"apelido"`
    Nome string `json:"nome"`
    Nascimento string `json:"nascimento"`
    Stack []string `json:"stack"`
}

type HttpResponse struct {
    Message string `json:"message"`
}

type Pessoa struct {
    Id string `gorm:"uniqueIndex;primary_key;type:string" json:"id"`
    Apelido string `json:"apelido"`
    Nome string `json:"nome"`
    Nascimento string `json:"nascimento"`
    Stack string `json:"stack"`
    SearchString string `json:"searchString"`
}

type ReturnPessoa struct {
    Id string `json:"id"`
    Apelido string `json:"apelido"`
    Nome string `json:"nome"`
    Nascimento string `json:"nascimento"`
    Stack []string `json:"stack"`
}

func CreateUUID () (string) {
    return uuid.NewV4().String()
}

func ValidaFormatoData (data string) (bool) {
    datePattern := `^\d{4}-\d{2}-\d{2}$`
    matched, err := regexp.Match(datePattern, []byte(data))
    if err != nil {
	return false
    }

    return matched
}


