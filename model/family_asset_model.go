package model

import "time"

type MemberAsset struct {
	MemberAssetID int        `gorm:"column:member_asset_id;primaryKey"`
	MemberID      int        `gorm:"column:member_id"`
	AssetID       int        `gorm:"column:asset_id"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at"`
}
