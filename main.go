package main

import (
	_ "github.com/joho/godotenv/autoload"
	application "github.com/naaltunian/wyn-search-go/app"
)

func main() {
	application.StartApplication()
}
