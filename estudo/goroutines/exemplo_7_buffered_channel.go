package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const numberGoroutines = 4
const numberTasks = 10

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	wg.Add(numberGoroutines)

	tasks := make(chan string, numberTasks)

	for n := 0; n < numberGoroutines; n++ {
		id := n + 1
		go job(id, tasks)
	}

	for i := 0; i < numberTasks; i++ {
		id := i + 1
		tasks <- fmt.Sprint("Tasks: ", id)
	}

	close(tasks)

	wg.Wait()
	fmt.Println("Finalizado")
}

func job(id int, tasks chan string) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Job %d desligando\n", id)
			return
		}

		fmt.Printf("\nJob %d iniciou a execução da task %s\n", id, task)

		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Job %d Finalizou a execução da task %s\n", id, task)
	}
}
