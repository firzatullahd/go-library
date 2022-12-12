package errors

type Errors interface {
	Error() string
	HttpStatusCode() int
	CustomCode() string
}

type errors struct {
	Err      error
	HttpCode int
	Code     int
}

func NewError(e error, httpCode int, customCode int) Errors {
	return &errors{
		Err:      e,
		HttpCode: httpCode,
		Code:     customCode,
	}
}

func (e *errors) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e *errors) HttpStatusCode() int {
	if e == nil {
		return 0
	}
	return e.HttpCode
}

func (e *errors) CustomCode() string {
	if e == nil || e.Code == 0 {
		return ""
	}

	return CustomMessage[e.Code]
}
