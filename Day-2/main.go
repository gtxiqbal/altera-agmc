package main

import (
	"github.com/gtxiqbal/altera-agmc/Day-2/config"
	"github.com/gtxiqbal/altera-agmc/Day-2/helper"
	"github.com/gtxiqbal/altera-agmc/Day-2/repository"
	"github.com/gtxiqbal/altera-agmc/Day-2/router"
	"github.com/gtxiqbal/altera-agmc/Day-2/service"
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

	r := router.NewRouter()
	r.SetRouteUser(userServiceImpl)
	r.SetRouteBook(bookServiceImpl)
	r.StartServer()
}
