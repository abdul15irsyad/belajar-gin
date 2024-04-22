package main

import (
	"gin-gorm-rest-api/database"
	"gin-gorm-rest-api/database/seeders"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	database.InitDatabase()

	seeders.SeedUsers()
	log.Println("all seeders executed")

	os.Exit(0)
}