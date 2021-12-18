package main

import (
	"fmt"
	"log"
	"context"
	"time"
	"flag"
	"strconv"
	"os"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-hexa/go-balance/internal/handlers/protobuf"
)

func main(){
	fmt.Println("Balance gRPC client")

	port := flag.String("port","","")
	flag.Parse()
	flag.VisitAll(func (f *flag.Flag) {
		if f.Value.String()=="" {
			fmt.Printf("A flag -%v não foi informado \n", f.Name )
			os.Exit(1)
		}
	})

	var host = "127.0.0.1:" + *port
	cc, err := grpc.Dial(host,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Failed to connect %v", err)
	}
	defer cc.Close()

	c := balancepb.NewBalanceServiceClient(cc)

	ts := timestamppb.Now()
	fmt.Printf("ts Timestamppb.AsTime() : %s\n", ts.AsTime().String())
	fmt.Println("start LoadBalance data")
	for i:=0 ; i < 500; i++ {

		id :=  strconv.Itoa(i)
		acc := "acc-" + strconv.Itoa(i)
		description := "description-"+ strconv.Itoa(i)
		
		b := balancepb.Balance{Id: id, 
			Account: acc, 
			Amount: 1, 
			DateBalance: ts, 
			Description: description,
		}
		
		req := &balancepb.AddBalanceRequest {
			Balance: &b,
		}
		LoadBalance(c , req ,15 * time.Second)
	}
	fmt.Println("end LoadBalance data")

	fmt.Println("-----------------------------")
	GetBalance(c , 15 * time.Second)
	fmt.Println("-----------------------------")
	ListBalance(c , 15 * time.Second)
	fmt.Println("-----------------------------")
	//AddBalance(c , 15 * time.Second)

	for i:=0; i < 3600; i++ {
		ListBalance(c , 15 * time.Second)
		time.Sleep(time.Second * time.Duration(1))
	}

}

func GetBalance(c balancepb.BalanceServiceClient, timeout time.Duration){
	fmt.Println("GetBalance")

	req := &balancepb.GetBalanceRequest {
		Id: "3",
	}

	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.GetBalance(ctx, req) 
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded  {
				log.Printf("---> Timeout, deadline exceeded")
			} else {
				log.Printf("Error unexpected : %v \n", statusErr)
			}
		} else {
			log.Printf("Error to call PRC : %v", err)
		}
	} else {
		fmt.Println(res)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func ListBalance(c balancepb.BalanceServiceClient, timeout time.Duration){
	fmt.Println("ListBalance")

	req := &balancepb.ListBalanceRequest {}

	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.ListBalance(ctx, req) 
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded  {
				log.Printf("---> Timeout, deadline exceeded")
			} else {
				log.Printf("Error unexpected : %v \n", statusErr)
			}
		} else {
			log.Printf("Error to call PRC : %v", err)
		}
	} else {
		fmt.Println(res)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func AddBalance(c balancepb.BalanceServiceClient, timeout time.Duration){
	fmt.Println("AddBalance")

	req := &balancepb.AddBalanceRequest {
		Balance: &balancepb.Balance {
			Id: "9999",
			Account: "acc-9999",
			Amount: 1, 
			DateBalance: timestamppb.Now(),
			Description: "Description-9999",
		},
	}

	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.AddBalance(ctx, req) 
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded  {
				log.Printf("---> Timeout, deadline exceeded")
			} else {
				log.Printf("Error unexpected : %v \n", statusErr)
			}
		} else {
			log.Printf("Error to call PRC : %v", err)
		}
	} else {
		fmt.Println(res)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}

func LoadBalance(c balancepb.BalanceServiceClient, req *balancepb.AddBalanceRequest, timeout time.Duration){
	header := metadata.New(map[string]string{"accept_language": "pt-BR", "jwt":"cookie"})
	ctx, cancel := context.WithTimeout(context.Background(), timeout) 
	ctx = metadata.NewOutgoingContext(ctx, header)
	defer cancel()

	res, err := c.AddBalance(ctx, req) 
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok{
			if statusErr.Code() == codes.DeadlineExceeded  {
				log.Printf("---> Timeout, deadline exceeded")
			} else {
				log.Printf("Error unexpected : %v \n", statusErr)
			}
		} else {
			log.Printf("Error to call PRC : %v", err)
		}
	} else {
		fmt.Println(res)
		//fmt.Printf("Timestamppb.AsTime() : %s\n", res.Balance.DateBalance.AsTime().String())
	}
}