package main

import (
	"github.com/1206yaya/go-echo-jwt-noteapp-api/controller"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/db"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/repository"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/router"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/usecase"
)

func main() {

	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
