package errorlocal

type IAppError interface {
	Kind() string
	StatusCode() int
	Error() string
	Trace() Trace
	OriginalErr() error
}

func (a *AppError) Kind() string {
	return a.name
}

func (a *AppError) StatusCode() int {
	return a.httpStatusCode
}

func (a *AppError) Error() string {
	return a.message
}

func (a *AppError) OriginalErr() error {
	return a.err
}

type Trace struct {
	InnerMessage string
	Ops          []string
}

func (a *AppError) Trace() Trace {
	return Trace{retrieveInnerMessage(a), retrieveOperations(a)}
}

func retrieveInnerMessage(a *AppError) string {
	if innerAppError, ok := a.err.(*AppError); ok {
		return retrieveInnerMessage(innerAppError)
	}

	if a.err == nil {
		return ""
	}

	return a.err.Error()
}

func retrieveOperations(a *AppError) []string {
	ops := []string{a.op}

	innerAppError, ok := a.err.(*AppError)
	if !ok {
		return ops
	}

	ops = append(ops, retrieveOperations(innerAppError)...)

	return ops
}
