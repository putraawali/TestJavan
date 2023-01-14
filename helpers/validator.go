package helpers

import (
	"testjavan/helpers/errs"

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
