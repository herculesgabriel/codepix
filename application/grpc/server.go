package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/herculesgabriel/codepix/application/grpc/pb"
	"github.com/herculesgabriel/codepix/application/usecase"
	"github.com/herculesgabriel/codepix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	address := fmt.Sprintf("0.0.0.0:%d", port)
	reflection.Register(grpcServer)

	pixKeyRepository := repository.PixKeyRepositoryDB{DB: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: &pixKeyRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start listener", err)
	}

	log.Printf("listener has been started on port %d", port)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

}
