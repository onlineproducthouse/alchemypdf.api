package requestmodelmock

import (
	"alchemypdf.api/lib/alchemypdf.api.model/requestmodel"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
)

type (
	RequestModelMockResolver struct {
		Insert                             func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError)
		SelectByID                         func(requestID int) ([]*requestmodel.Schema, *httperrorutil.AppError)
		SelectByClientReference            func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError)
		SelectWithContentByClientReference func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError)
		SelectByState                      func(stateKey string) ([]*requestmodel.Schema, *httperrorutil.AppError)
		StateUpdate                        func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError)
		AttemptCountIncrement              func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError)
	}

	RequestModelMock struct {
		resolver RequestModelMockResolver
	}
)

func NewRequestModelMock(resolver RequestModelMockResolver) RequestModelMock {
	return RequestModelMock{resolver}
}

func (mock RequestModelMock) Insert(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.Insert(payload)
}

func (mock RequestModelMock) SelectByID(requestID int) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.SelectByID(requestID)
}

func (mock RequestModelMock) SelectByClientReference(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.SelectByClientReference(clientReference)
}

func (mock RequestModelMock) SelectWithContentByClientReference(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.SelectWithContentByClientReference(clientReference)
}

func (mock RequestModelMock) SelectByState(stateKey string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.SelectByState(stateKey)
}

func (mock RequestModelMock) StateUpdate(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.StateUpdate(payload)
}

func (mock RequestModelMock) AttemptCountIncrement(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
	return mock.resolver.AttemptCountIncrement(payload)
}
