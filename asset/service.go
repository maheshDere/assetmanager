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
	asv *assetService
)

type Service interface {
	Dbconnection()
}

type assetService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type Server struct {
	}
	 
func (s *Server) CreateAsset(ctx context.Context, asset *CreateAssetReq) (*ListAssetResp, error) {
	fmt.Println("myserver",*s)
	tag := strings.Join(asset.tags,",")
	
	asset, err:= asv.store.CreateAsset(ctx, &db.Asset{
		Caption: asset.caption,
		Description: asset.description,
		AssetType: asset.asset_type,
		URL: asset.url,
		Source: asset.source,
		Owner: asset.owner,
		Private: asset.private,
        Tags: tag,
	})
	
	tags = strings.Split(asset.Tags, ",")
	
	return &pb.ListAssetResp{
		 Asset: Asset {
		 	id:   asset.ID,
		 	url:  asset.URL,
		 	asset_type: asset.AssetType,
		 	tags:        tags,
		 	caption:  asset.Caption,
		 	description: asset.Description,
		  },
		}, nil
}

func (s *Server) GetAsset(ctx context.Context, asset *pb.AssetReq) (*pb.ListAssetResp, error){
	return &pb.ListAssetResp{},nil
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
	fmt.Println(asv)
	asv = as
	fmt.Println(asv)
}















































// func (s *assetService) createAsset(ctx context.Context, request createAssetRequest) (assetResponse findByIDResponse, err error) {
// 	s.logger.Infof("Create asset request in service:  %#v",  request)
	
// 	if err != nil {
// 		s.logger.Errorw("Invalid request for create asset in service", "msg", err.Error(), "asset", request)
// 		return
// 	}	
	
	
// 	asset, err := s.store.CreateAsset(ctx, &db.Asset{
// 		URL: request.Asset.URL,
// 		Type: request.Asset.Type,
// 		IsPrivate: request.Asset.IsPrivate,
// 		Tags: request.Asset.Tags,
//         UserID: request.Asset.UserID,
// 	})

// 	if err != nil {
// 		s.logger.Errorw("Error creating asset in service", "err", err.Error())
// 		return
// 	}
//     assetResponse.Asset=asset
// 	return
// }

// func (s *assetService) listAssets(ctx context.Context, userID, assetType string) (response listAssetResponse, err error) {
// 	s.logger.Infof("List asset request ID and asset type : %s , %s ", userID, assetType)
	
// 	assets, err := asv.store.ListAssets(ctx, userID, assetType)
// 	if err != nil {
// 		s.logger.Errorw("Error listing assets", "err", err.Error())
// 		return
// 	}

// 	response.Assets = assets
// 	return
// }

// func (s *assetService) findByIDAsset(ctx context.Context, assetID string) (response findByIDResponse, err error) {
// 	s.logger.Infof("List asset request ID and  : %s ",assetID)
	
// 	response.Asset, err = s.store.FindByIDAsset(ctx, assetID)
// 	if err != nil {
// 		s.logger.Errorw("Error finding assets", "err", err.Error())
// 		return
// 	}
// 	return
// }


// func (s *assetService) listPublicAssets(ctx context.Context, assetType string) (response listAssetResponse, err error) {
// 	s.logger.Infof("List asset request ID and asset type : %s  ", assetType)
	
// 	assets, err := s.store.ListPublicAssets(ctx, assetType)
// 	if err != nil {
// 		s.logger.Errorw("Error listing assets", "err", err.Error())
// 		return
// 	}

// 	response.Assets = assets
// 	return
// }


// func (s *assetService) deleteAsset(ctx context.Context, id, userID string) (err error) {
// 	s.logger.Infof("Delete asset request ID: %s  ", id)
// 	err = s.store.DeleteAsset(ctx, id, userID)
// 	if err != nil {
// 		s.logger.Errorw("Error deleting asset: ", "msg", err.Error(), "id", id)
// 		return
// 	}
// 	return
// }



