package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func PostgresConnection() {
	var err error

	host := "localhost"
	port := "5432"
	username := "postgres"
	password := "postgres"
	databaseName := "rinha"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, databaseName)
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database, error: %v", err)
	}

	config.MaxConns = 40
	// config.MaxConnLifetime = 20
	config.MaxConnLifetime = time.Hour
	config.MinConns = 10
	config.HealthCheckPeriod = 30 * time.Second

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool, error: %v", err)
	}

	// connectionString := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	host, username, password, databaseName, port)
	// db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Successfully connected to postgres")
	// }
	//
	// return db
}
