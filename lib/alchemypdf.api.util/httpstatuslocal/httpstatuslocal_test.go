package httpstatuslocal_test

import (
	"testing"

	"alchemypdf.api/lib/alchemypdf.api.testhelpers/utils/asserterrmsg"
	"alchemypdf.api/lib/alchemypdf.api.util/httpstatuslocal"
)

func TestOk(t *testing.T) {
	statusCode, message := httpstatuslocal.Ok()

	if statusCode != 200 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 200, statusCode))
	}

	if message != "Ok" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "Ok", message))
	}
}
func TestBadRequest(t *testing.T) {
	statusCode, message := httpstatuslocal.BadRequest()

	if statusCode != 400 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 400, statusCode))
	}

	if message != "BadRequest" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "BadRequest", message))
	}
}
func TestUnauthorized(t *testing.T) {
	statusCode, message := httpstatuslocal.Unauthorized()

	if statusCode != 401 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 401, statusCode))
	}

	if message != "Unauthorized" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "Unauthorized", message))
	}
}
func TestForbidden(t *testing.T) {
	statusCode, message := httpstatuslocal.Forbidden()

	if statusCode != 403 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 403, statusCode))
	}

	if message != "Forbidden" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "Forbidden", message))
	}
}
func TestNotFound(t *testing.T) {
	statusCode, message := httpstatuslocal.NotFound()

	if statusCode != 404 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 404, statusCode))
	}

	if message != "NotFound" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "NotFound", message))
	}
}
func TestLocked(t *testing.T) {
	statusCode, message := httpstatuslocal.Locked()

	if statusCode != 423 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 423, statusCode))
	}

	if message != "Locked" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "Locked", message))
	}
}
func TestInternalServerError(t *testing.T) {
	statusCode, message := httpstatuslocal.InternalServerError()

	if statusCode != 500 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 500, statusCode))
	}

	if message != "InternalServerError" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "InternalServerError", message))
	}
}
func TestNotImplementedError(t *testing.T) {
	statusCode, message := httpstatuslocal.NotImplementedError()

	if statusCode != 501 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 501, statusCode))
	}

	if message != "NotImplementedError" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "NotImplementedError", message))
	}
}
func TestDeprecatedError(t *testing.T) {
	statusCode, message := httpstatuslocal.DeprecatedError()

	if statusCode != 410 {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Code", 410, statusCode))
	}

	if message != "DeprecatedError" {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("HTTP Status Message", "DeprecatedError", message))
	}
}
