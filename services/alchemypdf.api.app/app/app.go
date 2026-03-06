package app

import (
	"context"
	"fmt"

	"alchemypdf.api/docs"
	infrastructure "alchemypdf.api/lib/alchemypdf.api.infrastructure"
	ioc "alchemypdf.api/lib/alchemypdf.api.ioc"
	"alchemypdf.api/services/alchemypdf.api.app/app/mdlwr"
	"alchemypdf.api/services/alchemypdf.api.app/app/router"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Api struct {
	app   *echo.Echo
	infra infrastructure.IInfrastructure
	ioc   ioc.IIoc
	mdlwr mdlwr.MiddlewareAPI
}

type IApi interface {
	Build() *Api
	Start() error
	Stop(ctx context.Context) error
}

// @version 1.0
// @description AlcheMyPDF Rest API
// @contact.name Online Product House
// @contact.email info@onlineproducthouse.com
// @termsOfService https://www.onlineproducthouse.com
// @schemes http https
// @BasePath /api
func NewApi(infra infrastructure.IInfrastructure, ioc ioc.IIoc) *Api {
	docs.SwaggerInfo.Title = infra.Config().ProjectName()
	docs.SwaggerInfo.Host = fmt.Sprint(infra.Config().Host(), ":", infra.Config().Port())

	return &Api{
		echo.New(),
		infra,
		ioc,
		mdlwr.New(
			infra.Config(),
			infra.Logger(),
		),
	}
}

func (s *Api) Build() *Api {
	s.infra.Logger().Info("started: building API")

	s.app.Use(middleware.Recover())

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{""},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: false,
		MaxAge:           0,
		ExposeHeaders: []string{
			s.infra.Config().ReqHeaderApiKey(),
			s.infra.Config().ReqHeaderRequestID(),
		},
	}))

	s.app.Use(s.mdlwr.RequestID)
	s.app.Use(s.mdlwr.HTTPReqLog)

	s.infra.Logger().Info("started: registering routes")

	router.NewRouter(s.infra, s.ioc).BuildRouters(s.app)

	s.infra.Logger().Info("done: registering routes")

	s.infra.Logger().Info("done: building API")

	return s
}

func (s *Api) Start() error {
	s.infra.Logger().Info("starting API")
	return s.app.Start(fmt.Sprintf(":%s", s.infra.Config().Port()))
}

func (s *Api) Stop(ctx context.Context) error {
	s.infra.Logger().Info("stopping API")

	// if err := s.app.Shutdown(ctx); err != nil {
	// 	return err
	// }

	return nil
}
