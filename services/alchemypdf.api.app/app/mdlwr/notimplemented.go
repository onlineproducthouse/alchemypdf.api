package mdlwr

import (
	"strings"

	"alchemypdf.api/lib/alchemypdf.api.util/errorlocal"
	"alchemypdf.api/lib/alchemypdf.api.util/httpresponse"
	"github.com/labstack/echo/v4"
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
				err := errorlocal.NotImplementedErr("MiddlewareAPI.NotImplemented")
				api.logger.AppError(err)
				return c.JSON(err.StatusCode(), httpresponse.Default(err.Error(), err.StatusCode()))
			}

			return next(c)
		}
	}
}
