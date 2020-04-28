package application

import (
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/persistence"
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/services/statistics"
)

type SendTransaction struct {
	db          *persistence.Repository
	statService *statistics.Service
}

func NewSendTransactionUseCase(
	db *persistence.Repository,
	statService *statistics.Service,
) *SendTransaction {
	return &SendTransaction{db, statService}
}

func (uc *SendTransaction) Execute() error {
	return uc.statService.SendTransaction()
}
