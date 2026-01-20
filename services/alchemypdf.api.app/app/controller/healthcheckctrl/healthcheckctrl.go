package healthcheckctrl

import (
	"fmt"

	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/loglocal"
	"alchemypdf.api/lib/alchemypdf.api.service/healthchecksvc"
	"github.com/labstack/echo/v4"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatusutil"
)

type HealthCheck struct {
	logger loglocal.ILogger
	hc     healthchecksvc.IHealthCheckService
	config config.IConfig
}

type IHealthCheck interface {
	Ping(c echo.Context) error
}

func New(logger loglocal.ILogger, hc healthchecksvc.IHealthCheckService, config config.IConfig) HealthCheck {
	return HealthCheck{logger, hc, config}
}

// HealthCheck/Ping godoc
// @id HealthCheck.Ping
// @tags HealthCheck
// @summary Gets the health status of server.
// @router /HealthCheck/Ping [get]
// @accept */*
// @produce json
// @success 200 {object} httpresponseutil.Response
// @Failure 500 {object} httpresponseutil.Response
func (ctrl HealthCheck) Ping(c echo.Context) error {

	if err := ctrl.hc.Ping(); err != nil {
		ctrl.logger.Fatal(err.Error())
	}

	statusCode, statusCodeText := httpstatusutil.Ok()
	ctrl.logger.HTTPResponse(statusCode, c.Request().Method, c.Request().RequestURI, fmt.Sprint(c.Get(ctrl.config.RequestIDKey())))

	return c.JSON(statusCode, httpresponseutil.Default(statusCodeText, statusCode))
}
