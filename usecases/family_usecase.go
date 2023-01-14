package usecase

import (
	"context"
	"testjavan/repositories"
)

type FamilyUsecase interface {
	GetFamilyMemberByID(ctx context.Context, id int) error
}

type family struct {
	repo *repositories.Repository
}

func newFamilyUsecase(repo *repositories.Repository) FamilyUsecase {
	return &family{repo: repo}
}

func (f *family) GetFamilyMemberByID(ctx context.Context, id int) error {
	return nil
}
