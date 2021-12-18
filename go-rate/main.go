package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"context"

	"google.golang.org/grpc"

	proto "github.com/go-hexa/proto-shared/generated/go/rate"

)

type server struct{}

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Rate gRPC server")

	lis, err := net.Listen("tcp","0.0.0.0:60051")
	if err != nil{
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterRateServiceServer(s , &server{})

	go func(){
		log.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM )
	<-ch
	log.Println("Stopping server")
	s.Stop()
	log.Println("Stopping listener")
	lis.Close()
	log.Println("Stopped Done !!!")

}

func(*server) GetRateAccount(ctx context.Context, req *proto.RateRequest) (*proto.RateResponse, error) {

	rate := &proto.Rate{}
	rate.Rate = 12
	rate.Description = "Taxa do Cliente"

	res := &proto.RateResponse {
		Rate: rate,
	}

	return res, nil
}