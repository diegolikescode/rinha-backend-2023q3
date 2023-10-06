package config

import (
	"fmt"
	"log"
	"rinha-backend-2023q3/src/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func PostgresConnection() (db *gorm.DB){
    var err error

    host := "localhost"
    username := "admin"
    password := "admin"
    databaseName := "rinha"
    port := "5432"

    connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
    db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully connected to postgres")
    }

    db.Migrator().DropTable(&entities.Pessoa{})
    
    if err := db.AutoMigrate(&entities.Pessoa{}); err != nil {
	log.Fatalf("AutoMigrate failed: %v", err)
    }

    return db
}

