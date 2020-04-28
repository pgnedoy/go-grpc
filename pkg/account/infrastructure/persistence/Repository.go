package persistence

import (
	"github.com/pgnedoy/go-grpc/pkg/account/domain"
	"time"
)

type database struct {
	accounts []domain.Account
}

func (db *database)Find(id string) domain.Account {
	for _, account := range db.accounts {
		if account.Id == id {
			return account
		}
	}

	return domain.Account{
		Id:           "",
		UserId:       "",
		Transactions: nil,
		CreatedAt:    time.Time{},
	}
}


type Repository struct {
	//db *PostgresConnection
	db *database
}

func NewRepository() *Repository {
	transactions := []*domain.Transaction{
		{
			Id:          "1",
			AccountId:   "1",
			Type:        domain.Type_EXPENSES,
			Category:    domain.Category_FOOD,
			Count:       123.32,
			Description: "Awesome food",
			CreatedAt:   time.Now(),
		},
	}

	accounts := []domain.Account{
		{
			Id:           "1",
			UserId:       "11",
			Transactions: transactions,
			CreatedAt:    time.Time{},
		},
	}

	database := database{accounts}

	return &Repository{db: &database}
}

func (r *Repository)GetAccountById(id string) domain.Account {
	//mapping and error handling could be here
	return r.db.Find(id)
}


