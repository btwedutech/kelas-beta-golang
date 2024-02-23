package main

import "fmt"

func main() {
	urutanAngka := map[string]int{
		"satu":  1,
		"dua":   2,
		"tiga":  3,
		"empat": 4,
	}

	for key, val := range urutanAngka {
		fmt.Println("ini value apa: ", key)
		fmt.Println("value: ", val)

	}
}
