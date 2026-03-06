package configmock

import (
	"strings"

	"alchemypdf.api/lib/alchemypdf.api.infrastructure/config"
)

type MockConfig struct{}

func NewMockConfig() config.IConfig {
	return MockConfig{}
}

func (c MockConfig) ProjectName() string {
	return "example"
}

func (c MockConfig) ProjectShortName() string {
	return "e.g."
}

func (c MockConfig) EnvName() string {
	return ""
}

func (c MockConfig) Host() string {
	return "API_HOST"
}

func (c MockConfig) Port() string {
	return "API_PORT"
}

func (c MockConfig) ReqHeaderApiKey() string {
	return "x-api-key"
}

func (c MockConfig) ReqHeaderRequestID() string {
	return "x-request-id"
}

func (c MockConfig) RequestIDKey() string {
	return "x-request-id"
}

func (c MockConfig) DbConnectionString() string {
	return "db://0.0.0.0:0000"
}

func (c MockConfig) APIKeys() []string {
	return strings.Split("API_KEY", " ")
}

func (c MockConfig) RunSwagger() bool {
	return true
}

func (c MockConfig) LogLevel() string {
	return "Info"
}

func (c MockConfig) LogHTTPStatusCode() int {
	return 200
}
