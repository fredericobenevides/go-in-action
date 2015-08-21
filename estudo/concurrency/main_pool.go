package main

import (
	"fmt"
	"go-in-action/estudo/concurrency/pool"
	"io"
	"log"
	"math/rand"
	"time"
)

type databaseConnection struct {
	id uint
}

func (dbConn *databaseConnection) newConnection(nextId uint) databaseConnection {
	return databaseConnection{id: nextId}
}

func (dbConn *databaseConnection) Close() error {
	fmt.Println("Fechando a conexão do Banco ", dbConn.id)
	return nil
}

func createConnection() (io.Closer, error) {
	var currentId uint
	dbConn := databaseConnection{}
	dbConn.newConnection(currentId)
	return &dbConn, nil
}

func main() {
	// pool, err := pool.New(nil, 0)
	// pool, err := pool.New(nil, 2)
	pool, err := pool.New(createConnection, 2)
	if err != nil {
		fmt.Println(err)
		fmt.Println(pool)
	}

	executeQuery(1, pool)
	// pool.Acquire()
	pool.Close()
}

func executeQuery(query int, p pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println("Não foi possível obter uma conexão do pool")
	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*databaseConnection).id)
}
