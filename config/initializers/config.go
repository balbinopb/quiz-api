package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		} else {
			log.Println(".env loaded successfully")
		}
	}
}


func GetDBUrl() string {
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        user, pass, host, port, name)
}