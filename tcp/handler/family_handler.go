package handler

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testjavan/helpers"
	"testjavan/helpers/errs"
	"testjavan/model"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

func FamilyHandler(method string, args any) (model.Return, error) {
	var (
		result model.Return
		err    error
	)
	switch strings.ToLower(method) {
	case "get":
		argMap := args.(map[string]interface{})
		id := argMap["member_id"]
		idStr := ""
		switch id.(type) {
		case float64:
			idStr = strconv.Itoa(int(id.(float64)))
		case int:
			idStr = strconv.Itoa(id.(int))
		case string:
			idStr = id.(string)
		default:
			return result, errs.ErrInvalidParameter
		}

		result, err = GetFamily(idStr)

	default:
		return result, errors.New("Only accept get method")
	}

	return result, err
}

func GetFamily(args string) (model.Return, error) {
	var (
		result    model.Return
		err       error
		url       = fmt.Sprintf(os.Getenv("BASE_URL")+"/family/member/%s", args)
		requestID = uuid.New().String()
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client := resty.New()

	_, err = client.R().SetContext(ctx).
		SetHeaders(map[string]string{
			helpers.RequestID: requestID,
			"Content-Type":    "application/json",
		}).
		SetResult(&result).
		SetError(&result).
		Get(url)

	if err != nil {
		return result, err
	}

	return result, err
}
