package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"net"
	"strconv"
	"context"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/go-hexa/balance/internal/core"
	"github.com/go-hexa/balance/internal/handlers/protobuf"
	"github.com/go-hexa/balance/internal/adapters/repository"
	"github.com/go-hexa/balance/internal/adapters/client"
)

var my_pod core.Pod
var my_setup core.Setup 

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("*** WARNING .env file NOT FOUND, using the os.env")
	}
	return os.Getenv(key)
}

func initSetup(){
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Erro Fatal")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				my_pod.Ip = ipnet.IP.String()
			}
		}
	}
	my_pod.Port = getEnvVariable("PORT")
	if len(my_pod.Port) == 0 {
		my_pod.Port = "50051"
	}
	my_pod.Name = getEnvVariable("NAME_POD")
	if len(my_pod.Name) == 0 {
		my_pod.Name = "pod_no_name"
	}
	my_pod.Host = getEnvVariable("HOST")
	if len(my_pod.Host) == 0 {
		my_pod.Host = "0.0.0.0"
	}
	
	my_pod.PID = strconv.Itoa(os.Getpid())
	
	my_setup.ResponseTime = 0
	my_setup.IsRandomTime = false

	my_pod.Setup = my_setup

	log.Printf("Setup : %v \n",my_pod)
}

func authInterceptor(ctx context.Context,req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
	log.Printf("-------------------------------------------- \n")
	log.Printf("-- authInterceptor", info.FullMethod)
	log.Printf("-- req : %v", req)

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "INTERNAL_SERVER_ERROR")
	}

	log.Printf("-- jwt : %v", headers["jwt"])
	if len(headers["jwt"]) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Not Authorized")
	}

	log.Printf("-------------------------------------------- \n")
	return handler(ctx, req) 
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Starting gRPC Server 1.0")

	initSetup()

	handler_cliente	:= client.NewGrpcAdapterClient()
	repo 			:= repository.NewMemKV()

	service 		:= core.NewService(repo, handler_cliente)

	handler			:= hdl_grpc.NewGrpcAdapter(service)

	// ------------------------------------
	// Load dummy data
	// for i:=0 ; i < 500; i++ {
	// 	id :=  strconv.Itoa(i)
	// 	acc := "acc-" + strconv.Itoa(i)
	// 	description := "description-"+ strconv.Itoa(i)

	// 	b := core.Balance{Id: id, Account: acc, Amount: 1, DateBalance: time.Now(), Description: description}
	// 	service.AddBalance(b)
	// }
	// ------------------------------------

	var hostname = my_pod.Host + ":" +  my_pod.Port
	lis, err := net.Listen("tcp",hostname)
	if err != nil{
		log.Fatalf("Failed to connect listener : %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	balancepb.RegisterBalanceServiceServer(s, handler)

	go func(){
		log.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to connect server : %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	log.Println("-> Stopping initialized")
	s.Stop()
	log.Println("--> Stopping gPRC Server")
	lis.Close()
	log.Println("---> Stop DONE !!!")
}