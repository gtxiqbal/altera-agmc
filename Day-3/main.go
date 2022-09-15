package main

import (
	"github.com/gtxiqbal/altera-agmc/Day-3/config"
	"github.com/gtxiqbal/altera-agmc/Day-3/helper"
	"github.com/gtxiqbal/altera-agmc/Day-3/repository"
	"github.com/gtxiqbal/altera-agmc/Day-3/router"
	"github.com/gtxiqbal/altera-agmc/Day-3/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := config.NewDBMySQL()
	userRepositoryImpl := repository.NewUserRepositoryImpl(db)
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl)

	bookRepositoryImpl := repository.NewBookRepositoryImpl()
	bookServiceImpl := service.NewBookServiceImpl(bookRepositoryImpl)

	loginServiceImpl := service.NewLoginServiceImpl(userRepositoryImpl)

	r := router.NewRouter()
	r.SetRouteUser(userServiceImpl)
	r.SetRouteBook(bookServiceImpl)
	r.SetRouteLogin(loginServiceImpl)
	r.StartServer()
}
