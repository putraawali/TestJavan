package usecase

import (
	"testjavan/repositories"
)

type Usecase struct {
	Family FamilyUsecase
	Asset  AssetUsecase
}

// NewUsecase: Init new usecase
func NewUsecase(repo *repositories.Repository) *Usecase {
	return &Usecase{
		Family: newFamilyUsecase(repo),
		Asset:  newAssetUsecase(repo),
	}
}
