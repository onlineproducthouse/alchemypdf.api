package mdlwr

import (
	"github.com/labstack/echo/v4"
	alchemypdfapihttputils "github.com/onlineproducthouse/alchemypdf.api.httputils"
)

func (api MiddlewareAPI) Deprecated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := alchemypdfapihttputils.Deprecated("This operation has been deprecated.")
		api.logger.AppError(err)
		return c.JSON(err.StatusCode(), alchemypdfapihttputils.Default(err.Error(), err.StatusCode()))
	}
}
