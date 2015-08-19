package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Runner struct {
	tasks     []func(int)
	interrupt chan os.Signal
	timeout   <-chan time.Time
	complete  chan error
}

func New(timeout time.Duration) Runner {
	// return Runner{}
	return Runner{
		interrupt: make(chan os.Signal, 1),
		timeout:   time.After(timeout),
		complete:  make(chan error),
	}
}

func (r *Runner) Adiciona(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Executa task executa as tasks
func (r Runner) ExecutaTasks() error {
	fmt.Println("Executando ExecutaTasks")
	signal.Notify(r.interrupt, os.Interrupt)

	for id, task := range r.tasks {
		if r.foiInterrompido() {
			// fmt.Println("Saindo da execução pois foi interrompido")
			// return
			return errors.New("Task finalizado por motivos de interrupção")
		}

		fmt.Println("\nExecutando Task: ", id)
		task(id)
	}

	// como o loop está acima, não tem como entrar aqui
	// para verificar o timeout. Por isso criei outro método de
	// teste para fazer isso via goroutine
	// select {
	// case <-r.timeout:
	// 	return errors.New("Task finalizado por motivos de timeout")
	// }

	return nil
}

// executaTasks2 executa as task em goroutine
// usando wait group ao apertar Ctrl+C sai da routina
// sem usar o signal
func (r Runner) ExecutaTasksGoRoutine() {
	wg.Add(1)

	fmt.Println("Executando ExecutaTasksGoRoutine")
	// signal.Notify(r.interrupt, os.Interrupt)

	go r.executaTasks2()
	wg.Wait()
}

func (r Runner) ExecutaTasksGoRoutine2() {
	fmt.Println("Executando ExecutaTasksGoRoutine2")
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.executaTasks()
	}()

	select {
	case <-r.complete:
		fmt.Println("Completou a executação das tasks")
	case <-r.timeout:
		fmt.Println("Saindo da task pois ocorreu timeout")
	}
}

// func (r Runner) executaTasks2() error {
func (r Runner) executaTasks2() {
	defer wg.Done()

	fmt.Println("Executando Tasks2")
	for id, task := range r.tasks {
		if r.foiInterrompido() {
			fmt.Println("Saindo da execução pois foi interrompido")
			return
			// return errors.New("Task finalizado por motivos de interrupção")
		}

		fmt.Println("\nExecutando Task: ", id)
		task(id)
	}
	// return nil
}

func (r Runner) executaTasks() error {
	fmt.Println("Executando executaTasks")
	for id, task := range r.tasks {
		if r.foiInterrompido() {
			fmt.Println("Saindo da execução pois foi interrompido")
			return errors.New("Task finalizado por motivos de interrupção")
		}

		fmt.Println("\nExecutando Task: ", id)
		task(id)
	}
	return nil
}

func (r Runner) foiInterrompido() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r Runner) Teste() {
	fmt.Println("testado")
}

func (r Runner) ImprimeTasks() {
	for _, task := range r.tasks {
		fmt.Println("Imprimindo")
		task(10)
	}
}
