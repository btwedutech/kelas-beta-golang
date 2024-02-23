package main

import "fmt"

func hello() {
	fmt.Println("World")
}

func konversiMataUangUSD(uang int) int {
	var hasil int

	hasil = uang * 15000

	return hasil
}

func konversiMataUang(uang int, currency string) (int, string) {
	var hasil int

	switch currency {
	case "USD":
		hasil = uang * 15000
	case "JPY":
		hasil = uang * 300
	default:
		hasil = 0
	}

	return hasil, currency
}

func kalkulator(operator string, angka ...int) int {
	var hasil int
	for i := 0; i < len(angka); i++ {
		if operator == "+" {
			hasil += angka[i]
		} else if operator == "-" {
			hasil -= angka[i]
		}
	}

	return hasil
}

func main() {
	hello()
	hasilKonversi := konversiMataUangUSD(5)
	fmt.Println(hasilKonversi)

	hasilKonversi1, _ := konversiMataUang(5000, "JPY")
	fmt.Println(hasilKonversi1)

	angkaInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	hasilHitungKalkulator := kalkulator("+", angkaInput...)
	fmt.Println(hasilHitungKalkulator)

}

//todo void func
//todo func konversi mata uang 1 return
//todo func konversi mata uang 2 return
//todo func variadic calculator (...)
