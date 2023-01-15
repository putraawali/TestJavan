package usecase

import (
	"testjavan/repositories"
)

type Usecase struct {
	Family FamilyUsecase
	Asset  AssetUsecase
	Notif  NotificationUsecase
}

// NewUsecase: Init new usecase
func NewUsecase(repo *repositories.Repository) *Usecase {
	return &Usecase{
		Family: newFamilyUsecase(repo),
		Asset:  newAssetUsecase(repo),
		Notif:  newNotificationUsecase(repo),
	}
}
