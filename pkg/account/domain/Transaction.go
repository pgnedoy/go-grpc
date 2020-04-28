package domain

import "time"

type Transaction struct {
	Id string
	AccountId string
	Type Type
	Category Category
	Count float64
	Description string
	CreatedAt time.Time
}