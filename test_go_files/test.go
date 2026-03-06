package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Hi! I'm your new device! Pls write username:")
	fmt.Scan(&name)
	fmt.Println("Hello,", name, ",nice to meet you!")

	var age int
	fmt.Println("Pls write your age:")
	fmt.Scan(&age)
	fmt.Printf("OK! Username - %s, age - %d\n", name, age)
}
