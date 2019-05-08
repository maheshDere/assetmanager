package db

import (
	"strings"
	"context"
	"fmt"
	
	"github.com/lib/pq"
	"time"
	pb "github.com/maheshDere/assetmanager/assetproto"
)

const (
	createAssetQuery = `INSERT INTO asset ( caption, description, asset_type, url, source, owner, private, tags, created_at)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, url, asset_type, tags, caption ,description`

	listAssetsQuery = `SELECT id, url, asset_type, tags, caption, description FROM asset where `
)

type Asset struct {
	ID          string `db:"id"`
	Caption     string `db:"caption"`
	Description string `db:"description"`
	AssetType   string `db:"asset_type"`
	URL         string `db:"url"`
	Source      string `db:"source"`
	Owner       string `db:"owner"`
	Private     bool   `db:"private"`
	Tags        string `db:"tags"`
}

type AssetResponse struct {
	ID          string `db:"id"`
	URL         string `db:"url"`
	AssetType   string `db:"asset_type"`
	Tags        string `db:"tags"`
	Caption     string `db:"caption"`
	Description string `db:"description"`
}


func (s *store) CreateAsset(ctx context.Context, asset *Asset) (assetResponse AssetResponse, err error) {
	fmt.Println("asset:::", asset.AssetType)
	now := time.Now()
	err = s.db.GetContext(
		ctx,
		&assetResponse,
		createAssetQuery,
		asset.Caption,
		asset.Description,
		asset.AssetType,
		asset.URL,
		asset.Source,
		asset.Owner,
		asset.Private,
		asset.Tags,
		now,
	)
	fmt.Println("assetResp:::", assetResponse)

	return
}

func (s *store) GetAsset(ctx context.Context, assetReq *pb.AssetReq) ([]AssetResponse, error) {
	fmt.Println("I cam in GetAsset")
	var err error
	listAssetsQuery := listAssetsQuery
    

	var assets = make([]AssetResponse, 0)
	var assetsType = make([]string, 0)
	if assetReq.Video == true{
		assetsType = append(assetsType, "video")
	}
	if assetReq.Audio == true{
		assetsType =append(assetsType, "audio")
	}
	if assetReq.Image == true{
		assetsType =append(assetsType, "image")
	}
	fmt.Println("I cam in GetAsset")
	
	param := "{" + strings.Join(assetsType, ",") + "}"
	fmt.Println("Params", param)
	
	fmt.Println(assetReq.Id)


	if assetReq.Id != ""{
		fmt.Println("************************************************************")
		listAssetsQuery0 := listAssetsQuery + `string_to_array() && array[assetsType] `
		err = s.db.SelectContext(ctx, &assets, listAssetsQuery0, pq.Array([]string{"7a6c3539-bf3c-4adf-8b0c-ac97977ee33f","fb3a7d45-a55a-4ef9-a5e2-30a47207d409"}), pq.Array(assetsType))
	} else {
		if assetReq.Tag != "" && assetReq.Private == true {
        	listAssetsQuery1 := listAssetsQuery + `owner = $1 `	
			err = s.db.SelectContext(ctx, &assets, listAssetsQuery1, assetReq.Owner)		
		} else {
			listAssetsQuery2 := listAssetsQuery + ` tags = $1 AND asset_type = ANY ($3) and private = false`
			err = s.db.SelectContext(ctx, &assets, listAssetsQuery2, assetReq.Tag, param, assetReq.Private)
		}
	}
	
	return assets, err
}

func (s *store) DB() store {
	return *s
}
