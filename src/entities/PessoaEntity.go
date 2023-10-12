package entities

import (
	"regexp"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
    gorm.Model
    Id uuid.UUID `gorm:"uniqueIndex;primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
    Apelido string `json:"apelido"`
    Nome string `json:"nome"`
    Nascimento string `json:"nascimento"`
    Stack string `json:"stack"`
}

func ValidaFormatoData (data string) (bool) {
    datePattern := `^\d{4}-\d{2}-\d{2}$`
    matched, err := regexp.Match(datePattern, []byte(data))
    if err != nil {
	return false
    }

    return matched
}


