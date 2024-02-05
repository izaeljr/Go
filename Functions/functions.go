package main

import "fmt"

func sum(n1 int16, n2 int16) int16 {
	return n1 + n2
}

func calc(n1, n2 int16) (int16, int16) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := sum(10, 42)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Text")
	println(result)

	resultSum, resultSub := calc(15, 23)
	println(resultSum, resultSub)
}
