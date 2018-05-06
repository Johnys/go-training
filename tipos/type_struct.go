package main

import "fmt"

type Pessoa struct {
	nome  string
	idade uint
}

func (pessoa *Pessoa) getNome() string {
	return pessoa.nome
}

func main() {
	p := Pessoa{"Johnys", 10}
	p2 := Pessoa{nome: "Johnys 2", idade: 20}
	fmt.Println(p.getNome())
}
