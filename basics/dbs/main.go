package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func getConnect(user, password string) *sql.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres", user, password)
	conn, err := sql.Open("pgx", dbURL)
	if err != nil {
		slog.Error("error while connecting db", "err", err)
	}
	return conn
}

var tbQuery = `
	CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

type user struct {
	username, email string
}

type userResult struct {
	id              int
	username, email string
	cretedAt        time.Time
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	conn := getConnect(postgresUser, postgresPassword)
	err = conn.PingContext(ctx)

	if err != nil {
		slog.Error("Error while pinging db", "err", err)
	}

	//create table
	_, err = conn.ExecContext(ctx, tbQuery)

	if err != nil {
		slog.Error("Error while creating table ", "err", err)
	}

	jon := user{"jon doe", "jon@doe.com"}

	_, err = conn.QueryContext(ctx, "INSERT INTO users (username, email) VALUES ($1,$2)", jon.username, jon.email)

	if err != nil {
		slog.Error("Error while creating table ", "err", err)
	}

	rows, errQuerty := conn.QueryContext(ctx, "SELECT * FROM users")

	if errQuerty != nil {
		slog.Error("Error while querying table ", "err", errQuerty)
	}

	var uResult []userResult

	for rows.Next() {
		rec := userResult{}
		rows.Scan(&rec.id, &rec.username, &rec.email, &rec.cretedAt)

		uResult = append(uResult, rec)

	}

	fmt.Println(uResult)

	conn.BeginTx()
}
