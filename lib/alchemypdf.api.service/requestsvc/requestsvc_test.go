package requestsvc_test

import (
	"testing"

	constant "alchemypdf.api/lib/alchemypdf.api.constant"
	"alchemypdf.api/lib/alchemypdf.api.contract/requestcontract"
	"alchemypdf.api/lib/alchemypdf.api.model/requestmodel"
	"alchemypdf.api/lib/alchemypdf.api.service/requestsvc"
	"alchemypdf.api/lib/alchemypdf.api.testhelpers/infrastructure/loglocalmock"
	"alchemypdf.api/lib/alchemypdf.api.testhelpers/models/requestmodelmock"

	"github.com/google/uuid"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/helpers/asserterrmsg"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
)

func TestCreate(t *testing.T) {
	//setup
	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			Insert: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{{
					RequestID:         1,
					RequestExternalID: uuid.NewString(),
				}}, nil
			},
		}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: uuid.NewString(),
		Content:         "Content",
		CallbackURL:     "http://api.example.org/path/to/callback",
	})

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}
}

func TestCreate_ErrHandling_UnknownError_FailedToCreate(t *testing.T) {
	//setup
	errMsg := "failed to create request"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			Insert: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{}, nil
			},
		}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: uuid.NewString(),
		Content:         "Content",
		CallbackURL:     "http://api.example.org/path/to/callback",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestCreate_ErrHandling_CatchError_FailedToInsert(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			Insert: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestCreate_ErrHandling_CatchError_FailedToInsert", nil)
			},
		}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: uuid.NewString(),
		Content:         "Content",
		CallbackURL:     "http://api.example.org/path/to/callback",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestCreate_ErrHandling_ValidationError_Required_CallbackURL(t *testing.T) {
	//setup
	errMsg := "request callback url is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: uuid.NewString(),
		Content:         "Content",
		CallbackURL:     "",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestCreate_ErrHandling_ValidationError_Required_Content(t *testing.T) {
	//setup
	errMsg := "request content is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: uuid.NewString(),
		Content:         "",
		CallbackURL:     "",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestCreate_ErrHandling_ValidationError_Required_ClientReference(t *testing.T) {
	//setup
	errMsg := "request client reference is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference: "",
		Content:         "",
		CallbackURL:     "",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestGetByID(t *testing.T) {
	//setup
	requestID := 1

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByID: func(requestID int) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{{RequestID: requestID}}, nil
			},
		}),
	)

	//execute
	res, err := service.GetByID(requestID)

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}

	if res.RequestID != requestID {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res.RequestID", requestID, res.RequestID))
	}
}

func TestGetByID_ErrHandling_NotFoundError(t *testing.T) {
	//setup
	errMsg := "request not found"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByID: func(requestID int) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{}, nil
			},
		}),
	)

	//execute
	res, err := service.GetByID(1)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByID_ErrHandling_CatchError_SelectByID(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByID: func(requestID int) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestGetByID_ErrHandling_CatchError_SelectByID", nil)
			},
		}),
	)

	//execute
	res, err := service.GetByID(1)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByID_ErrHandling_ValidationError_Required_RequestID(t *testing.T) {
	//setup
	errMsg := "request id is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	res, err := service.GetByID(0)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByClientReference(t *testing.T) {
	//setup
	clientReference := uuid.NewString()

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByClientReference: func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{{ClientReference: clientReference}}, nil
			},
		}),
	)

	//execute
	res, err := service.GetByClientReference(clientReference)

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}

	if len(res) != 1 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("len(res)", 1, len(res)))
	}

	if res[0].ClientReference != clientReference {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res[0].ClientReference", clientReference, res[0].ClientReference))
	}
}

func TestGetByClientReference_ErrHandling_CatchError_SelectByID(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByClientReference: func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestGetByClientReference_ErrHandling_CatchError_SelectByID", nil)
			},
		}),
	)

	//execute
	res, err := service.GetByClientReference(uuid.NewString())

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByClientReference_ErrHandling_ValidationError_Required_ClientReference(t *testing.T) {
	//setup
	errMsg := "request client reference is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	res, err := service.GetByClientReference("")

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetWithContentByClientReference(t *testing.T) {
	//setup
	clientReference := uuid.NewString()

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectWithContentByClientReference: func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{{ClientReference: clientReference}}, nil
			},
		}),
	)

	//execute
	res, err := service.GetWithContentByClientReference(clientReference)

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}

	if len(res) != 1 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("len(res)", 1, len(res)))
	}

	if res[0].ClientReference != clientReference {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res[0].ClientReference", clientReference, res[0].ClientReference))
	}
}

