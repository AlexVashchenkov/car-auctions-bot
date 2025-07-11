package user

import (
	"car-auctions-telegram-bot/internal/common"
	"car-auctions-telegram-bot/internal/common/models"
	"car-auctions-telegram-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type UserHandler struct {
	Bot        *tgbotapi.BotAPI
	Repository *storage.UserRepository
}

func (h *UserHandler) HandleUpdate(update tgbotapi.Update) {
	user, err := h.Repository.GetByTelegramID(update.Message.From.ID)
	if err != nil {
		log.Println("DB error:", err)
		return
	}

	if update.Message.Text == "/start" || user == nil {
		h.HandleStart(update)
		return
	}

	switch *user.State {
	case "awaiting_agreement":
		h.handleAgreement(update, user)
	case "awaiting_name":
		h.handleName(update, user)
	case "awaiting_phone":
		h.handlePhone(update, user)
	case "awaiting_email":
		h.handleEmail(update, user)
	case "main_menu":
		h.sendMainMenu(update)
	default:
		h.sendMainMenu(update)
	}
}

func (h *UserHandler) HandleStart(update tgbotapi.Update) {
	user, err := h.Repository.GetByTelegramID(update.Message.From.ID)
	if err != nil {
		log.Fatal(err)
	}

	if user == nil {
		user = &models.User{
			TelegramID: update.Message.From.ID,
			LastName:   common.Ptr(""),
			FirstName:  common.Ptr(""),
			Phone:      common.Ptr(""),
			Email:      common.Ptr(""),
		}

		user.State = common.Ptr("awaiting_agreement")
		err = h.Repository.Create(user)
		if err != nil {
			log.Fatal(err)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, common.GreetingRegistrationMessage)
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Начать регистрацию"),
			),
		)
		h.Bot.Send(msg)
	} else {
		h.sendMainMenu(update)
	}
}

func (h *UserHandler) sendMainMenu(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Мои ставки"),
		),
	)
	h.Bot.Send(msg)
}
