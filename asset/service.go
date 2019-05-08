package asset

import (
	"fmt"
	"context"
	"strings"
	pb "github.com/maheshDere/assetmanager/assetproto"

	"github.com/maheshDere/assetmanager/db"
	"go.uber.org/zap"
)

var(
	asc *assetService
)

type Service interface {
	Dbconnection()
}

type assetService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type Server struct {}

	
func (s *Server) CreateAsset(ctx context.Context, asset *pb.CreateAssetReq) (*pb.ListAssetResp, error) {
	fmt.Println("myserver",*s)
	tag := strings.Join(asset.Tags,",")
	
	a, err:= asc.store.CreateAsset(ctx, &db.Asset{
		Caption: asset.Caption,
		Description: asset.Description,
		AssetType: asset.AssetType,
		URL: asset.Url,
		Source: asset.Source,
		Owner: asset.Owner,
		Private: asset.Private,
    Tags: tag,
	})
	
	tags := strings.Split(a.Tags, ",")
  if err != nil{
    fmt.Println("error ::",err)
   }
	
	 return &pb.ListAssetResp{
		Assets: []*pb.Asset{{
			Id:   a.ID,
			Url:  a.URL,
			AssetType: a.AssetType,
			Tags:        tags,
			Description: a.Description,
			Caption:  a.Caption,
		  },},
		}, nil
}

func (s *Server) GetAsset(ctx context.Context, asset *pb.AssetReq) (*pb.ListAssetResp,error){
	var assetResp = pb.ListAssetResp{}
	assets, err := asc.store.GetAsset(ctx, asset)
	if err != nil {
	fmt.Println("error::in list assets in asset service", err )
	}
	for _,a := range assets{
	 tags := strings.Split(a.Tags, ",")
	 as:=pb.Asset{
	 Url : a.URL,
	 AssetType: a.AssetType,
	 Id : a.ID,
	 Tags : tags,
	 Caption: a.Caption,
	 Description: a.Description,
	}
	assetResp.Assets = append(assetResp.Assets ,&as) 
	} 
	return &assetResp,err
}
func (s *Server) GetMediaLibrary(ctx context.Context, asset *pb.AssetReq) (*pb.ListAssetResp, error){
	return &pb.ListAssetResp{},nil
} 

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &assetService{
		store:  s,
		logger: l,
	}
}

func(as *assetService) Dbconnection(){
	asc = as	
}

