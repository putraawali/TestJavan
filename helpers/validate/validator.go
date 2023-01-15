package validate

import (
	"context"
	"strconv"
	"testjavan/helpers/errs"
	"testjavan/model"
	"testjavan/repositories"

	"github.com/labstack/echo/v4"
)

func ValidateContentType() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if contentType := c.Request().Header.Get("Content-Type"); contentType != "application/json" {
				return errs.Wrap(c, errs.ErrContentType)
			}

			return next(c)
		}
	}
}

func ValidateUserDeviceToken(repo *repositories.Repository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header

			skipper := header.Get("Skipper")
			if skipper != "tcp-skipper" {
				deviceType := header.Get("Device-Type")
				deviceToken := header.Get("Device-Token")
				if deviceType == "" || deviceToken == "" {
					return errs.Wrap(c, errs.ErrMissingDevice)
				}

				memberID := header.Get("Member-ID")
				if memberID == "" {
					return errs.Wrap(c, errs.ErrMissingMemberID)
				}

				memberIDInt, _ := strconv.Atoi(memberID)
				go repo.Device.UpsertDevice(context.Background(), model.Device{
					MemberID:    memberIDInt,
					DeviceType:  deviceType,
					DeviceToken: deviceToken,
				})
			}

			return next(c)
		}
	}
}
