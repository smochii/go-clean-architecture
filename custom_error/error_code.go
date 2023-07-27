package custom_error

var ErrorCodes *errorCodes

type errorCode string

type errorCodes struct {
	InternalError        errorCode
	PermissionError      errorCode
	BadRequestError      errorCode
	NotFoundError        errorCode
	UnauthorizedError    errorCode
	MaintenanceModeError errorCode
	DisabledAccountError errorCode
}

func init() {
	ErrorCodes = &errorCodes{}
	ErrorCodes.InternalError = "internal_error"
	ErrorCodes.PermissionError = "permission_error"
	ErrorCodes.BadRequestError = "bad_request_error"
	ErrorCodes.NotFoundError = "not_found_error"
	ErrorCodes.UnauthorizedError = "unauthorized_error"
}
