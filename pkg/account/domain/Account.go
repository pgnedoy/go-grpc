package domain

import "time"

type Account struct {
	Id string
	UserId string
	Transactions []*Transaction
	CreatedAt time.Time
}