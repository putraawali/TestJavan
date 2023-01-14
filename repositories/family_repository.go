package repositories

import (
	"context"

	"gorm.io/gorm"
)

type FamilyRepository interface {
	GetFamilyMemberByID(ctx context.Context, id int) error
}

type family struct {
	db *gorm.DB
}

func newFamilyRepository(db *gorm.DB) FamilyRepository {
	return &family{db: db}
}

func (f *family) GetFamilyMemberByID(ctx context.Context, id int) error {
	var (
		err error
	)

	return err
}
