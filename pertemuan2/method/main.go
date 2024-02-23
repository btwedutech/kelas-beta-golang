package main

import "fmt"

type hewan struct {
	Nama  string
	Suara string
}

func (animal hewan) bunyi() {
	fmt.Println(animal.Suara)
}

func main() {
	Hew := hewan{
		Nama:  "kucing",
		Suara: "meooongggg",
	}

	Hew.bunyi()
}
