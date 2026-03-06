package config

type IConfig interface {
	ProjectName() string
	ProjectShortName() string
	EnvName() string
	Host() string
	Port() string

	ReqHeaderApiKey() string
	ReqHeaderRequestID() string
	RequestIDKey() string

	DbConnectionString() string

	APIKeys() []string

	RunSwagger() bool

	// Panic, Fatal, Error, Debug, Warn, Info
	LogLevel() string
	LogHTTPStatusCode() int
}
