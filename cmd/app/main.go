package main

import (
	"LibSongs/internal/app"
	"github.com/joho/godotenv"
	"log"
)

const configsDir = "configs"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// @title LibrarySongs API
// @version 1.0
// @description API Server for LibrarySongs Application

// @host localhost:8000
// @BasePath /

func main() {

	app.Run(configsDir)

}
