package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	m          sync.Mutex
	connection func() (io.Closer, error)
	resources  chan io.Closer
	closed     bool
}

func New(fn func() (io.Closer, error), size int) (Pool, error) {
	if size <= 0 {
		// return nil, errors.New("Valor do pool está pequeno")
		return Pool{}, errors.New("Valor do pool está pequeno")
	}
	// return nil, nil
	// return Pool{
	// 	connection: fn,
	// 	resources:  make(chan io.Closer, size),
	// }, nil
	pool := Pool{
		connection: fn,
		resources:  make(chan io.Closer, size),
	}

	//inicializa o pool com a quantidade de resource necessário
	go func(size int) {
		fmt.Printf("Adicionado %d conexões no pool\n", size)
		for i := 0; i < 2; i++ {
			c, err := pool.connection()
			if err != nil {
				fmt.Println("New: Ocorreu um problema ao obter a conexão")
			}
			pool.resources <- c
		}
	}(size)

	return pool, nil
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	fmt.Println("Fechando o pool")

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}

func (p *Pool) Acquire() (io.Closer, error) {
	fmt.Println("Obtendo uma conexão com o pool")
	select {
	case r, ok := <-p.resources:
		fmt.Println("Acquire: usando pool compartilhado")
		if !ok {
			return nil, errors.New("pool closed")
		}
		return r, nil

	default:
		fmt.Println("Acquire: usando UM NOVO pool")
		return p.connection()
	}
}

func (p *Pool) AcquireSemNovaConexao() (io.Closer, error) {
	fmt.Println("Obtendo uma conexão com o pool")
	for {
		r, ok := <-p.resources
		if !ok {
			return nil, errors.New("pool closed")
		}
		fmt.Println("Retornando a conexão do pool")
		return r, nil
	}
}

func (p *Pool) Release(resource io.Closer) {
	if p.closed {
		resource.Close()
		return
	}

	select {
	case p.resources <- resource:
		fmt.Println("Release: Colocando em queue")
	default:
		// resources já está cheio
		fmt.Println("Release: Closing")
	}

}
