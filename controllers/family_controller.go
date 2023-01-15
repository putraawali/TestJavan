package controllers

import (
	"context"
	"strconv"
	"testjavan/helpers"
	"testjavan/helpers/errs"
	"testjavan/model"

	"github.com/labstack/echo/v4"
)

func (ct *controllerImpl) GetFamilyMemberByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	result, err := ct.usecase.Family.GetFamilyMemberByID(ctx, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Data = result
	response.Status = "success"

	return c.JSON(200, response)
}

func (ct *controllerImpl) UpdateFamilyMemberByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	req := model.FamilyRequest{}
	err = c.Bind(&req)
	if err != nil {
		return errs.Wrap(c, err)
	}

	if req.MemberName == "" || req.Gender == "" {
		err = errs.ErrMissingParameter
		return errs.Wrap(c, err)
	}

	gender := []string{"F", "M"}
	count := 0
	for _, g := range gender {
		if g != req.Gender {
			count++
		}
	}

	if count == 2 {
		err = errs.ErrInvalidParameter
		return errs.Wrap(c, err)
	}

	err = ct.usecase.Family.UpdateFamilyMemberByID(ctx, req, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	return c.JSON(200, response)
}

func (ct *controllerImpl) RemoveFamilyMemberByID(c echo.Context) error {
	var (
		response = model.Return{}
		ctx      = context.WithValue(context.Background(), helpers.RequestIDKey, c.Request().Header.Get(string(helpers.RequestIDKey)))
	)

	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return errs.Wrap(c, errs.ErrInvalidParameter)
	}

	err = ct.usecase.Family.RemoveFamilyMemberByID(ctx, id)
	if err != nil {
		return errs.Wrap(c, err)
	}

	response.Status = "success"

	return c.JSON(200, response)
}
