package httpresponse_test

import (
	"fmt"
	"testing"

	"alchemypdf.api/lib/alchemypdf.api.testhelpers/utils/asserterrmsg"
	"alchemypdf.api/lib/alchemypdf.api.util/httpresponse"
)

func TestResponseDefault(t *testing.T) {
	_testCases := []httpresponse.Response{
		{"Ok", 200},
		{"Bad Request", 400},
		{"Unauthorized", 401},
		{"Forbidden", 403},
		{"Not Found", 404},
		{"Locked", 423},
		{"Internal Server Error", 500},
	}

	for _, expect := range _testCases {
		t.Run(fmt.Sprintf("HTTP Response %s", expect.Message), func(t *testing.T) {
			got := httpresponse.Default(expect.Message, expect.StatusCode)

			if got.Message != expect.Message {
				t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Message", expect.Message, got.Message))
			}

			if got.StatusCode != expect.StatusCode {
				t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("StatusCode", expect.StatusCode, got.StatusCode))
			}
		})
	}
}

func BenchmarkResponseDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		httpresponse.Default("Ok", 200)
	}
}
