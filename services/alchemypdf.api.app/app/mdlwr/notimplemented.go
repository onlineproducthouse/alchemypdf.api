package mdlwr

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponse"
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
				err := httperror.NotImplementedErr("MiddlewareAPI.NotImplemented")
				api.logger.AppError(err)
				return c.JSON(err.StatusCode(), httpresponse.Default(err.Error(), err.StatusCode()))
			}

			return next(c)
		}
	}
}
