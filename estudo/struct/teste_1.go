package main

import (
	"fmt"
)

type informacao interface {
	dados()
}

type pessoa_fisica struct {
	cpf string
}

type pessoa_juridica struct {
	cnpj string
}

type pessoa struct {
	nome string
	pessoa_fisica
	pessoa_juridica
}

func (pf *pessoa_fisica) dados() {
	fmt.Println("\nImprimindo dados da pessoa fisica: ")
	fmt.Println(pf.cpf)
}

func (pj *pessoa_juridica) dados() {
	fmt.Println("\nImprimindo dados da pessoa jurídica: ")
	fmt.Println(pj.cnpj)
}

func (p *pessoa) dados() {
	fmt.Println("\nImprimindo dados da pessoa: ")
	fmt.Println(p.pessoa_fisica.cpf)
	fmt.Println(p.pessoa_juridica.cnpj)
}

func imprimir_informacao(i informacao) {
	i.dados()
}

func main() {
	pf := pessoa_fisica{
		cpf: "123.456.789.00",
	}

	pj := pessoa_juridica{
		cnpj: "000.000.011/1",
	}

	// pf.dados()
	// pj.dados()
	imprimir_informacao(&pf)
	imprimir_informacao(&pj)

	p := pessoa{
		pessoa_fisica:   pf,
		pessoa_juridica: pj,
		nome:            "Frederico",
	}

	fmt.Println("Testando de outra maneira")
	p.pessoa_fisica.dados()
	p.pessoa_juridica.dados()

	p.dados()
	imprimir_informacao(&p)
}
