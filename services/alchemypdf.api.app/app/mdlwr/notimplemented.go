package mdlwr

import (
	"strings"

	"github.com/labstack/echo/v4"
	alchemypdfapihttputils "github.com/onlineproducthouse/alchemypdf.api.httputils"
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
				err := alchemypdfapihttputils.NotImplementedErr("MiddlewareAPI.NotImplemented")
				api.logger.AppError(err)
				return c.JSON(err.StatusCode(), alchemypdfapihttputils.Default(err.Error(), err.StatusCode()))
			}

			return next(c)
		}
	}
}
