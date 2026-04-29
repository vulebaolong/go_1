package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	IsProduction bool
	Port         string
	Host         string
}

func New() *Env {
	godotenv.Load()

	isProduction := os.Getenv("IS_PRODUCTION") == "true"
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Println("isProduction", isProduction)
	fmt.Println("port", port)
	fmt.Println("host", host)

	return &Env{
		IsProduction: isProduction,
		Port:         port,
		Host:         host,
	}
}
