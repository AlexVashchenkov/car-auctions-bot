package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error, while opening .env file", err)
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	fmt.Println(token)
}
