package model

import (
	"time"
)

type Asset struct {
	AssetID   int        `gorm:"column:asset_id;primaryKey" json:"asset_id"`
	AssetName string     `gorm:"column:asset_name" json:"asset_name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type AssetWithPrice struct {
	AssetID   int        `json:"asset_id"`
	AssetName string     `json:"asset_name"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type AssetRequest struct {
	AssetID   int    `json:"asset_id"`
	AssetName string `json:"asset_name"`
}
