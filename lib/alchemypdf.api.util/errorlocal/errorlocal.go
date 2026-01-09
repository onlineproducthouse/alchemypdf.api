package errorlocal

import "alchemypdf.api/lib/alchemypdf.api.util/httpstatuslocal"

const (
	NoErrStr             string = "Ok"                  // status 200
	ValidationErrStr     string = "ValidationError"     // status 400
	AuthErrStr           string = "AuthenticationError" // status 401
	ForbiddenErrStr      string = "ForbiddenError"      // status 403
	NotFoundErrStr       string = "NotFoundError"       // status 404
	ResorceLockedErrStr  string = "ResorceLockedError"  // status 423
	UnknownErrStr        string = "UnknownError"        // status 500
	NotImplementedErrStr string = "NotImplementedError" // status 501
	DeprecatedErrStr     string = "DeprecatedError"     // status 410
)

type AppError struct {
	name           string
	message        string
	op             string
	httpStatusCode int
	err            error
}

func ValidationErr(msg, op string, err error) *AppError {
	code, _ := httpstatuslocal.BadRequest()
	return &AppError{ValidationErrStr, msg, op, code, err}
}

func AuthErr(msg, op string, err error) *AppError {
	code, _ := httpstatuslocal.Unauthorized()
	return &AppError{AuthErrStr, msg, op, code, err}
}

func ForbiddenErr(msg, op string, err error) *AppError {
	code, _ := httpstatuslocal.Forbidden()
	return &AppError{ForbiddenErrStr, msg, op, code, err}
}

func NotFoundErr(msg, op string, err error) *AppError {
	code, text := httpstatuslocal.NotFound()

	if msg != "" {
		text = msg
	}

	return &AppError{NotFoundErrStr, text, op, code, err}
}

func ResorceLockedErr(msg, op string, err error) *AppError {
	code, _ := httpstatuslocal.Locked()
	return &AppError{ResorceLockedErrStr, msg, op, code, err}
}

func UnknownErr(msg, op string, err error) *AppError {
	code, _ := httpstatuslocal.InternalServerError()
	return &AppError{UnknownErrStr, msg, op, code, err}
}

func NotImplementedErr(op string) *AppError {
	code, text := httpstatuslocal.NotImplementedError()
	return &AppError{NotImplementedErrStr, text, op, code, nil}
}

func Deprecated(op string) *AppError {
	code, text := httpstatuslocal.DeprecatedError()
	return &AppError{DeprecatedErrStr, text, op, code, nil}
}

func CatchErr(err IAppError, op string) *AppError {
	return &AppError{err.Kind(), err.Error(), op, err.StatusCode(), err}
}
