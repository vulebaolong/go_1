package env

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	IsProduction bool
	Port         string
	Host         string
	DatabaseUrl  string

	ExpiresAtAccessToken time.Duration
	SecretAccessToken    string

	ExpiresAtRefreshToken time.Duration
	SecretRefreshToken    string
}

func New() *Env {
	godotenv.Load()

	isProduction := os.Getenv("IS_PRODUCTION") == "true"
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	databaseUrl := os.Getenv("DATABASE_URL")

	// ACCESS TOKEN
	expiresAtAccessTokenString := os.Getenv("EXPIRES_AT_ACCESS_TOKEN")
	expiresAtAccessToken := getDuration(expiresAtAccessTokenString)
	secretAccessToken := os.Getenv("SECRET_ACCESS_TOKEN")

	// REFRESH TOKEN
	expiresAtRefreshTokenString := os.Getenv("EXPIRES_AT_REFRESH_TOKEN")
	expiresAtRefreshToken := getDuration(expiresAtRefreshTokenString)
	secretRefreshToken := os.Getenv("SECRET_REFRESH_TOKEN")

	fmt.Println("isProduction", isProduction)
	fmt.Println("port", port)
	fmt.Println("host", host)
	fmt.Println("databaseUrl", databaseUrl)

	fmt.Println("expiresAtAccessTokenString", expiresAtAccessTokenString)
	fmt.Println("secretAccessToken", secretAccessToken)

	fmt.Println("expiresAtRefreshTokenString", expiresAtRefreshTokenString)
	fmt.Println("secretRefreshToken", secretRefreshToken)

	return &Env{
		IsProduction:          isProduction,
		Port:                  port,
		Host:                  host,
		DatabaseUrl:           databaseUrl,
		ExpiresAtAccessToken:  expiresAtAccessToken,
		SecretAccessToken:     secretAccessToken,
		ExpiresAtRefreshToken: expiresAtRefreshToken,
		SecretRefreshToken:    secretRefreshToken,
	}
}

func getDuration(durationString string) time.Duration {
	durationTime, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatal("Parser expiresAtAccessToken error")
	}
	return durationTime
}
