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
	user, err := h.Repository.GetByTelegramID(update.FromChat().ID)
	if err != nil {
		log.Println("DB error:", err)
		return
	}

	if update.Message.Text == "/start" || user == nil {
		h.HandleStart(update)
		return
	} else if update.Message.Text == common.UserKeyboardMyBids {
		h.HandleGetBids(update)
	}

	switch *user.State {
	case common.UserStateAwaitingAgreement:
		h.handleAgreement(update, user)
	case common.UserRegistrationAwaitingInitials:
		h.handleInitials(update, user)
	case common.UserRegistrationAwaitingPhone:
		h.handlePhone(update, user)
	case common.UserRegistrationAwaitingEmail:
		h.handleEmail(update, user)
	case common.UserStateMainMenu:
		h.sendMainMenu(update)
	default:
		h.Bot.Send(tgbotapi.NewMessage(update.FromChat().ID, "Извините, я Вас не понял."))
		h.sendMainMenu(update)
	}
}

func (h *UserHandler) HandleStart(update tgbotapi.Update) {
	user, err := h.Repository.GetByTelegramID(update.FromChat().ID)
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

		user.State = common.Ptr(common.UserStateAwaitingAgreement)
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
	msg := tgbotapi.NewMessage(update.FromChat().ID, "Выберите действие:")
	msg.ReplyMarkup = common.UserKeyBoard
	h.Bot.Send(msg)
}

func (h *UserHandler) HandleGetBids(update tgbotapi.Update) {
	bids, err := h.Repository.GetBidsByTelegramID(update.FromChat().ID)
	if err != nil {
		log.Fatal(err)
	}

	if bids == nil || len(bids) == 0 {
		msg := tgbotapi.NewMessage(update.FromChat().ID, "У Вас сейчас нет активных ставок. Перейдите в канал с аукционами и сделайте свою первую ставку уже сейчас!")
		h.Bot.Send(msg)
	} else {
		msg := tgbotapi.NewMessage(update.FromChat().ID, models.BidsToString(bids))
		h.Bot.Send(msg)
	}
}
