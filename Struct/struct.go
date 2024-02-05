package main

import "fmt"

type user struct {
	Name string
	Age  uint8
}

func main() {
	fmt.Println("Structs")

	var u user
	u.Age = 24
	u.Name = "Izael"

	fmt.Println(u.Name, u.Age)

	u2 := user{"Julia", 24}

	fmt.Println(u2)
}
