package db

import (	
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
	Limit       int32
	Offset      int32
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
	var assetsType [3]string
	fmt.Println("assets type", assetsType)
	if assetReq.Video == true{
		assetsType[0] = "video"
	}
	if assetReq.Audio == true{
		assetsType[1] = "audio"
	}
	if assetReq.Image == true{
		assetsType[2] = "image"
	}

	fmt.Println("assetType", assetsType)
	if assetReq.Id != ""{
		listAssetsQuery0 := listAssetsQuery + `id = $1 order by created_at limit $2 offset $3` 
        err = s.db.SelectContext(ctx, &assets, listAssetsQuery0, assetReq.Id, assetReq.Limit,assetReq.Offset )
		} else {
		if assetReq.Tag != "" && assetReq.Private == true {
			fmt.Println("inside else if")
        	listAssetsQuery1 := listAssetsQuery + `owner = $1 and asset_type = ANY($2) or string_to_array(tags, ',') && array['`+assetReq.Tag+`'] order by created_at limit $3 offset $4`	
			err = s.db.SelectContext(ctx, &assets, listAssetsQuery1, assetReq.Owner, pq.Array(assetsType), assetReq.Limit, assetReq.Offset )		
		} else {
			fmt.Println("inside else")
			fmt.Println("tags",assetReq.Tag)
			listAssetsQuery2 := listAssetsQuery + `asset_type = ANY($1)  and private = false or string_to_array(tags, ',') && array['`+assetReq.Tag+`'] order by created_at limit  $2 offset $3`
			err = s.db.SelectContext(ctx, &assets, listAssetsQuery2, pq.Array(assetsType), assetReq.Limit,assetReq.Offset)
		}
	}
	
	return assets, err
}

func (s *store) DB() store {
	return *s
}

