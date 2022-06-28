package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URI = ""
	API_PORT     = 0
	SECRET_KEY   []byte
)

func LoadConfig() {

	var err error

	if err := godotenv.Load(); err != nil {
		log.Print(err)
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 9000
	}

	DATABASE_URI = os.Getenv("DATABASE_URI")
}
