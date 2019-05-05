package db

import (
	
	"context"
	"time"
)

const (
		createAssetQuery = `INSERT INTO asset ( caption, description, asset_type, url, source, owner, private, tags, created_at)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id, url, asset_type, tags, caption ,description`
	
	listAssetsQuery = `SELECT id, url, type FROM assets WHERE user_id = $1 AND type = $2  ORDER BY
	created_at DESC;
   `
   listPublicAssetsQuery = `SELECT id, url, type FROM assets WHERE is_private = false AND type = $1  ORDER BY
   created_at DESC;
  `	
  FindAssetByIDQuery = `SELECT id, url, type FROM assets WHERE id = $1;`	
  deleteAssetByIDQuery = `DELETE FROM assets WHERE id = $1 AND user_id = $2`
	
)

 
type Asset struct {
	ID  string           `db:"id"`
	Caption string       `db:"caption"`
	Description string   `db:"description"`
	AssetType string     `db:"asset_type"`
	URL  string          `db:"url"`
	Source string        `db:"source"`
	Owner string         `db:"owner"`
	Private bool         `db:"private"`
	Tags string          `db:"tags"`
}


type AssetResponse struct {
	ID          string      	
	URL 		    string    	
	Type    	  string    	
	Tags	      string
	Caption     string
	Description string
}

func (s *store) CreateAsset(ctx context.Context, asset *Asset) (assetResponse AssetResponse, err error) {
		now := time.Now()
		err = s.db.GetContext(
				ctx,
				&assetResponse,
				createAssetQuery,
				asset.Caption,	
				asset.AssetType,		
				asset.URL,
				asset.Source,
				asset.Owner,
				asset.Private,
				asset.Tags,
				now,			
			)
		return 
	}
	
	









































// type Asset struct {
// 	ID          string    	`db:"id" json:"id"`
// 	URL 		string    	`db:"url" json:"url"`
// 	Type    	string    	`db:"type" json:"type"`
// 	IsPrivate   bool   		`db:"is_private" json:"is_private"`
// 	Tags        []string 	`db:"tags" json:"tags"`
// 	CreatedAt   time.Time 	`db:"created_at" json:"created_at"`	
//   UserID      string    	`db:"user_id" json:"user_id"`
// }

// type AssetResponse struct {
// 	ID          string    	`json:"id"`
// 	URL 		string    	`json:"url"`
// 	Type    	string    	`json:"type"`	
// }


// func (s *store) CreateAsset(ctx context.Context, asset *Assett) (assetResponse AssetResponse, err error) {
// 	now := time.Now()
	 
		
// 	err = s.db.GetContext(
// 			ctx,
// 			&assetResponse,
// 			createAssetQuery,
// 			asset.Type,	
// 			asset.URL,		
// 			asset.IsPrivate,
// 			tags,
// 			now,
// 			asset.UserID,		
// 		)
// 	return 
// }

// func (s *store) ListAssets(ctx context.Context, userID, assetType string) ([]AssetResponse, error) {
// 	var assets = make([]AssetResponse, 0)

// 	err :=  s.db.SelectContext(ctx, &assets, listAssetsQuery, userID, assetType )

// 	return assets, err
// }

// func (s *store) FindByIDAsset(ctx context.Context, assetID  string) (assetResponse AssetResponse,err error){
// 	fmt.Println("Id",assetID)
// 	err =  s.db.SelectContext(ctx, &assetResponse, FindAssetByIDQuery, assetID )

// 	return 
// }

// func (s *store) ListPublicAssets(ctx context.Context, assetType string) ([]AssetResponse, error) {
// 	var assets = make([]AssetResponse, 0)

// 	err :=  s.db.SelectContext(ctx, &assets, listPublicAssetsQuery, assetType )

// 	return assets, err
// }

// func (s *store) DeleteAsset(ctx context.Context, id, userID string) (err error){	 
// 		_, err = s.db.Exec(deleteAssetByIDQuery, id, userID)
// 		return err	
// }

func (s *store) DB()(store){
 return *s
}