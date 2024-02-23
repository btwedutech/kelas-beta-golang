package main

import "fmt"

func main() {
	var nama string = "adit"
	var namaPointer *string = &nama

	fmt.Println("nama asal :", nama)
	fmt.Println("ini nama dengan &: ", &nama)
	fmt.Println("ini variabel namaPointer: ", namaPointer)

	nama = "aditya"

	fmt.Println(*namaPointer)
	fmt.Println(namaPointer)

}

//todo implementasi pointer
//todo ubah data asal
//todo ubah data melalui pointer
