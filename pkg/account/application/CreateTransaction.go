package application

import (
	"github.com/pgnedoy/go-grpc/pkg/account/domain"
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/persistence"
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/services/statistics"
)

type CreateTransaction struct {
	db          *persistence.Repository
	statService *statistics.Service
}

func NewCreateTransactionUseCase(
	db *persistence.Repository,
	statService *statistics.Service,
) *CreateTransaction {
	return &CreateTransaction{db, statService}
}

func (uc *CreateTransaction) Execute(transaction domain.Transaction) error {
	//save transaction to db
	//get account with transactions
	account := uc.db.GetAccountById(transaction.AccountId)
	return uc.statService.SendAccount(account)
}
