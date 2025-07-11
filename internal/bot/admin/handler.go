package admin

import (
	"car-auctions-telegram-bot/internal/common"
	"car-auctions-telegram-bot/internal/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "regexp"
)

type AdminHandler struct {
	Bot       *tgbotapi.BotAPI
	AdminRepo *storage.AdminRepository
	UserRepo  *storage.UserRepository
}

func (h *AdminHandler) HandleUpdate(update tgbotapi.Update) {
	admin, err := h.AdminRepo.GetByTelegramID(update.Message.From.ID)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, common.ErrDatabase)
		h.Bot.Send(msg)
		return
	}

	if admin == nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, common.AdminBotProhibited)
		h.Bot.Send(msg)
		return
	}

	if update.Message.Text == "/start" {
		h.HandleStart(update)
		return
	}

	switch *admin.State {
	case "main_menu":
		h.sendMainMenu(update)
	default:
		h.sendMainMenu(update)
	}
}

func (h *AdminHandler) HandleStart(update tgbotapi.Update) {
	h.sendMainMenu(update)
	return
}

func (h *AdminHandler) sendMainMenu(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Мои ставки"),
		),
	)
	h.Bot.Send(msg)
}
