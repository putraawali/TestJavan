package errs

import (
	"errors"
	"net/http"
	"testjavan/model"

	"github.com/labstack/echo/v4"
)

var (
	ErrMethodNotAllowed = errors.New("error: method is not allowed")
	ErrContentType      = errors.New("error: invalid Content-Type")
	ErrMissingParameter = errors.New("error: missing parameter")
	ErrRecordNotfound   = errors.New("error: data not found")
)

var ErrHttpStatusMap = map[string]int{
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	ErrContentType.Error():      http.StatusBadRequest,
	ErrMissingParameter.Error(): http.StatusBadRequest,
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
		Error: errView,
		Data:  data,
	})
}
