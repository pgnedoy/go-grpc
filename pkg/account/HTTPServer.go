package account

import (
	"github.com/pgnedoy/go-grpc/pkg/account/application"
	"github.com/pgnedoy/go-grpc/pkg/account/domain"
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/persistence"
	"github.com/pgnedoy/go-grpc/pkg/account/infrastructure/services/statistics"
	"log"
	"net/http"
	"time"
)

type HTTPService struct {
	CreateTransaction application.CreateTransaction
}

func StartHTTPServer() {
	repository := persistence.NewRepository()
	statService := statistics.NewService()
	createTransactionUseCase := application.NewCreateTransactionUseCase(
		repository,
		statService,
	)

	sendTransactionUseCase := application.NewSendTransactionUseCase(
		repository,
		statService,
	)

	createTransactionHandler := func(w http.ResponseWriter, req *http.Request) {
		transaction := domain.Transaction{
			Id:"1", AccountId:   "1", Type: 0, Category: 0, Count: 0, Description: "", CreatedAt: time.Time{},
		}
		createTransactionUseCase.Execute(transaction)
	}

	sendTransactionHandler := func(w http.ResponseWriter, req *http.Request) {
		sendTransactionUseCase.Execute()
	}

	http.HandleFunc("/create", createTransactionHandler)
	http.HandleFunc("/send", sendTransactionHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
