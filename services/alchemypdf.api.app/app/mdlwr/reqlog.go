package mdlwr

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (api MiddlewareAPI) HTTPReqLog(next echo.HandlerFunc) echo.HandlerFunc {
	// const op string = "MiddlewareAPI.HTTPReqLog"

	return func(c echo.Context) error {
		api.logger.HTTPRequest(c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(api.config.RequestIDKey())))
		return next(c)
	}

}
