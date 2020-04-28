package statistics

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "github.com/pgnedoy/go-grpc/pb/statistics"
	"log"
	"net"
	"time"
)

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)
	fmt.Println(info.FullMethod)
	fmt.Println(time.Since(start))
	return h, err
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

type server struct {}

func (s *server)SendTransaction(stream pb.Statistics_SendTransactionServer) error {
	var transactionCount int32
	startTime := time.Now()

	for {
		transaction, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.Summary{
				Count: transactionCount,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}

		if err != nil {
			return err
		}
		fmt.Println(transaction)
		transactionCount++
	}
}

func (s *server)SendAccount(ctx context.Context, account *pb.Account) (*pb.Response, error) {
	fmt.Println(account)
	//time.Sleep(10*time.Second)
	return &pb.Response{
		Code:    200,
		Message: "OK",
	}, nil
}

func StartGRPCServer() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		withServerUnaryInterceptor(),
	)
	pb.RegisterStatisticsServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}