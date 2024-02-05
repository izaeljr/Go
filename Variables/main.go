package main

import "fmt"

func main() {
	var test string = "Teste"
	test2 := "Teste 2"

	fmt.Println(test)
	fmt.Println(test2)

	var (
		test3 string = "Teste 3"
		test4 int    = 4
	)

	fmt.Println(test3)
	fmt.Println(test4)

	test, test2 = test2, test

	fmt.Println(test)
	fmt.Println(test2)
}
