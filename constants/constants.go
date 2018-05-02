package main

import "fmt"

const inteiro int = 10
const boolean bool = false
const texto string = "teste"

const (
	segunda = iota
	terca
	quarta
	quinta
	sexta
)

func imprimeConstantesSimples() {
	fmt.Println("O valor das constantes são: ", inteiro, boolean, texto)
}

func imprimeConstantesIOTA() {
	fmt.Println("O valor das constantes da semana são: ", segunda, terca, quarta, quinta, sexta)
}

func main() {
	imprimeConstantesSimples()
	imprimeConstantesIOTA()
}
