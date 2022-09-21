package main

import (
	"github.com/gtxiqbal/altera-agmc/Day-6/config"
	"github.com/gtxiqbal/altera-agmc/Day-6/helper"
	"github.com/gtxiqbal/altera-agmc/Day-6/repository"
	"github.com/gtxiqbal/altera-agmc/Day-6/router"
	"github.com/gtxiqbal/altera-agmc/Day-6/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := config.NewDBMySQL()
	mongoDb := config.NewDBMongo()
	userRepositoryImpl := repository.NewUserRepositoryImpl(db)
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl)

	bookRepositoryImpl := repository.NewBookRepositoryImpl(mongoDb)
	bookServiceImpl := service.NewBookServiceImpl(bookRepositoryImpl)

	loginServiceImpl := service.NewLoginServiceImpl(userRepositoryImpl)

	r := router.NewRouter()
	r.SetRouteUser(userServiceImpl)
	r.SetRouteBook(bookServiceImpl)
	r.SetRouteLogin(loginServiceImpl)
	r.StartServer()
}
