package cmd

import (
	"github.com/H0tab/todo-app"
	"github.com/H0tab/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.Initroutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