func TestGetWithContentByClientReference_ErrHandling_CatchError_SelectByID(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectWithContentByClientReference: func(clientReference string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestGetWithContentByClientReference_ErrHandling_CatchError_SelectByID", nil)
			},
		}),
	)

	//execute
	res, err := service.GetWithContentByClientReference(uuid.NewString())

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetWithContentByClientReference_ErrHandling_ValidationError_Required_ClientReference(t *testing.T) {
	//setup
	errMsg := "request client reference is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	res, err := service.GetWithContentByClientReference("")

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByState(t *testing.T) {
	//setup
	requestID := 1

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByState: func(stateKey string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{{RequestID: requestID}}, nil
			},
		}),
	)

	//execute
	res, err := service.GetByState(constant.RequestStatePending)

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}

	if res.RequestID != requestID {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res.RequestID", requestID, res.RequestID))
	}
}

func TestGetByState_ErrHandling_NotFoundError(t *testing.T) {
	//setup
	errMsg := "request not found"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByState: func(stateKey string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{}, nil
			},
		}),
	)

	//execute
	res, err := service.GetByState(constant.RequestStatePending)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByState_ErrHandling_CatchError_SelectByState(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			SelectByState: func(stateKey string) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestGetByState_ErrHandling_CatchError_SelectByState", nil)
			},
		}),
	)

	//execute
	res, err := service.GetByState(constant.RequestStatePending)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestGetByState_ErrHandling_ValidationError_Required_StateKey(t *testing.T) {
	//setup
	errMsg := "request state key is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	res, err := service.GetByState("")

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}

	if res != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("res", "nil", res))
	}
}

func TestStateUpdate(t *testing.T) {
	//setup
	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			StateUpdate: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{&payload}, nil
			},
		}),
	)

	//execute
	err := service.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       1,
		RequestStateKey: constant.RequestStateAttemptLimitReached,
	})

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}
}

func TestStateUpdate_ErrHandling_UnknownError_FailedToStateUpdate(t *testing.T) {
	//setup
	errMsg := "failed to update request state"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			StateUpdate: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{}, nil
			},
		}),
	)

	//execute
	err := service.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       1,
		RequestStateKey: constant.RequestStateAttemptLimitReached,
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestStateUpdate_ErrHandling_CatchError_FailedToStateUpdate(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			StateUpdate: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestStateUpdate_ErrHandling_CatchError_FailedToStateUpdate", nil)
			},
		}),
	)

	//execute
	err := service.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       1,
		RequestStateKey: constant.RequestStateAttemptLimitReached,
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestStateUpdate_ErrHandling_ValidationError_Required_RequestStateKey(t *testing.T) {
	//setup
	errMsg := "request state key is required"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       1,
		RequestStateKey: "",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestStateUpdate_ErrHandling_ValidationError_Required_RequestID(t *testing.T) {
	//setup
	errMsg := "request id must be greater than zero"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.StateUpdate(requestcontract.StateUpdateRequest{
		RequestID:       0,
		RequestStateKey: "",
	})

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestAttemptCountIncrement(t *testing.T) {
	//setup
	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			AttemptCountIncrement: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{&payload}, nil
			},
		}),
	)

	//execute
	err := service.AttemptCountIncrement(1)

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}
}

func TestAttemptCountIncrement_ErrHandling_UnknownError_FailedToAttemptCountIncrement(t *testing.T) {
	//setup
	errMsg := "failed to update request attempt count"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			AttemptCountIncrement: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return []*requestmodel.Schema{}, nil
			},
		}),
	)

	//execute
	err := service.AttemptCountIncrement(1)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestAttemptCountIncrement_ErrHandling_CatchError_FailedToAttemptCountIncrement(t *testing.T) {
	//setup
	errMsg := "an error occurred"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
			AttemptCountIncrement: func(payload requestmodel.Schema) ([]*requestmodel.Schema, *httperrorutil.AppError) {
				return nil, httperrorutil.UnknownErr(errMsg, "TestAttemptCountIncrement_ErrHandling_CatchError_FailedToAttemptCountIncrement", nil)
			},
		}),
	)

	//execute
	err := service.AttemptCountIncrement(1)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}

func TestAttemptCountIncrement_ErrHandling_ValidationError_Required_RequestID(t *testing.T) {
	//setup
	errMsg := "request id must be greater than zero"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{}),
	)

	//execute
	err := service.AttemptCountIncrement(0)

	//assert
	if err.Error() != errMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err.Error()", errMsg, err))
	}
}
