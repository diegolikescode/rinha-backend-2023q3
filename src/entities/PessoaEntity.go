package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomDate struct {
    time.Time
}

type Pessoa struct {
    gorm.Model
    Id uuid.UUID `gorm:"uniqueIndex;primary_key;type:uuid;default:gen_random_uuid()" json:"id"`
    Apelido string `json:"apelido"`
    Nome string `json:"nome"`
    // Nascimento CustomDate `json:"nascimento"`
    // Stack []string `json:"stack"`
}

func (c *CustomDate) UnmarshalJSON(data []byte) error {
    dateStr := strings.Trim(string(data), "\"")
    parsedDate, err := time.Parse("2006-04-17", dateStr)
    if err != nil {
	return err
    }

    c.Time = parsedDate
    return nil
}

