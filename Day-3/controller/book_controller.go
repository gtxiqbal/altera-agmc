package controller

import (
	"github.com/gtxiqbal/altera-agmc/Day-3/helper"
	"github.com/gtxiqbal/altera-agmc/Day-3/model/dto"
	"github.com/gtxiqbal/altera-agmc/Day-3/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type BookController struct {
	bookSvc service.BookService
}

func NewBookController(bookSvc service.BookService) *BookController {
	return &BookController{bookSvc: bookSvc}
}

func (bookCtrl *BookController) GetAll(c echo.Context) error {
	return c.JSON(200, bookCtrl.bookSvc.GetAll())
}

func (bookCtrl *BookController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(200, bookCtrl.bookSvc.GetByID(uint(id)))
}

func (bookCtrl *BookController) Create(c echo.Context) error {
	var bookDtoReq dto.BookDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&bookDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(bookDtoReq))
	return c.JSON(200, bookCtrl.bookSvc.Create(bookDtoReq))
}

func (bookCtrl *BookController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var bookDtoReq dto.BookDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&bookDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(bookDtoReq))

	bookDtoReq.ID = uint(id)
	return c.JSON(200, bookCtrl.bookSvc.Update(bookDtoReq))
}

func (bookCtrl *BookController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(200, bookCtrl.bookSvc.Delete(uint(id)))
}
