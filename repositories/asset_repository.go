package repositories

import (
	"context"
	"testjavan/helpers/errs"
	"testjavan/model"
	"time"

	"gorm.io/gorm"
)

type AssetRepository interface {
	GetAssetMemberByID(ctx context.Context, id int) ([]model.Asset, error)
}

type asset struct {
	db *gorm.DB
}

func newAssetRepository(db *gorm.DB) AssetRepository {
	return &asset{db: db}
}

func (a *asset) GetAssetMemberByID(ctx context.Context, id int) ([]model.Asset, error) {
	var (
		result []model.Asset
		err    error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	selectTable := `
		a.asset_id,
		a.asset_name,
		a.created_at,
		a.updated_at,
		a.deleted_at
	`

	db := a.db.Table("member_assets ma")
	db = db.Select(selectTable)
	db = db.Joins("INNER JOIN assets a ON a.asset_id = ma.asset_id")
	db = db.Where("ma.member_id = ? AND a.deleted_at IS NULL", id).WithContext(ctx)
	err = db.Find(&result).Error
	if err != nil {
		return result, err
	}

	if len(result) == 0 {
		err = errs.ErrRecordNotfound
		return result, err
	}

	return result, err
}
