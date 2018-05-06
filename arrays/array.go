package main

import (
	"fmt"
)

func main() {
	numeros := [6]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(numeros); i++ {
		fmt.Println("Posição numero: ", i, &numeros[i])
	}
}
