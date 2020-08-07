package grpc_service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"oss-helper/grpc_service/contract"
	"oss-helper/internal/config"
)

func Run(cfg config.Config) {

	s := grpc.NewServer() //创建gRPC服务

	contract.RegisterStorageServer(s, &StorageService{})

	// 在gRPC服务器上注册反射服务
	reflection.Register(s)

	lis, err := net.Listen("tcp", cfg.Host)
	if err != nil {
		panic(err)
	}

	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
