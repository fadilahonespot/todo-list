package cuserr

import (
	"net/http"

	"github.com/fadilahonespot/todo-list/utils/model"
	"github.com/gofiber/fiber/v2"
)

type ApplicationError struct {
	ErrorCode      int
	Message        string
	Data           interface{}
	OverideMessage bool
}

func SetError(errorCode int, message string) error {
	return &ApplicationError{
		ErrorCode: errorCode,
		Message:   message,
	}
}

func SetErrorMessage(errorCode int, message string) error {
	return &ApplicationError{
		ErrorCode:      errorCode,
		Message:        message,
		OverideMessage: true,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func (e *ApplicationError) Code() int {
	return e.ErrorCode
}

func GetErrorCode(err error) int {
	if err == nil {
		return 0
	}

	if se, ok := err.(interface {
		Code() int
	}); ok {
		return se.Code()
	}
	return 0
}

func ErrorHandle(c *fiber.Ctx, err error) error {
	code := GetErrorCode(err)
	if code == 0 {
		code = http.StatusInternalServerError
	}

	message := http.StatusText(code)
	if he, ok := err.(*ApplicationError); ok {
		if he.OverideMessage {
			message = he.Message
		}
	}

	resp := model.Response{
		Status:  http.StatusText(code),
		Message: message,
		Data:    struct{}{},
	}

	return c.Status(code).JSON(resp)
}
