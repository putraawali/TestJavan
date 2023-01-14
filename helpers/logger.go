package helpers

import (
	"bytes"
	"io/ioutil"
	"testjavan/helpers/constants"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(e *echo.Echo, skipURL []string) {
	skip := logSkipperURL(skipURL)

	e.Use(logRequest(skip))
	e.Use(logResponse(skip))
}

func logSkipperURL(skipURLs []string) middleware.Skipper {
	return func(c echo.Context) bool {
		path := c.Request().URL.Path
		for _, url := range skipURLs {
			if path == url {
				return true
			}
		}
		return false
	}
}

func logRequest(skipper middleware.Skipper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper != nil && skipper(c) {
				return next(c)
			}

			req := c.Request()
			header := req.Header
			reqBody, _ := ioutil.ReadAll(req.Body)
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

			requestID := ""
			if requestID = header.Get("X-Request-ID"); requestID == "" {
				id := uuid.New()
				requestID = id.String()
				c.Request().Header.Set("X-Request-ID", requestID)
			}

			type log struct {
				logger *logrus.Logger
			}

			logrus.WithFields(logrus.Fields{
				"type":         "request",
				"method":       req.Method,
				"uri":          req.URL.String(),
				"at":           time.Now().Format(constants.TimeFormat),
				"user_agent":   req.UserAgent(),
				"ip":           req.RemoteAddr,
				"X-Request-ID": requestID,
			}).Info(string(reqBody))

			return next(c)
		}
	}
}

func logResponse(skipper middleware.Skipper) echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: skipper,
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			status := c.Response().Status

			req := c.Request()
			logrus.WithFields(logrus.Fields{
				"type":         "request",
				"method":       req.Method,
				"uri":          req.URL.String(),
				"at":           time.Now().Format(constants.TimeFormat),
				"status":       status,
				"X-Request-ID": req.Header.Get("X-Request-ID"),
				"ip":           req.RemoteAddr,
				"user_agent":   req.UserAgent(),
			}).Info(string(resBody))
		},
	})
}
