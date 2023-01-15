package errs

import (
	"errors"
	"net/http"
	"testjavan/model"

	"github.com/labstack/echo/v4"
)

var (
	ErrMethodNotAllowed = errors.New("error: method is not allowed")
	ErrContentType      = errors.New("error: Content-Type must be application/json")
	ErrMissingParameter = errors.New("error: missing parameter")
	ErrInvalidParameter = errors.New("error: invalid parameter")
	ErrRecordNotfound   = errors.New("error: data not found")
	ErrMissingDevice    = errors.New("error: Device-Type and Device-Token is mandatory")
	ErrMissingMemberID  = errors.New("error: Member-ID is mandatory")
)

var ErrHttpStatusMap = map[string]int{
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrContentType.Error():      http.StatusBadRequest,
	ErrMissingParameter.Error(): http.StatusBadRequest,
	ErrInvalidParameter.Error(): http.StatusBadRequest,
	ErrMissingDevice.Error():    http.StatusBadRequest,
	ErrMissingMemberID.Error():  http.StatusBadRequest,
	ErrRecordNotfound.Error():   http.StatusNotFound,
}

func Wrap(c echo.Context, err error) error {
	msg := err.Error()
	code := ErrHttpStatusMap[msg]

	if code == 0 {
		code = http.StatusInternalServerError
		msg = "error: internal server error"
	}

	errView := model.Error{
		Message:    msg,
		StatusCode: code,
	}

	data := map[string]interface{}{}

	return c.JSON(code, model.Return{
		Error:  errView,
		Data:   data,
		Status: "error",
	})
}
