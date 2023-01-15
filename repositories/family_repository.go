package repositories

import (
	"context"
	"errors"
	"testjavan/helpers/errs"
	"testjavan/model"
	"time"

	"gorm.io/gorm"
)

type FamilyRepository interface {
	GetFamilyMemberByID(ctx context.Context, id int) (model.Family, error)
	UpdateFamilyMemberByID(ctx context.Context, data model.Family) error
	RemoveFamilyMemberByID(ctx context.Context, id int) error
}

type family struct {
	db *gorm.DB
}

func newFamilyRepository(db *gorm.DB) FamilyRepository {
	return &family{db: db}
}

func (f *family) GetFamilyMemberByID(ctx context.Context, id int) (model.Family, error) {
	var (
		result model.Family
		err    error
	)

	selectTable := `
		member_id,
		member_name,
		gender,
		created_at,
		updated_at,
		deleted_at
	`

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	db := f.db.Table("members")
	db = db.Select(selectTable)
	db = db.Where("member_id = ? AND deleted_at IS NULL", id).WithContext(ctx)
	if err = db.First(&result).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = errs.ErrRecordNotfound
		return result, err
	}

	return result, err
}

func (f *family) UpdateFamilyMemberByID(ctx context.Context, data model.Family) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return f.db.Table("members").Updates(data).WithContext(ctx).Error
}

func (f *family) RemoveFamilyMemberByID(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	now := time.Now()

	data := model.Family{
		MemberID:  id,
		DeletedAt: &now,
	}

	return f.db.Table("members").Updates(data).WithContext(ctx).Error
}
