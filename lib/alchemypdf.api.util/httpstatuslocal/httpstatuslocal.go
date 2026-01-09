package httpstatuslocal

func Ok() (int, string) {
	return 200, "Ok"
}

func BadRequest() (int, string) {
	return 400, "BadRequest"
}

func Unauthorized() (int, string) {
	return 401, "Unauthorized"
}

func Forbidden() (int, string) {
	return 403, "Forbidden"
}

func NotFound() (int, string) {
	return 404, "NotFound"
}

func Locked() (int, string) {
	return 423, "Locked"
}

func InternalServerError() (int, string) {
	return 500, "InternalServerError"
}

func NotImplementedError() (int, string) {
	return 501, "NotImplementedError"
}

func DeprecatedError() (int, string) {
	return 410, "DeprecatedError"
}
