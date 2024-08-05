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

	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	if OPENAI_API_KEY == "" {
		log.Fatalln("You must provide an API key!")
	}

	if os.Getenv("DEV_MODE") == "true" {
		DEV_MODE = true
	}
}

var OPENAI_API_KEY string = "not initialized"
var DEV_MODE bool
