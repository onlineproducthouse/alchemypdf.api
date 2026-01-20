package mdlwr

import (
	"errors"

	"slices"

	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
)

func (api MiddlewareAPI) APIKey(next echo.HandlerFunc) echo.HandlerFunc {
	const op string = "MiddlewareAPI.APIKey"

	return func(c echo.Context) error {
		valid := false

		apiKey := c.Request().Header.Get(api.config.ReqHeaderApiKey())
		if apiKey == "" {
			msg := "api key not found"
			err := httperrorutil.NotFoundErr(msg, op, errors.New(msg))

			api.logger.AppError(err)

			return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
		}

		if slices.Contains(api.config.APIKeys(), apiKey) {
			valid = true
		}

		if !valid {
			return errors.New("invalid api key provided")
		}

		return next(c)
	}
}
