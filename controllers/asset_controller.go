package controllers

import (
	"context"
	"net/http"
	"strconv"
	"testjavan/helpers"
	"testjavan/helpers/errs"
	"testjavan/model"

	"github.com/labstack/echo/v4"
)

func (ct *controllerImpl) GetAssetByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
		err      error
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	response.Data, err = ct.usecase.Asset.GetAssetByID(ctx, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	return c.JSON(http.StatusOK, response)
}

func (ct *controllerImpl) UpdateAssetByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
		err      error
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	req := model.AssetRequest{}
	err = c.Bind(&req)
	if err != nil {
		return errs.Wrap(c, err)
	}

	if req.AssetName == "" {
		return errs.Wrap(c, errs.ErrMissingParameter)
	}

	err = ct.usecase.Asset.UpdateAssetByID(ctx, req, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	go ct.usecase.Notif.SendNotif(ctx, "update")

	return c.JSON(http.StatusOK, response)
}

func (ct *controllerImpl) CreateAsset(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
		err      error
	)

	req := model.AssetRequest{}
	err = c.Bind(&req)
	if err != nil {
		return errs.Wrap(c, err)
	}

	if req.AssetName == "" {
		err = errs.ErrInvalidParameter
		return errs.Wrap(c, err)
	}

	err = ct.usecase.Asset.CreateAsset(ctx, req)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	go ct.usecase.Notif.SendNotif(ctx, "create")

	return c.JSON(http.StatusCreated, response)
}

func (ct *controllerImpl) DeleteAssetByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
		err      error
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	err = ct.usecase.Asset.DeleteAssetByID(ctx, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	go ct.usecase.Notif.SendNotif(ctx, "remove")

	return c.JSON(http.StatusCreated, response)
}
