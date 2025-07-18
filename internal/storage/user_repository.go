package storage

import (
	"car-auctions-telegram-bot/internal/common/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE telegram_id=$1`, telegramID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.db.NamedExec(`
        INSERT INTO users (telegram_id, last_name, first_name, middle_name, phone, email, state)
        VALUES (:telegram_id, :last_name, :first_name, :middle_name, :phone, :email, :state)
    `, user)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	query := "UPDATE users SET "
	params := map[string]interface{}{
		"telegram_id": user.TelegramID,
	}
	var setParts []string

	if user.LastName != nil {
		setParts = append(setParts, "last_name=:last_name")
		params["last_name"] = *user.LastName
	}
	if user.FirstName != nil {
		setParts = append(setParts, "first_name=:first_name")
		params["first_name"] = *user.FirstName
	}
	if user.MiddleName != nil {
		setParts = append(setParts, "middle_name=:middle_name")
		params["middle_name"] = *user.MiddleName
	}
	if user.Phone != nil {
		setParts = append(setParts, "phone=:phone")
		params["phone"] = *user.Phone
	}
	if user.Email != nil {
		setParts = append(setParts, "email=:email")
		params["email"] = *user.Email
	}
	if user.State != nil {
		setParts = append(setParts, "state=:state")
		params["state"] = *user.State
	}

	if len(setParts) == 0 {
		return fmt.Errorf("nothing to update")
	}

	// Always update updated_at
	setParts = append(setParts, "updated_at=now()")

	query += strings.Join(setParts, ", ")
	query += " WHERE telegram_id=:telegram_id"

	_, err := r.db.NamedExec(query, params)
	return err
}

func (r *UserRepository) GetBidsByTelegramID(id int64) ([]models.Bid, error) {
	var bids []models.Bid
	err := r.db.Get(&bids, `SELECT (id, user_id, auction_id, amount, created_at, updated_at) FROM bids WHERE user_id=$1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return bids, nil
	}
}
