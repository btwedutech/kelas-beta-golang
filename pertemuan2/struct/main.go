package main

import "fmt"

type Person struct {
	Name      string
	DateBirth string
	Weight    string
	Mail      []Email
}

type Email struct {
	To   string
	Name string
}

func main() {
	//struct
	person1 := Person{
		Name:      "saya",
		DateBirth: "01-01-1998",
		Weight:    "60",
		Mail: []Email{
			{
				To:   "adit@gmail.com",
				Name: "adit",
			},
			{
				To:   "adit1@gmail.com",
				Name: "adit",
			},
		},
	}

	fmt.Println(person1)

	persons := []Person{
		{Name: "saya1", DateBirth: "01-01-1998", Weight: "60"},
		{Name: "saya1", DateBirth: "01-01-1998", Weight: "60"},
		{Name: "saya1", DateBirth: "01-01-1998", Weight: "60"},
		{Name: "saya1", DateBirth: "01-01-1998", Weight: "60"},
	}

	fmt.Println(persons)

}

//Struct
//Slice Struct
//Slice Struct didalam Slice
