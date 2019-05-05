package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type StorerMock struct {
	mock.Mock
}

func (m *StorerMock) CreateAsset(ctx context.Context, asset *Asset) (assetResponse AssetResponse, err error) {
	args := m.Called(ctx, asset)
	return assetResponse, args.Error(1)
}

func (m *StorerMock) ListAssets(ctx context.Context, userID,AssetType string) (assets []Asset, err error) {
	args := m.Called(ctx, userID)
	assets, _ = args.Get(0).([]Asset)
	return assets, args.Error(1)
}

func (m *StorerMock) ListPublicAssets(ctx context.Context, AssetType string) (assets []Asset, err error) {
	args := m.Called(ctx, AssetType)
	assets, _ = args.Get(0).([]Asset)
	return assets, args.Error(1)
}

func (m *StorerMock) DeleteDream(ctx context.Context, id, userID string) (err error) {
	args := m.Called(ctx, id, userID)
	return args.Error(0)
}
func (m *StorerMock) UpdateDream(ctx context.Context, dream *Asset, id, userID string) (err error) {
	args := m.Called(ctx, dream, id, userID)
	return args.Error(0)
}

