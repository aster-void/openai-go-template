package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if envfile := os.Getenv("ENV_FILE"); envfile != "" {
		err := godotenv.Load(envfile)
		if err != nil {
			log.Fatalln("Failed to load env file:", err)
		}
	}

	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEy")
	if OPENAI_API_KEY == "" {
		log.Fatalln("You must provide an API key!")
	}
}

var OPENAI_API_KEY string = "not initialized"
