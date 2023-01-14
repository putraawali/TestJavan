package repositories

import (
	"context"

	"gorm.io/gorm"
)

type AssetRepository interface {
	GetAssetMemberByID(ctx context.Context, id int) error
}

type asset struct {
	db *gorm.DB
}

func newAssetRepository(db *gorm.DB) AssetRepository {
	return &asset{db: db}
}

func (f *asset) GetAssetMemberByID(ctx context.Context, id int) error {
	var (
		err error
	)

	return err
}
