package asset

import (
		 "github.com/maheshDere/assetmanager/db"
)


type createAssetRequest struct {
	Asset db.Asset `json:"asset"`
}

type listAssetResponse struct {
	Assets []db.AssetResponse `json:"assets"`
}

type findByIDResponse struct {
	Asset db.AssetResponse `json:"asset"`
}
