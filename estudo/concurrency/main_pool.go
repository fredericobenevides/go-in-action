package main

import (
	"fmt"
	"go-in-action/estudo/concurrency/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type databaseConnection struct {
	id uint
}

func (dbConn *databaseConnection) newConnection(nextId uint) databaseConnection {
	fmt.Println("Criando uma nova conexão com id: ", nextId)
	return databaseConnection{id: nextId}
}

func (dbConn *databaseConnection) Close() error {
	fmt.Println("Fechando a conexão do Banco ", dbConn.id)
	return nil
}

func createConnection() (io.Closer, error) {
	var currentId uint
	dbConn := databaseConnection{}
	dbConn.newConnection(1)
	currentId = currentId + 1
	return &dbConn, nil
}

func createConnection2() (io.Closer, error) {
	atomic.AddInt32(&idDatabase, 1)
	return &databaseConnection{id: uint(idDatabase)}, nil
}

const (
	maxgoroutines = 10
	// sizeResources = 2
	sizeResources = 2
)

var idDatabase int32

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxgoroutines)

	// pool, err := pool.New(nil, 0)
	// pool, err := pool.New(nil, 2)
	// pool, err := pool.New(createConnection, sizeResources)
	pool, err := pool.New(createConnection2, sizeResources)
	if err != nil {
		fmt.Println(err)
		fmt.Println(pool)
	}

	for i := 0; i < maxgoroutines; i++ {
		go func(query int) {
			executeQuery(query, pool)
			wg.Done()
		}(i)

	}

	wg.Wait()

	// pool.Acquire()
	pool.Close()
}

func executeQuery(query int, p pool.Pool) {
	// conn, err := p.Acquire()
	conn, err := p.AcquireSemNovaConexao()
	if err != nil {
		log.Println("Não foi possível obter uma conexão do pool")
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*databaseConnection).id)
}
