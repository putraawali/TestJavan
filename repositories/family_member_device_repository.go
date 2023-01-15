package repositories

import (
	"context"
	"testjavan/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberDevice interface {
	UpsertDevice(ctx context.Context, data model.Device) error
	GetDevices(ctx context.Context) ([]model.Device, error)
}

type device struct {
	db *gorm.DB
}

func newDeviceRepository(db *gorm.DB) MemberDevice {
	return &device{db: db}
}

func (d *device) UpsertDevice(ctx context.Context, data model.Device) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	return d.db.Table("family_member_devices").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "device_token"}, {Name: "member_id"}},
			DoNothing: true,
		}).
		Create(&data).
		WithContext(ctx).
		Error
}

func (d *device) GetDevices(ctx context.Context) ([]model.Device, error) {
	var (
		result []model.Device
		err    error
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	err = d.db.Table("family_member_devices").Find(&result).WithContext(ctx).Error

	return result, err
}
