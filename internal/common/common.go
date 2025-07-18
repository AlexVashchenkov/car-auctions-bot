package common

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const GreetingRegistrationMessage = `
Приветствуем Вас в боте-аукционисте! 
Так как сейчас вы не зарегистрированный пользователь, мне понадобятся Ваши данные.
Начнём регистрацию?
`

const AdminBotErrProhibited = `
Вы не являетесь зарегистрированным админом. Пожалуйста, обратитесь к руководству проекта
`

const ErrDatabase = `
Произошла ошибка. Пожалуйста, обратитесь в поддержку по реквизитам:
`

const AdminKeyboardCreateAuction = "🔨 Создать аукцион"
const AdminKeyboardAuctionList = "📜 Список аукционов"
const AdminKeyBoardEndAuction = "⏳ Завершить аукцион"
const AdminKeyBoardDeleteAuction = "🗑️ Удалить аукцион"
const AdminKeyboardAddAdmin = "👔 Добавить админа"

var AdminKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AdminKeyboardCreateAuction),
		tgbotapi.NewKeyboardButton(AdminKeyboardAuctionList),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AdminKeyBoardEndAuction),
		tgbotapi.NewKeyboardButton(AdminKeyBoardDeleteAuction),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(AdminKeyboardAddAdmin),
	),
)

const AdminStateInitial = "initial"
const AdminStateMainMenu = "admin_main_menu"
const AdminStateCreateAuctionStep1 = "admin_auction_create_step_1"
const AdminStateCreateAuctionStep2 = "admin_auction_create_step_2"
const AdminStateCreateAuctionStep3 = "admin_auction_create_step_3"

const AdminStateDeleteAuctionStep1 = "admin_auction_delete_step_1"
const AdminStateDeleteAuctionStep2 = "admin_auction_delete_step_2"
const AdminStateDeleteAuctionStep3 = "admin_auction_delete_step_3"

const AdminStateEndAuctionStep1 = "admin_auction_end_step_1"
const AdminStateEndAuctionStep2 = "admin_auction_end_step_2"
const AdminStateEndAuctionStep3 = "admin_auction_end_step_3"

const AdminStateAuctionListStep1 = "admin_auction_list_step_1"
const AdminStateAuctionListStep2 = "admin_auction_list_step_2"
const AdminStateAuctionListStep3 = "admin_auction_list_step_3"

const AdminStateAddAdminStep1 = "admin_admin_add_step_1"
const AdminStateAddAdminStep2 = "admin_admin_add_step_2"
const AdminStateAddAdminStep3 = "admin_admin_add_step_3"

func Ptr[T any](v T) *T {
	return &v
}

const UserKeyboardMyBids = "📜 Мои ставки"

var UserKeyBoard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(UserKeyboardMyBids),
	),
)

const UserStateAwaitingAgreement = "user_awaiting_agreement"
const UserStateMainMenu = "user_main_menu"

const UserRegistrationAwaitingInitials = "user_registration_awaiting_initials"
const UserRegistrationAwaitingPhone = "user_registration_step_2"
const UserRegistrationAwaitingEmail = "user_registrations_step_3"
