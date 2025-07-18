package models

import (
	"fmt"
	"time"
)

type Bid struct {
	ID         int64     `db:"id"`
	TelegramID int64     `db:"user_id"`
	AuctionID  int64     `db:"auction_id"`
	Amount     int64     `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (b *Bid) ToString() string {
	return fmt.Sprintf("Ваша ставка на аукцион №%d: %d", b.AuctionID, b.Amount)
}

func BidsToString(bids []Bid) string {
	output := "Ваши ставки на аукционах:\n"
	for _, bid := range bids {
		formattedTime := bid.UpdatedAt.Format("02.01.2006 15:04:05")
		output += fmt.Sprintf("\nАукцион №%d --- %d рублей (сделана %s)\n", bid.ID, bid.Amount, formattedTime)
	}
	return output
}
