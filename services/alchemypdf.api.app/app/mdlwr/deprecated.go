package mdlwr

import (
	"github.com/labstack/echo/v5"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
)

func (api MiddlewareAPI) Deprecated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		err := httperrorutil.Deprecated("This operation has been deprecated.")
		api.logger.AppError(err)
		return c.JSON(err.StatusCode(), httpresponseutil.Default(err.Error(), err.StatusCode()))
	}
}
