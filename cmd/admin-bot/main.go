package main

import (
	"car-auctions-telegram-bot/internal/bot/admin"
	"car-auctions-telegram-bot/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while opening .env file", err)
	}

	token := os.Getenv("ADMIN_BOT_TOKEN")
	if token == "" {
		log.Fatal("ADMIN_BOT_TOKEN not set")
	}

	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		log.Fatal("Database connection string not set")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("User Bot authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	db, err := sqlx.Connect("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}

	adminRepo := storage.NewAdminRepository(db)
	userRepo := storage.NewUserRepository(db)

	handler := admin.AdminHandler{AdminRepo: adminRepo, UserRepo: userRepo, Bot: bot}

	for update := range updates {
		handler.HandleUpdate(update)
	}
}
