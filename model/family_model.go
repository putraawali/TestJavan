package model

import (
	"time"
)

type Family struct {
	MemberID   int        `gorm:"column:member_id;primaryKey" json:"member_id"`
	MemberName string     `gorm:"column:member_name" json:"member_name"`
	Gender     string     `gorm:"column:gender" json:"gender"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime:mili;column:updated_at" json:"updated_at,omitempty"`
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
}

type FamilyAssets struct {
	MemberID   int              `json:"member_id"`
	MemberName string           `json:"member_name"`
	Gender     string           `json:"gender"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  *time.Time       `json:"updated_at,omitempty"`
	DeletedAt  *time.Time       `json:"deleted_at,omitempty"`
	Assets     []AssetWithPrice `json:"assets"`
}

type FamilyRequest struct {
	MemberName string `json:"member_name"`
	Gender     string `json:"gender"`
}
