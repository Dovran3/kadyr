package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST   string
	PORT   string
	DB_URI string
}

var ENV Config

func Init_config() {
	godotenv.Load()
	ENV.HOST = os.Getenv("HOST")
	ENV.PORT = os.Getenv("PORT")
	ENV.DB_URI = os.Getenv("DB_URI")
}
