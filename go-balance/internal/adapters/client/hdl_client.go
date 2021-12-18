package client

import(
	"log"
//	"context"

	"github.com/go-hexa/balance/internal/core"
	"google.golang.org/grpc"
	// "github/snipet-go-hexa/balance/internal/handlers/protobuf"
	// "google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/grpc/status"
	// "google.golang.org/grpc/codes"
)

type GrpcAdapterClient struct {
	cliente core.BalanceClientPort
}

func NewGrpcAdapterClient() *GrpcAdapterClient {
	return &GrpcAdapterClient{	}
}

func (g *GrpcAdapterClient) GetRate(account string) (int32, error) {
	log.Printf("GetRate")
	
	log.Printf("--------------------------------------")
	log.Printf("- GetRate Calling another Service !!!!")
	log.Printf("--------------------------------------")

	var host = "127.0.0.1:60051" 
	cc, err := grpc.Dial(host,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Failed to connect %v", err)
	}
	defer cc.Close()

	//c := balancepb.NewBalanceServiceClient(cc)

	return 3, nil
}