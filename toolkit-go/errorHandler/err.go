package errorHandler

import (
	"fmt"
	"net/http"
)

type ErrorType int

const (
	VALIDATION_ERROR     ErrorType = 1
	BIZ_VALIDATION_ERROR ErrorType = 2
	NOT_FOUND_ERROR      ErrorType = 3
	DUPLICATED_ERROR     ErrorType = 4
	DB_ERROR             ErrorType = 5
	SERVER_ERROR         ErrorType = 6
	OTHER_ERROR          ErrorType = 99
)

type errorInfo struct {
	eorigin  error
	etype    ErrorType
	ecode    string
	emessage string
}

func ErrorInstance(origin error, errorType ErrorType, errorCode string, message string, messageVars ...string) IError {
	return errorInfo{
		origin,
		errorType,
		errorCode,
		fmt.Sprintf(message, messageVars),
	}
}

func (f errorInfo) GetError() error {
	return f.eorigin
}

func (f errorInfo) GetCode() string {
	return f.ecode
}

func (f errorInfo) GetMessage() string {
	return f.emessage
}

func (f errorInfo) GetErrorType() ErrorType {
	return f.etype
}

func (f errorInfo) GetHttpCode() int {
	switch f.etype {
	case NOT_FOUND_ERROR:
		return http.StatusNotFound
	case DUPLICATED_ERROR:
		return http.StatusConflict
	case VALIDATION_ERROR:
		return http.StatusBadRequest
	case BIZ_VALIDATION_ERROR:
		return http.StatusBadRequest
	case SERVER_ERROR:
		return http.StatusInternalServerError
	case DB_ERROR:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
