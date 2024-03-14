package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/pandishpancheta/tokenization-service/pkg/config"
	tokenization "github.com/pandishpancheta/tokenization-service/pkg/pb"
	"github.com/pandishpancheta/tokenization-service/pkg/service"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Starting the service...")

	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", "localhost:"+cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	t := service.NewTokenizationService(*cfg)

	s := grpc.NewServer()

	log.Println("Registering tokenization service...")

	tokenization.RegisterTokenizationServiceServer(s, t)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
