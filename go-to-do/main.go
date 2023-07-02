package main

import (
	"fmt"
	"go-to-do/logic"
)

func main() {
	fmt.Println("vim-go")
	todo := logic.NewToDo(1, "prueba", "Desc", "Pruebas")
	fmt.Println(todo)
}
