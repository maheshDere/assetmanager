package main

import (
	"log"
    "fmt"
	pb "github.com/maheshDere/assetmanager/assetproto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func createAsset(client pb.AssetServiceClient, asset *pb.CreateAssetReq) {
	resp, err := client.CreateAsset(context.Background(), asset)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	fmt.Println("resp", resp)
}

func listAsset(client pb.AssetServiceClient, asset *pb.AssetReq) {
	resp, err := client.GetAsset(context.Background(), asset)
	fmt.Println("Hey")
	if err != nil {
		log.Fatalf("Could not list Customer: %v", err)
	}
	fmt.Println("resp", resp)
}
 
func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	
    // asset := &pb.CreateAssetReq{
	// 	Caption : "Study",
	// 	Description : "must study",
	// 	Url: "s.com",
	// 	AssetType: "audio",
	// 	Source : "S3",
	// 	Owner :	"058ece9b-a1b2-4135-9b0d-d28fcb6aa162",
	// 	Private : false,
    //     Tags: []string{"study", "student", "teacher"},
		
	// }
	assetReq := &pb.AssetReq{
		Id:"",
		Tag: "",
		Owner :	"058ece9b-a1b2-4135-9b0d-d28fcb6aa169",
		Private : false,
		Video  : true, 
		Audio   :false,
		Image :true,
		Limit: 10,
		Offset: 0,		
	}
	client := pb.NewAssetServiceClient(conn)
	 //createAsset(client, asset)
	 
	listAsset(client, assetReq)
}
