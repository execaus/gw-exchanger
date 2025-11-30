package main

import (
	"fmt"
	"gw-exchanger/config"
	"gw-exchanger/internal/handler"
	gw_grpc "gw-exchanger/internal/pb/exchange"
	"gw-exchanger/internal/repository"
	"gw-exchanger/internal/service"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	cfg := config.LoadConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		zap.L().Fatal("failed to listen", zap.Error(err))
	}

	r, closeConnection := repository.NewRepository(&cfg.Database)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	grpcServer := grpc.NewServer()
	gw_grpc.RegisterExchangeServiceServer(grpcServer, h)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		zap.L().Info(fmt.Sprintf("gRPC server started on :%v", cfg.Server.Port))
		if err := grpcServer.Serve(lis); err != nil {
			zap.L().Error("failed to serve", zap.Error(err))
		}
	}()

	<-stop
	zap.L().Info("Shutting down gRPC server...")

	grpcServer.GracefulStop()
	closeConnection()
	zap.L().Info("gRPC server stopped")
}
