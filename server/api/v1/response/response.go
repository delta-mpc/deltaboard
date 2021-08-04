package response

import (
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
)

type ResponseMessage interface {
	SetMessage(message string)
	GetMessage() string
}

type ErrorResponseMessage interface {
	GetErrorType() string
	SetErrorType(errType string)
}

type ExceptionResponseMessage interface {
	GetException() string
}

type Response struct {
	Message string `json:"message" description:"The response message. Will be 'success' or a detailed error type"`
}

func (r *Response) SetMessage(message string) {
	r.Message = message
}

func (r *Response) GetMessage() string {
	return r.Message
}

type ErrorResponse struct {
	Response
}

func (r *ErrorResponse) GetErrorType() string {
	return r.GetMessage()
}

func (r *ErrorResponse) SetErrorType(errType string) {
	r.SetMessage(errType)
}

func (r *ErrorResponse) Error() string {
	return r.GetErrorType()
}

type ExceptionResponse struct {
	Response
}

func (e *ExceptionResponse) Error() string {
	return e.Message
}

func (e *ExceptionResponse) GetException() string {
	return e.Message
}

func NewExceptionResponse(err error) *ExceptionResponse {
	r := &ExceptionResponse{}
	r.SetMessage(err.Error())
	return r
}

func TonicErrorResponse(ctx *gin.Context, err error) (int, interface{}) {
	if e, ok := err.(tonic.BindError); ok {
		validationErrorResponse := NewValidationErrorResponse()
		validationErr := e.ValidationErrors()
		// We return only the first error
		for _, err := range validationErr {
			validationErrorResponse.SetFieldName(err.Field())
			validationErrorResponse.SetFieldMessage(err.Tag())
			return 400, validationErrorResponse
		}
		validationErrorResponse.SetFieldName(e.GetField())
		validationErrorResponse.SetFieldMessage(e.GetMessage())
		return 400, validationErrorResponse
	}

	if err, ok := err.(ErrorResponseMessage); ok {
		return 400, err
	}

	if err, ok := err.(ExceptionResponseMessage); ok {
		return 500, err
	}

	return 500, NewExceptionResponse(err)
}

func TonicRenderResponse(ctx *gin.Context, statusCode int, payload interface{}) {

	if payload, ok := payload.(ResponseMessage); ok {
		if payload.GetMessage() == "" {
			payload.SetMessage("success")
		}
	}

	tonic.DefaultRenderHook(ctx, statusCode, payload)
}
