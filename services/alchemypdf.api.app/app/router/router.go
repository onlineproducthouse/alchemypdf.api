package router

import (
	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	infrastructure "alchemypdf.api/lib/alchemypdf.api.infrastructure"
	ioc "alchemypdf.api/lib/alchemypdf.api.ioc"
	"alchemypdf.api/services/alchemypdf.api.app/app/controller/healthcheckctrl"
	"alchemypdf.api/services/alchemypdf.api.app/app/controller/requestctrl"
	"alchemypdf.api/services/alchemypdf.api.app/app/mdlwr"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	infra      infrastructure.IInfrastructure
	ioc        ioc.IIoc
	middleware mdlwr.MiddlewareAPI
}

type IRouter interface {
	BuildRouters(app *echo.Echo)
}

func NewRouter(infra infrastructure.IInfrastructure, container ioc.IIoc) Router {
	return Router{
		infra,
		container,
		mdlwr.New(
			infra.Config(),
			infra.Logger(),
		),
	}
}

func (r Router) BuildRouters(app *echo.Echo) {

	echo.New()

	r.buildSwaggerRouters(app)

	apiGroup := app.Group("/api")
	apiGroupV1 := apiGroup.Group("/v1", r.middleware.APIKey)

	r.buildHealthCheckRouters(apiGroup)
	r.buildRequestRouters(apiGroupV1)
}

func (r Router) buildSwaggerRouters(app *echo.Echo) {
	if r.infra.Config().RunSwagger() {
		r.infra.Logger().Info("swagger enabled")
		app.GET("/swagger/*", echoSwagger.WrapHandler, r.middleware.NotImplemented([]string{constant.ENV_PROD}))
	}
}

func (r Router) buildHealthCheckRouters(apiGroup *echo.Group) {
	ctrl := healthcheckctrl.New(r.infra.Logger(), r.ioc.HealthCheck(), r.infra.Config())
	apiGroup.GET("/HealthCheck/Ping", ctrl.Ping)
}

func (r Router) buildRequestRouters(apiGroupV1 *echo.Group) {
	ctrl := requestctrl.NewRequestCtrl(
		r.infra.Config(),
		r.infra.Logger(),
		r.ioc.RequestService(),
	)

	group := apiGroupV1.Group("/Request")

	group.POST("/Create", ctrl.HandleCreate)
	group.GET("/GetByClientReference/:ClientReference", ctrl.HandleGetByClientReference)
	group.GET("/GetWithContentByClientReference/:ClientReference", ctrl.HandleGetWithContentByClientReference)
	group.GET("/GetPending", ctrl.HandleGetPending)
	group.POST("/Complete", ctrl.HandleComplete)
	group.POST("/Callback", ctrl.HandleCallback, r.middleware.NotImplemented([]string{constant.ENV_PROD}))
}
