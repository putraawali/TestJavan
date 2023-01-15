package controllers

import (
	usecase "testjavan/usecases"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetFamilyMemberByID(c echo.Context) error
	UpdateFamilyMemberByID(c echo.Context) error
	RemoveFamilyMemberByID(c echo.Context) error
	CreateFamilyMember(c echo.Context) error
	GetAssetByID(c echo.Context) error
	UpdateAssetByID(c echo.Context) error
	CreateAsset(c echo.Context) error
	DeleteAssetByID(c echo.Context) error
	AddAssetToFamilyMember(c echo.Context) error
	RemoveAssetFromFamilyMember(c echo.Context) error
}

type controllerImpl struct {
	usecase *usecase.Usecase
}

func NewController(e *echo.Echo, usecase *usecase.Usecase) {
	ctrl := &controllerImpl{usecase: usecase}

	familyRouter(e, ctrl)
	assetRouter(e, ctrl)
}

func familyRouter(e *echo.Echo, controller *controllerImpl) {
	g := e.Group("/family")

	g.GET("/member/:id", controller.GetFamilyMemberByID)
	g.POST("/member", controller.CreateFamilyMember)
	g.PUT("/member/:id", controller.UpdateFamilyMemberByID)
	g.DELETE("/member/:id", controller.RemoveFamilyMemberByID)

	// Add new asset to family member
	g.POST("/member/:id/asset", controller.AddAssetToFamilyMember)
	g.DELETE("/member/:id/asset", controller.RemoveAssetFromFamilyMember)
}

func assetRouter(e *echo.Echo, controller *controllerImpl) {
	a := e.Group("/asset")

	a.GET("/:id", controller.GetAssetByID)
	a.POST("", controller.CreateAsset)
	a.PUT("/:id", controller.UpdateAssetByID)
	a.DELETE("/:id", controller.DeleteAssetByID)
}
