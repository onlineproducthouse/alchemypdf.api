package errorlocal_test

import (
	"errors"
	"testing"

	"alchemypdf.api/lib/alchemypdf.api.testhelpers/utils/asserterrmsg"
	"alchemypdf.api/lib/alchemypdf.api.util/errorlocal"
	"alchemypdf.api/lib/alchemypdf.api.util/httpstatuslocal"
)

func TestValidationErr(t *testing.T) {
	op := "TestValidationErr"
	msg := "Testing Validation Error"
	innerMsg := "Inner Testing Validation Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ValidationErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ValidationErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ValidationErrStr, err.Kind()))
	}

	code, _ := httpstatuslocal.BadRequest()
	if err.StatusCode() != code {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error StatusCode", code, err.StatusCode()))
	}

	if err.Error() != msg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error", msg, err.Error()))
	}

	if err.OriginalErr().Error() != originalErr.Error() {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("OriginalErr", originalErr, err.OriginalErr()))
	}

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Trace InnerMessage", innerMsg, err.Trace().InnerMessage))
	}
}

func TestAuthErr(t *testing.T) {
	op := "TestAuthErr"
	msg := "Testing Auth Error"
	innerMsg := "Inner Testing Auth Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.AuthErr(msg, op, originalErr)

	if err.Kind() != errorlocal.AuthErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.AuthErrStr, err.Kind()))
	}
}

func TestForbiddenErr(t *testing.T) {
	op := "TestForbiddenErr"
	msg := "Testing ForbiddenErr Error"
	innerMsg := "Inner Testing ForbiddenErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ForbiddenErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ForbiddenErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ForbiddenErrStr, err.Kind()))
	}
}

func TestNotFoundErr(t *testing.T) {
	op := "TestNotFoundErr"
	msg := "Testing NotFoundErr Error"
	innerMsg := "Inner Testing NotFoundErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.NotFoundErr(msg, op, originalErr)

	if err.Kind() != errorlocal.NotFoundErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.NotFoundErrStr, err.Kind()))
	}
}

func TestResorceLockedErr(t *testing.T) {
	op := "TestResorceLockedErr"
	msg := "Testing ResorceLockedErr Error"
	innerMsg := "Inner Testing ResorceLockedErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.ResorceLockedErr(msg, op, originalErr)

	if err.Kind() != errorlocal.ResorceLockedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.ResorceLockedErrStr, err.Kind()))
	}
}

func TestUnknownErr(t *testing.T) {
	op := "TestUnknownErr"
	msg := "Testing UnknownErr Error"
	innerMsg := "Inner Testing UnknownErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.UnknownErr(msg, op, originalErr)

	if err.Kind() != errorlocal.UnknownErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.UnknownErrStr, err.Kind()))
	}
}

func TestNotImplementedErr(t *testing.T) {
	op := "TestNotImplementedErr"
	err := errorlocal.NotImplementedErr(op)

	if err.Kind() != errorlocal.NotImplementedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.NotImplementedErrStr, err.Kind()))
	}
}

func TestDeprecated(t *testing.T) {
	op := "TestDeprecated"
	err := errorlocal.Deprecated(op)

	if err.Kind() != errorlocal.DeprecatedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", errorlocal.DeprecatedErrStr, err.Kind()))
	}
}

func TestCatchErr(t *testing.T) {
	op := "TestCatchErr"
	msg := "Testing CatchErr Error"
	innerMsg := "Inner Testing CatchErr Error"
	originalErr := errors.New(innerMsg)

	err := errorlocal.CatchErr(errorlocal.UnknownErr(msg, op, originalErr), op)

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Trace Message", msg, err.Trace().InnerMessage))
	}
}
