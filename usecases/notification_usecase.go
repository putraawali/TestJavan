package usecase

import (
	"context"
	"fmt"
	"testjavan/repositories"
)

type NotificationUsecase interface {
	SendNotif(ctx context.Context, notifType string)
}

type notification struct {
	repo *repositories.Repository
}

func newNotificationUsecase(repo *repositories.Repository) NotificationUsecase {
	return &notification{repo: repo}
}

func (n *notification) SendNotif(ctx context.Context, notifType string) {
	devices, err := n.repo.Device.GetDevices(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, device := range devices {
		if device.DeviceType == "desktop" {
			//Sent notif using desktop notif services
			continue
		}

		if device.DeviceType == "android" {
			//Sent notif using android notification services
			continue
		}

		if device.DeviceToken == "ios" {
			//Sent notif using ios notification services
			continue
		}
	}
}
