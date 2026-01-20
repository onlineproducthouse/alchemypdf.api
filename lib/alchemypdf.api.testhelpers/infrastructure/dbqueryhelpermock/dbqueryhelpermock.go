package dbqueryhelpermock

import "github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"

type (
	DBQueryHelperMockResolver struct {
		Execute func(dst any, query string, args ...any) (any, *httperrorutil.AppError)
	}

	DBQueryHelperMock struct {
		resolver DBQueryHelperMockResolver
	}
)

func NewDBQueryHelperMock(resolver DBQueryHelperMockResolver) DBQueryHelperMock {
	return DBQueryHelperMock{resolver}
}

func (mock DBQueryHelperMock) Execute(dst any, query string, args ...any) (any, *httperrorutil.AppError) {
	return mock.resolver.Execute(dst, query, args...)
}
