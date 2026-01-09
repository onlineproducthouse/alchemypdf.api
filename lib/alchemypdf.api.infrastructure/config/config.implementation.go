package config

import (
	"fmt"
	"net/url"
	"strings"
)

func (c TConfig) ProjectName() string {
	return get("PROJECT_NAME")
}

func (c TConfig) ProjectShortName() string {
	return get("PROJECT_SHORT_NAME")
}

func (c TConfig) EnvName() string {
	return get("ENVIRONMENT_NAME")
}

func (c TConfig) Host() string {
	return get("API_HOST")
}

func (c TConfig) Port() string {
	return get("API_PORT")
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

func (c TConfig) DbConnectionString() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		get("DB_PROTOCOL"),
		get("DB_USERNAME"),
		url.QueryEscape(get("DB_PASSWORD")),
		get("DB_HOST"),
		get("DB_PORT"),
		get("DB_NAME"),
	)
}

func (c TConfig) APIKeys() []string {
	return strings.Split(get("API_KEYS"), ",")
}

func (c TConfig) RunSwagger() bool {
	return get("RUN_SWAGGER") == "true"
}
