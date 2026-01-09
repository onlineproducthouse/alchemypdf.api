package mdlwr

import (
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/loglocal"
)

type MiddlewareAPI struct {
	config config.IConfig
	logger loglocal.ILogger
}

func New(
	config config.IConfig,
	logger loglocal.ILogger,
) MiddlewareAPI {
	return MiddlewareAPI{
		config,
		logger,
	}
}
