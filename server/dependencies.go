package server

import (
	"github.com/maheshDere/assetmanager/app"
	"github.com/maheshDere/assetmanager/db"
	"github.com/maheshDere/assetmanager/asset"
)

type dependencies struct {
	AssetService        asset.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB, logger)
	assetService := asset.NewService(dbStore, logger)
	assetService.Dbconnection()
	
	return dependencies{
		AssetService:        assetService,
	}, nil
}
