package server

import (
	"net"
	"log"
	"fmt"
	"strconv"
	"google.golang.org/grpc"
	pb "github.com/maheshDere/assetmanager/assetproto" 
	"github.com/maheshDere/assetmanager/config"
	"github.com/maheshDere/assetmanager/asset"
)

func StartAPIServer() {
	port := config.AppPort()
	
	p := strconv.Itoa(port)
	
	fmt.Println("p",p)
	
	_, err := initDependencies()
	if err != nil {
		panic(err)
	}
	
	server := asset.Server{}

	lis, err := net.Listen("tcp",":"+p)
	
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterAssetServiceServer(s, &server)
	s.Serve(lis)
	
}
