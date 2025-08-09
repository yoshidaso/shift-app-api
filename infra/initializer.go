package infra

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	// .env が無くても環境変数が設定されていれば続行できるため、致命的エラーにはしない
	if err := godotenv.Load(); err != nil {
		log.Printf(".env file not found; continuing with existing environment variables")
	}
}
