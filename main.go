package main

import (
	"01APi/config"
	"01APi/controller"
	"01APi/helper"
	"01APi/repository"
	"01APi/router"
	"01APi/service"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start server")
	//database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	//  controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
