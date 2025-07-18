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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, common.AdminBotErrProhibited)
		h.Bot.Send(msg)
		return
	}

	if update.Message.Text == "/start" {
		h.handleStart(update)
		return
	} else if update.Message.Text == common.AdminKeyboardCreateAuction {
		h.handleCreateAuction(update)
	} else if update.Message.Text == common.AdminKeyboardAuctionList {
		h.handleAuctionList(update)
	} else if update.Message.Text == common.AdminKeyBoardEndAuction {
		h.handleEndAuction(update)
	} else if update.Message.Text == common.AdminKeyBoardDeleteAuction {
		h.handleDeleteAuction(update)
	} else if update.Message.Text == common.AdminKeyboardAddAdmin {
		h.handleAddAdmin(update)
	}

	switch *admin.State {
	case common.AdminStateMainMenu:
		h.sendMainMenu(update)
	default:
		h.sendMainMenu(update)
	}
}

func (h *AdminHandler) sendMainMenu(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "Выберите действие:")
	msg.ReplyMarkup = common.AdminKeyboard
	h.Bot.Send(msg)
}

func (h *AdminHandler) handleCreateAuction(update tgbotapi.Update) {

}

func (h *AdminHandler) handleAuctionList(update tgbotapi.Update) {

}

func (h *AdminHandler) handleEndAuction(update tgbotapi.Update) {

}

func (h *AdminHandler) handleDeleteAuction(update tgbotapi.Update) {

}

func (h *AdminHandler) handleAddAdmin(update tgbotapi.Update) {

}

func (h *AdminHandler) handleStart(update tgbotapi.Update) {
	h.sendMainMenu(update)
	return
}
