package cmd

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	envf := "conf/.env.local"
	err := godotenv.Load(envf)
	if err != nil {
		log.Fatalf("Error loading %s file.\n", envf)
	} else {
		log.Printf("Successfully load %s file.\n", envf)
	}
}
