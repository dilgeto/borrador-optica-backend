package Injector

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func InjectDependencies(rout *gin.Engine) {
	errData := godotenv.Load()
	if errData != nil {
		log.Fatalf("Error while loading .env: -%v", errData)
	}

	db, err := ConnectPostgreSQL(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("optica"))
	if err != nil {
		log.Fatalf("Failed to connect to database: - %v", err)
	}

	// TODO: Rellenar con las entidades
}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, error) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + "port" + "/" + dbname + "?sslmode=disable"
	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}
