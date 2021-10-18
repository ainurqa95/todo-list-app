package main

import (
	"log"

	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/handler"
)

func main() {

	srv := new(todo.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
