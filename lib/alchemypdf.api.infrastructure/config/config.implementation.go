package config

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (c TConfig) EnvName() string {
	return get("ENVIRONMENT_NAME")
}

func (c TConfig) ProjectName() string {
	return get("ALCHEMYPDF_PROJECT_NAME")
}

func (c TConfig) ProjectShortName() string {
	return get("ALCHEMYPDF_PROJECT_SHORT_NAME")
}

func (c TConfig) RunSwagger() bool {
	return get("RUN_SWAGGER") == "true"
}

func (c TConfig) Host() string {
	return get("ALCHEMYPDF_API_HOST")
}

func (c TConfig) Port() string {
	return get("ALCHEMYPDF_API_PORT")
}

func (c TConfig) APIKeys() []string {
	return strings.Split(get("ALCHEMYPDF_API_KEYS"), ",")
}

func (c TConfig) DbConnectionString() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		get("ALCHEMYPDF_DB_PROTOCOL"),
		get("ALCHEMYPDF_DB_USERNAME"),
		url.QueryEscape(get("ALCHEMYPDF_DB_PASSWORD")),
		get("ALCHEMYPDF_DB_HOST"),
		get("ALCHEMYPDF_DB_PORT"),
		get("ALCHEMYPDF_DB_NAME"),
	)
}

func (c TConfig) LogLevel() string {
	l := get("ALCHEMYPDF_LOG_LEVEL")
	if l == "" {
		return "Info"
	}
	return l
}

func (c TConfig) LogHTTPStatusCode() int {
	result, _ := strconv.Atoi(get("ALCHEMYPDF_LOG_HTTP_STATUS_CODE"))
	return result
}

func (c TConfig) ReqHeaderApiKey() string {
	return "x-api-key"
}

func (c TConfig) ReqHeaderRequestID() string {
	return "x-request-id"
}

func (c TConfig) RequestIDKey() string {
	return "x-request-id"
}
