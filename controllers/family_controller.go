package controllers

import "github.com/labstack/echo/v4"

func (ct *controllerImpl) GetFamilyMemberByID(c echo.Context) error {

	return c.JSON(200, "OK")
}
