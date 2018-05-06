package main

import (
	"fmt"
)

func comPonteiro(numero *int) {
	fmt.Println(numero)
	*numero = 20
}

func semPonteiro(numero int) {
	numero = 20
}

func main() {
	i := 10
	fmt.Println(i)
	semPonteiro(i)
	fmt.Println(i)
	comPonteiro(&i)
	fmt.Println(i)
}
