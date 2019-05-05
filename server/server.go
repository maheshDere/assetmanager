package server

import (
	"fmt"
	"strconv"
    "net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/maheshDere/assetmanager/assetproto" 
	"github.com/maheshDere/assetmanager/config"
	"github.com/maheshDere/assetmanager/asset"
)

func StartAPIServer() {
	port := config.AppPort()
	p:=string(port)
	
	dependencies, err := initDependencies()
	if err != nil {
		panic(err)
	}
	
	lis, err := net.Listen("tcp", p)
	
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &asset.Server{})
	s.Serve(lis)
	
}
