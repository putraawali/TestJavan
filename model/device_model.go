package model

type Device struct {
	MemberDeviceID int    `gorm:"column:member_device_id;primaryKey"`
	MemberID       int    `gorm:"column:member_id"`
	DeviceType     string `gorm:"column:device_type"`
	DeviceToken    string `gorm:"column:device_token"`
}
