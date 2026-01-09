package loglocal

import (
	"strings"

	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	config config.IConfig
}

type ILogger interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(err error)
	Fatal(msg string)
	Panic(msg string)

	AppError(err httperror.IAppError)

	HTTPRequest(method, url, requestID string)
	HTTPResponse(status int, method, url, requestID string)
}

func New(c config.IConfig) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return Logger{c}
}

func (logger Logger) Info(msg string) {
	if strings.ToLower(logger.config.EnvName()) != constant.ENV_PROD {
		log.Info().Msg(msg)
	}
}

func (logger Logger) Debug(msg string) {
	if strings.ToLower(logger.config.EnvName()) != constant.ENV_PROD {
		log.Debug().Msg(msg)
	}
}

func (logger Logger) Warn(msg string) {
	if strings.ToLower(logger.config.EnvName()) != constant.ENV_PROD {
		log.Warn().Msg(msg)
	}
}

func (logger Logger) Error(err error) {
	log.Err(err)
}

func (logger Logger) Fatal(msg string) {
	log.Fatal().Msg(msg)
}

func (logger Logger) Panic(msg string) {
	log.Panic().Msg(msg)
}
