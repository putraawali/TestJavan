package controllers

import (
	usecase "testjavan/usecases"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetFamilyMemberByID(c echo.Context) error
	UpdateFamilyMemberByID(c echo.Context) error
	RemoveFamilyMemberByID(c echo.Context) error
}

type controllerImpl struct {
	usecase *usecase.Usecase
}

func NewController(e *echo.Echo, usecase *usecase.Usecase) {
	ctrl := &controllerImpl{usecase: usecase}

	familyRouter(e, ctrl)
}

func familyRouter(e *echo.Echo, controller *controllerImpl) {
	g := e.Group("/family")

	g.GET("/member/:id", controller.GetFamilyMemberByID)
	g.PUT("/member/:id", controller.UpdateFamilyMemberByID)
	g.DELETE("/member/:id", controller.RemoveFamilyMemberByID)
}
