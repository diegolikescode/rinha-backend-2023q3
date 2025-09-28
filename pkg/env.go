package pkg

import (
	"fmt"
	"os"
)

func GetEnvOrDieTrying(key string) string {
	value := os.Getenv(key)
	if value == "" {
		err := fmt.Errorf("Missing env var %s", key)
		panic(err)
	}

	return value
}
