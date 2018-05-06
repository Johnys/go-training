package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi("-42"))
	fmt.Println(strconv.Itoa(20))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseInt("39", 10, 32))
	fmt.Println(strconv.FormatInt(20, 10))
}
