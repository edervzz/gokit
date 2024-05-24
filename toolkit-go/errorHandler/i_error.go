package errorHandler

type IError interface {
	GetError() error
	GetCode() string
	GetMessage() string
	GetErrorType() ErrorType
	GetHttpCode() int
}
