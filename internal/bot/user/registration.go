package user

import (
	"car-auctions-telegram-bot/internal/common"
	"car-auctions-telegram-bot/internal/common/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
	"strings"
)

func (h *UserHandler) handleAgreement(update tgbotapi.Update, user *models.User) {
	if update.Message.Text != "Начать регистрацию" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мы не можем допустить вас до аукционов без предварительной регистрации. Чтобы продолжить, нажмите 'Начать регистрацию'")
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Начать регистрацию"),
			),
		)
		h.Bot.Send(msg)
		return
	}

	user.State = common.Ptr("awaiting_name")
	_ = h.Repository.Update(user)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите Вашу фамилию, имя и отчество:")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	h.Bot.Send(msg)
}

func (h *UserHandler) handleName(update tgbotapi.Update, user *models.User) {
	parts := strings.Split(update.Message.Text, " ")
	if len(parts) < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите корректные данные о Вашей фамилии, имени и отчестве:")
		h.Bot.Send(msg)
		return
	}

	user.FirstName = common.Ptr(parts[1])
	user.LastName = common.Ptr(parts[0])
	if len(parts) > 2 {
		user.MiddleName = common.Ptr(parts[2])
	}

	user.State = common.Ptr("awaiting_phone")
	_ = h.Repository.Update(user)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ваш номер телефона или воспользуйтесь кнопкой ниже:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Поделиться контактом"),
		))
	h.Bot.Send(msg)
}

func (h *UserHandler) handlePhone(update tgbotapi.Update, user *models.User) {
	if update.Message.Contact != nil {
		if ok, _ := regexp.MatchString("[+][0-9][1-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]", update.Message.Contact.PhoneNumber); !ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите корректный номер телефона, или воспользуйтесь кнопкой ниже:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButtonContact("Поделиться контактом"),
				))
			h.Bot.Send(msg)
			return
		} else {
			user.Phone = common.Ptr(update.Message.Contact.PhoneNumber)
		}
	} else {
		if ok, _ := regexp.MatchString("[+][0-9][1-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]", update.Message.Text); !ok {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите корректный номер телефона, или воспользуйтесь кнопкой ниже:")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButtonContact("Поделиться контактом"),
				))
			h.Bot.Send(msg)
			return
		} else {
			common.Ptr(update.Message.Contact.PhoneNumber)
		}
	}

	user.State = common.Ptr("awaiting_email")

	_ = h.Repository.Update(user)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите ваш адрес электронной почты:")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	h.Bot.Send(msg)
}

func (h *UserHandler) handleEmail(update tgbotapi.Update, user *models.User) {
	if ok, _ := regexp.MatchString("[a-zA-Z]+@[a-z]+[.][a-z]", update.Message.Text); !ok {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введите корректный адрес электронной почты")
		h.Bot.Send(msg)
		return
	}

	user.Email = common.Ptr(update.Message.Text)
	user.State = common.Ptr("main_menu")
	_ = h.Repository.Update(user)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отлично! Вы успешно зарегистрировались.")
	h.Bot.Send(msg)
}
