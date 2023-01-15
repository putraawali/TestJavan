package model

import (
	"time"
)

type Asset struct {
	AssetID   int        `gorm:"column:asset_id" json:"asset_id"`
	AssetName string     `gorm:"column:asset_name" json:"asset_name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at,autoUpdateTime:mili" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

type AssetWithPrice struct {
	AssetID   int        `json:"asset_id"`
	AssetName string     `json:"asset_name"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
