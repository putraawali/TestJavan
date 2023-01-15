package repositories

import (
	"context"
	"testjavan/model"
	"time"

	"gorm.io/gorm"
)

type FamilyAssetRepository interface {
	AddAssetToFamilyMemberByMemberID(ctx context.Context, data model.MemberAsset) error
	UpdateFamilyAssetByMemberIDandAssetID(ctx context.Context, data model.MemberAsset) error
}

type familyAsset struct {
	db *gorm.DB
}

func newFamilyAsset(db *gorm.DB) FamilyAssetRepository {
	return &familyAsset{db: db}
}

func (fa *familyAsset) AddAssetToFamilyMemberByMemberID(ctx context.Context, data model.MemberAsset) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	return fa.db.Table("member_assets").Create(&data).WithContext(ctx).Error
}

func (fa *familyAsset) UpdateFamilyAssetByMemberIDandAssetID(ctx context.Context, data model.MemberAsset) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	return fa.db.Table("member_assets").Where("member_id = ? AND asset_id = ?", data.MemberID, data.AssetID).Updates(data).WithContext(ctx).Error
}
