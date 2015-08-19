package main

import (
	"fmt"
	"go-in-action/estudo/concurrency/runner"
	"time"
)

const timeout = 2 * time.Second

func main() {
	r := runner.New(timeout)
	r.Adiciona(createTask(), createTask(),
		createTask(), createTask())

	// r.ExecutaTasks()

	// if err := r.ExecutaTasks(); err != nil {
	// 	fmt.Println(err)
	// }

	// r.ExecutaTasksGoRoutine()

	r.ExecutaTasksGoRoutine2()

	// r.ImprimeTasks()

	// f := createTask()
	// f('1')
}

func createTask() func(int) {

	return func(id int) {
		fmt.Println("Processando Task:", id)
		time.Sleep(time.Duration(id) * time.Second)
		fmt.Println("Finalizando Task:", id)
	}
}
