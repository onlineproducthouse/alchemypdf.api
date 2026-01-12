package loglocalmock

import "github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"

type LoggerMock struct{}

func NewLoggerMock() LoggerMock {
	return LoggerMock{}
}

func (mock LoggerMock) Info(msg string)                                        {}
func (mock LoggerMock) Debug(msg string)                                       {}
func (mock LoggerMock) Warn(msg string)                                        {}
func (mock LoggerMock) Error(err error)                                        {}
func (mock LoggerMock) Fatal(msg string)                                       {}
func (mock LoggerMock) Panic(msg string)                                       {}
func (mock LoggerMock) AppError(err httperror.IAppError)                       {}
func (mock LoggerMock) HTTPRequest(method, url, requestID string)              {}
func (mock LoggerMock) HTTPResponse(status int, method, url, requestID string) {}
