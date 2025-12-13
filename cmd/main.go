package main

import (
	"fmt"

	"github.com/luiztapioca/todo-cli/todo"
)

func main() {
	fmt.Println(todo.SaveTask(todo.Task{ID: 3, Title: "teste", Completed: false}))
}
