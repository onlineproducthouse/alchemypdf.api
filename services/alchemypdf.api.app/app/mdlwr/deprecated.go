package mdlwr

import (
	"alchemypdf.api/lib/alchemypdf.api.util/errorlocal"
	"alchemypdf.api/lib/alchemypdf.api.util/httpresponse"
	"github.com/labstack/echo/v4"
)

func (api MiddlewareAPI) Deprecated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := errorlocal.Deprecated("This operation has been deprecated.")
		api.logger.AppError(err)
		return c.JSON(err.StatusCode(), httpresponse.Default(err.Error(), err.StatusCode()))
	}
}
