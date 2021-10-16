package main

import (
	"log"

	"github.com/ainurqa95/go-todo-list-app"
)

func main() {

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
