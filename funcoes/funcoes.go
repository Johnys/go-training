package main

import "fmt"

func multiplosRetornos() (int, int) {
	var i, j int = 10, 20
	return i, j
}

func testMultiplosRetornos() {
	_, b := multiplosRetornos()
	fmt.Println("Os retornos foram:", b)
}

func incrementa() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testIncrementa() {
	inc := incrementa()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func main() {
	//testMultiplosRetornos()
	testIncrementa()
}
