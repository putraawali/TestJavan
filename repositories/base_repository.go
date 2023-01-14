package repositories

import "gorm.io/gorm"

type Repository struct {
	Family FamilyRepository
	Asset  AssetRepository
}

// NewRepository: Init new repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Family: newFamilyRepository(db),
		Asset:  newAssetRepository(db),
	}
}
