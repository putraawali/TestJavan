package repositories

import (
	"context"
	"errors"
	"testjavan/helpers/errs"
	"testjavan/model"
	"time"

	"gorm.io/gorm"
)

type AssetRepository interface {
	GetAssetMemberByID(ctx context.Context, id int) ([]model.Asset, error)
	GetAssetByID(ctx context.Context, id int) (model.Asset, error)
	UpdateAssetByID(ctx context.Context, data model.Asset) error
	CreateAsset(ctx context.Context, data model.Asset) (model.Asset, error)
	GetAssetByAssetName(ctx context.Context, name string) (model.Asset, error)
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

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
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
	db = db.Where("ma.member_id = ? AND a.deleted_at IS NULL AND ma.deleted_at IS NULL", id).WithContext(ctx)
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

func (a *asset) GetAssetByID(ctx context.Context, id int) (model.Asset, error) {
	var (
		result model.Asset
		err    error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	selectTable := `
		asset_id,
		asset_name,
		created_at,
		updated_at,
		deleted_at
	`

	db := a.db.Table("assets")
	db = db.Select(selectTable)
	db = db.First(&result).WithContext(ctx)

	if err = db.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = errs.ErrRecordNotfound
		return result, err
	} else if err != nil {
		return result, err
	}

	return result, err
}

func (a *asset) UpdateAssetByID(ctx context.Context, data model.Asset) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	return a.db.Table("assets").Updates(data).WithContext(ctx).Error
}

func (a *asset) CreateAsset(ctx context.Context, data model.Asset) (model.Asset, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	err := a.db.Table("assets").Create(&data).WithContext(ctx).Error

	return data, err
}

func (a *asset) GetAssetByAssetName(ctx context.Context, name string) (model.Asset, error) {
	var (
		result model.Asset
		err    error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	selectTable := `
		asset_id,
		asset_name,
		created_at,
		updated_at,
		deleted_at
	`

	db := a.db.Table("assets")
	db = db.Select(selectTable)
	db = db.First(&result, "asset_name = ?", name).WithContext(ctx)

	if err = db.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = errs.ErrRecordNotfound
		return result, err
	} else if err != nil {
		return result, err
	}

	return result, err
}
