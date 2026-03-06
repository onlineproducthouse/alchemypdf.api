package mdlwr

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func (api MiddlewareAPI) RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	// const op string = "MiddlewareAPI.RequestID"

	return func(c *echo.Context) error {
		c.Set(api.config.RequestIDKey(), uuid.NewString())
		return next(c)
	}
}
