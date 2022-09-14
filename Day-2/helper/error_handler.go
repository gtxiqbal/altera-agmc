package helper

import (
	"fmt"
	"github.com/gtxiqbal/altera-agmc/Day-2/model/dto"
	"github.com/labstack/echo/v4"
)

func PanicIfError(err error) {
	PanicIfErrorCode(500, err)
}

func PanicIfErrorCode(code int, err error) {
	if err != nil {
		PanicErrorCode(code, err)
	}
}

func PanicErrorCode(code int, err error) {
	err = echo.NewHTTPError(code, err)
	panic(err)
}

func CustomHttpErrorHandler(c echo.Context, err error, stack []byte) error {
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		httpError = &echo.HTTPError{
			Code:     500,
			Message:  err.Error(),
			Internal: err,
		}
	}

	c.Logger().Error(httpError.Message)
	fmt.Println(string(stack))

	return c.JSON(httpError.Code, dto.ResponseDTO[any]{
		Code:    httpError.Code,
		Status:  dto.StatusFailed,
		Message: httpError.Message.(error).Error(),
	})
}
