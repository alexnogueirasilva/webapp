package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// APIURL API URL is the string connection with the database
	APIURL = ""

	// Port where the API will be running
	Port = 0

	// HashKey is the key used to sign the JWT
	HashKey []byte

	// BlockKey is the key used to encrypt the JWT
	BlockKey []byte
)

// Load loads the configuration variables
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
