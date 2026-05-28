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

	GoogleClientId     string
	GoogleClientSecret string
	GoogleRedirectUrl  string

	DomainFe string

	CloudinaryUrl string
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

	// GOOGLE
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectUrl := os.Getenv("GOOGLE_REDIRECT_URL")

	domainFe := os.Getenv("DOMAIN_FE")

	cloudinaryUrl := os.Getenv("CLOUDINARY_URL")

	fmt.Println("isProduction", isProduction)
	fmt.Println("port", port)
	fmt.Println("host", host)
	fmt.Println("databaseUrl", databaseUrl)

	fmt.Println("expiresAtAccessTokenString", expiresAtAccessTokenString)
	fmt.Println("secretAccessToken", secretAccessToken)

	fmt.Println("expiresAtRefreshTokenString", expiresAtRefreshTokenString)
	fmt.Println("secretRefreshToken", secretRefreshToken)

	fmt.Println("googleClientId", googleClientId)
	fmt.Println("googleClientSecret", googleClientSecret)
	fmt.Println("googleRedirectUrl", googleRedirectUrl)

	fmt.Println("domainFe", domainFe)

	fmt.Println("cloudinaryUrl", cloudinaryUrl)

	return &Env{
		IsProduction:          isProduction,
		Port:                  port,
		Host:                  host,
		DatabaseUrl:           databaseUrl,
		ExpiresAtAccessToken:  expiresAtAccessToken,
		SecretAccessToken:     secretAccessToken,
		ExpiresAtRefreshToken: expiresAtRefreshToken,
		SecretRefreshToken:    secretRefreshToken,
		GoogleClientId:        googleClientId,
		GoogleClientSecret:    googleClientSecret,
		GoogleRedirectUrl:     googleRedirectUrl,
		DomainFe:              domainFe,
		CloudinaryUrl:         cloudinaryUrl,
	}
}

func getDuration(durationString string) time.Duration {
	durationTime, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatal("Parser expiresAtAccessToken error")
	}
	return durationTime
}
