package requestsvc_test

import (
	"testing"

	"alchemypdf.api/lib/alchemypdf.api.contract/requestcontract"
	"alchemypdf.api/lib/alchemypdf.api.model/requestmodel"
	"alchemypdf.api/lib/alchemypdf.api.service/requestsvc"
	"alchemypdf.api/lib/alchemypdf.api.testhelpers/models/requestmodelmock"
	"alchemypdf.api/lib/alchemypdf.api.testhelpers/infrastructure/loglocalmock"


	"github.com/onlineproducthouse/alchemypdf.api.httputils/helpers/asserterrmsg"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
)

func  TestCreate(t *testing.T){
	//setup
	requestID := 1
	requestExternalID := "uuid.NewString()"

	service := requestsvc.NewRequestService(
		loglocalmock.NewLoggerMock(),
		requestmodelmock.NewRequestModelMock(requestmodelmock.RequestModelMockResolver{
	Insert:	func(payload requestmodel.Schema) ([]*requestmodel.Schema,*httperror.AppError){
			return []*requestmodel.Schema{{
				RequestID: requestID,
				RequestExternalID: requestExternalID,
			}}, nil
			},
		}),
	)

	//execute
	err := service.Create(requestcontract.CreateRequest{
		ClientReference :"ClientReference",
		Content         :"Content",
		CallbackURL     :"CallbackURL",
	})

	//assert
	if err != nil {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("err", "nil", err))
	}
}
