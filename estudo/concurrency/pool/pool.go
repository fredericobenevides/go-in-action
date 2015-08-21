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
	return Pool{
		connection: fn,
		resources:  make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	fmt.Println("Fechando o pool")
	p.closed = true

}

func (p *Pool) Acquire() (io.Closer, error) {
	fmt.Println("Obtendo uma conexão com o pool")
	return p.connection()
}
