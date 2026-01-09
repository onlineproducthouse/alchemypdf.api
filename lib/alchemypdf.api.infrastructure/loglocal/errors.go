package loglocal

import (
	"strings"

	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
	"github.com/rs/zerolog/log"
)

func (logger Logger) AppError(err httperror.IAppError) {
	if strings.ToLower(logger.config.EnvName()) == constant.ENV_LOCAL ||
		strings.ToLower(logger.config.EnvName()) == constant.ENV_TEST {
		log.Err(err)
	}

	log.Error().
		Str("kind", err.Kind()).
		Int("httpStatusCode", err.StatusCode()).
		// Str("innerMessage", err.Trace().InnerMessage).
		Str("trace", strings.Join(err.Trace().Ops, ", ")).
		Msg(err.Error())
}
