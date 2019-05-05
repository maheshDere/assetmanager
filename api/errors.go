package api

import (
	"errors"
)

const (
	ErrorForeignKeyViolationpg = "23503"
)

var (
	ErrJSONFormat = errors.New("Invalid request")
	//User
	ErrInvalidUser    		= errors.New("invalid user credential")	
	//Asset
	ErrEmptyAssetURL         			= errors.New("Asset URL must be present")
	ErrEmptyAssetType         			= errors.New("Asset type must be present")	
	ErrInvalidAssetID                	= errors.New("invalid asset credential")
	ErrInvalidAssetType                	= errors.New("invalid asset type")
	ErrInvalidAssetStatus               = errors.New("invalid asset status")
)
