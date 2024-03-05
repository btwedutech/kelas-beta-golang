package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	defer fmt.Println("guys")
	defer fmt.Println("haha")

	fmt.Println("world")
}
