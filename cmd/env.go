package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env() {
	log.Printf("System Enviroment : %s", os.Getenv("GO_ACTIVE_PROFILE"))
	err := godotenv.Load(os.Getenv("GO_ACTIVE_PROFILE") + ".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
