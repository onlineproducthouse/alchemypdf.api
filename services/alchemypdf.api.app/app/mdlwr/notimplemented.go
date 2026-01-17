package mdlwr

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
)

func (api MiddlewareAPI) NotImplemented(envList []string) func(next echo.HandlerFunc) echo.HandlerFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {
			isNotImplemented := false

			for _, env := range envList {
				if strings.EqualFold(api.config.EnvName(), env) {
					isNotImplemented = true
					break
				}
			}

			if isNotImplemented {
				err := httperrorutil.NotImplementedErr("MiddlewareAPI.NotImplemented")
				api.logger.AppError(err)
				return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
			}

			return next(c)
		}
	}
}
