package models

import (
	"time"
)

type User struct {
	ID         int64     `db:"id"`
	TelegramID int64     `db:"telegram_id"`
	LastName   *string   `db:"last_name"`
	FirstName  *string   `db:"first_name"`
	MiddleName *string   `db:"middle_name"`
	Phone      *string   `db:"phone"`
	Email      *string   `db:"email"`
	State      *string   `db:"state"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
