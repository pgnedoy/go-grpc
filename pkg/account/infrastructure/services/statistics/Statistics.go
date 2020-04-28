package statistics

import (
	"context"
	"errors"
	timestamp "github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	pb "github.com/pgnedoy/go-grpc/pb/statistics"
	"github.com/pgnedoy/go-grpc/pkg/account/domain"
	"log"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendAccount(account domain.Account) error {
	serverAddr := "github.com/pgnedoy/go-grpc:8080"

	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewStatisticsClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//context.
	//defer cancel()

	var transactions []*pb.Transaction

	for _, transaction := range account.Transactions {
		createdAt, err := timestamp.TimestampProto(transaction.CreatedAt)
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}

		transactions = append(transactions, &pb.Transaction{
			Id:          transaction.Id,
			Type:        pb.Type(transaction.Type),
			Category:    pb.Category(transaction.Category),
			Count:       transaction.Count,
			Description: transaction.Description,
			CreatedAt:   createdAt,
		})
	}

	pbAccount := pb.Account{
		Id:           account.Id,
		UserId:       account.UserId,
		Transactions: transactions,
	}
	res, err := client.SendAccount(context.Background(), &pbAccount)
	if err != nil {
		log.Fatalf("%v.SendAccount(_) = _, %v: ", client, err)
		return errors.New("Error sending account to Stats service")
	}
	log.Println(res)
	return nil
}

func (s *Service) SendTransaction() error {
	serverAddr := "github.com/pgnedoy/go-grpc:8080"
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewStatisticsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	time, _ := timestamp.TimestampProto(time.Now())
	transactions := []*pb.Transaction{
		{Id:"1", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 101, Description: "", CreatedAt: time },
		{Id:"2", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 102, Description: "", CreatedAt: time },
		{Id:"3", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 103, Description: "", CreatedAt: time },
		{Id:"4", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 104, Description: "", CreatedAt: time },
		{Id:"5", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 105, Description: "", CreatedAt: time },
		{Id:"6", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 106, Description: "", CreatedAt: time },
		{Id:"7", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 107, Description: "", CreatedAt: time },
		{Id:"8", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 108, Description: "", CreatedAt: time },
		{Id:"9", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 109, Description: "", CreatedAt: time },
		{Id:"10", Type: pb.Type_EXPENSES, Category: pb.Category_FOOD, Count: 110, Description: "", CreatedAt: time },
	}

	stream, err := client.SendTransaction(ctx)
	if err != nil {
		log.Fatalf("%v.SendTransaction(_) = _, %v", client, err)
		return err
	}

	for _, transaction := range transactions {
		if err := stream.Send(transaction); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, transaction, err)
			return err
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		return err
	}
	log.Printf("Route summary: %v", reply)


	return nil
}

