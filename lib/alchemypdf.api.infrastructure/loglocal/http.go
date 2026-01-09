package loglocal

import "github.com/rs/zerolog/log"

func (logger Logger) HTTPRequest(method, url, requestID string) {
	log.Info().
		Str("method", method).
		Str("url", url).
		Str("requestId", requestID).
		Msg("starting API request")
}

func (logger Logger) HTTPResponse(status int, method, url, requestID string) {
	log.Info().
		Int("status", status).
		Str("method", method).
		Str("url", url).
		Str("requestId", requestID).
		Msg("completed API request")
}
