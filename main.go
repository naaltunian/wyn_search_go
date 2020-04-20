package main

import (
	_ "github.com/joho/godotenv/autoload"
	application "github.com/naaltunian/wyn-search-go/app"
	"github.com/naaltunian/wyn-search-go/database"
)

func main() {
	database.ConnectToDatabase()
	application.StartApplication()
}
