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
	// 	Caption : "sport",
	// 	Description : "sport is go for health",
	// 	Url: "abc.com",
	// 	AssetType: "image",
	// 	Source : "S3",
	// 	Owner :	"058ece9b-a1b2-4135-9b0d-d28fcb6aa169",
	// 	Private : true,
    //     Tags: []string{"cricket", "sport", "football"},
		
	// }
	assetReq := &pb.AssetReq{
		Id:" 7a6c3539-bf3c-4adf-8b0c-ac97977ee33f",
		Tag: "cricket",
		Owner :	"058ece9b-a1b2-4135-9b0d-d28fcb6aa169",
		Private : true,
		Video  : true, 
		Audio   :true,
		Image :true,
		Limit: 10,
		Offset: 0,		
	}
	client := pb.NewAssetServiceClient(conn)
	 //createAsset(client, asset)
	 
	listAsset(client, assetReq)
}
