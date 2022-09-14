package controller

import (
	"github.com/gtxiqbal/altera-agmc/Day-2/helper"
	"github.com/gtxiqbal/altera-agmc/Day-2/model/dto"
	"github.com/gtxiqbal/altera-agmc/Day-2/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type UserController struct {
	userSvc service.UserService
}

func NewUserController(userSvc service.UserService) *UserController {
	return &UserController{userSvc: userSvc}
}

func (userCtrl *UserController) GetAll(c echo.Context) error {
	return c.JSON(200, userCtrl.userSvc.GetAll(c.Request().Context()))
}

func (userCtrl *UserController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)
	return c.JSON(200, userCtrl.userSvc.GetByID(c.Request().Context(), uint(id)))
}

func (userCtrl *UserController) Create(c echo.Context) error {
	var userDtoReq dto.UserDtoReq
	err := c.Bind(&userDtoReq)
	helper.PanicIfErrorCode(400, err)
	return c.JSON(200, userCtrl.userSvc.Create(c.Request().Context(), userDtoReq))
}

func (userCtrl *UserController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)

	var userDtoReq dto.UserDtoReq
	err = c.Bind(&userDtoReq)
	helper.PanicIfErrorCode(400, err)

	userDtoReq.ID = uint(id)
	return c.JSON(200, userCtrl.userSvc.Update(c.Request().Context(), userDtoReq))
}

func (userCtrl *UserController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)
	return c.JSON(200, userCtrl.userSvc.Delete(c.Request().Context(), uint(id)))
}
