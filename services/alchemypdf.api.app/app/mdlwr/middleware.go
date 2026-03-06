package mdlwr

import (
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"github.com/onlineproducthouse/alchemypdf.api.logger/loggylog"
)

type MiddlewareAPI struct {
	config config.IConfig
	logger loggylog.ILoggyLog
}

func New(
	config config.IConfig,
	logger loggylog.ILoggyLog,
) MiddlewareAPI {
	return MiddlewareAPI{
		config,
		logger,
	}
}
