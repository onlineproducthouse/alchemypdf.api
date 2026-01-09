package mdlwr

import (
	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponse"
)

func (api MiddlewareAPI) Deprecated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := httperror.Deprecated("This operation has been deprecated.")
		api.logger.AppError(err)
		return c.JSON(err.StatusCode(), httpresponse.Default(err.Error(), err.StatusCode()))
	}
}
