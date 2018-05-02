package main

import "fmt"

func variaveisInteiras() {
	var x int = 3
	var y int = 10
	fmt.Println("O valor da soma é: ", x+y)
}

func multiplasVariaveisEmUmaLinha() {
	var x, y int = 4, 11
	fmt.Println("O valor da soma é: ", x+y)
}

func valorPadraoVariavel() {
	var x int
	fmt.Println("O valor de x é: ", x)
}

func declararEInstanciarVariavel() {
	x := 10
	fmt.Println("O valor de x é:", x)
}

func main() {
	variaveisInteiras()
	multiplasVariaveisEmUmaLinha()
	valorPadraoVariavel()
	declararEInstanciarVariavel()
}
