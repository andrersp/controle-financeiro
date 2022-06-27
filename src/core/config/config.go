package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDb  = ""
	API_PORT  = 0
	SecretKey []byte
)

// Load envs
func Load() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro on load env")
	}
	API_PORT, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		API_PORT = 9000
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_HOST := os.Getenv("DB_HOST")

	StringDb = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER, DB_PASSWORD, DB_HOST, DB_DATABASE)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
