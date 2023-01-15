package repositories

import (
	"context"
	"encoding/json"
	"testjavan/helpers"
	"testjavan/helpers/constants"
	"testjavan/model"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type ProductRepository interface {
	GetAllProduct(ctx context.Context) (model.Products, error)
}

type product struct{}

func newProductRepository() ProductRepository {
	return &product{}
}

func (p *product) GetAllProduct(ctx context.Context) (model.Products, error) {
	var (
		result    model.Products
		err       error
		requestID = ctx.Value(helpers.RequestIDKey).(string)
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	client := resty.New()

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		by, err := json.Marshal(r.Body)
		if err != nil {
			return err
		}

		//Log request
		logrus.WithFields(logrus.Fields{
			"type":         "request-api",
			"X-Request-ID": requestID,
			"method":       r.Method,
			"url":          r.URL,
			"at":           time.Now().Format(constants.TimeFormat),
		}).Info(string(by))

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		//Log request
		logrus.WithFields(logrus.Fields{
			"type":         "response-api",
			"X-Request-ID": requestID,
			"status":       r.StatusCode(),
			"method":       r.Request.Method,
			"duration":     r.Time().String(),
			"url":          r.Request.URL,
			"at":           time.Now().Format(constants.TimeFormat),
		}).Info()

		return nil
	})

	url := "https://dummyjson.com/products"

	_, err = client.R().SetContext(ctx).
		SetHeader(helpers.RequestID, requestID).
		SetQueryParams(map[string]string{
			"limit":  "100",
			"select": "id,title,price",
		}).SetResult(&result).
		Get(url)

	if err != nil {
		return result, err
	}

	return result, err
}
