package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Car struct {
	ID           string `json:"uuid"`
	Year         string `json:"year"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Trim         string `json:"trim"`
	Body         string `json:"body"`
	Transmission string `json:"trasmission"`
	State        string `json:"state"`
	Condition    string `json:"condition"`
	Odometer     string `json:"odometer"`
	Color        string `json:"color"`
	Interior     string `json:"interior"`
	Seller       string `json:"seller"`
	Mmr          string `json:"mmr"`
	SellingPrice string `json:"selling_price"`
	SaleDate     string `json:"sale_date"`
}

func main() {
	err := os.RemoveAll("json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Removing json directory")

	err = os.Mkdir("json", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating json directory")

	fileCsv, err := os.Open("cars_100000.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer fileCsv.Close()

	reader := csv.NewReader(fileCsv)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	cars := csvToStruct(records)

	startedAt := time.Now()

	wg := sync.WaitGroup{}

	var channel = make(chan Car)

	jml := 3
	fmt.Println("Menjalankan ", jml, "process goroutine")

	for i := 0; i < jml; i++ {
		wg.Add(1)
		go process(channel, &wg)
	}

	// wg.Add(2)

	// go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)

	for _, car := range cars {
		channel <- car
	}

	close(channel)

	wg.Wait() //block

	fmt.Println("Success")
	fmt.Println(time.Since(startedAt))

}

func csvToStruct(records [][]string) []Car {
	cars := []Car{}

	for _, car := range records {
		cars = append(cars, Car{
			ID:           car[0],
			Year:         car[1],
			Make:         car[2],
			Model:        car[3],
			Trim:         car[4],
			Body:         car[5],
			Transmission: car[6],
			State:        car[7],
			Condition:    car[8],
			Odometer:     car[9],
			Color:        car[10],
			Interior:     car[11],
			Seller:       car[12],
			Mmr:          car[13],
			SellingPrice: car[14],
			SaleDate:     car[15],
		})
	}

	return cars
}

func convertToJson(car Car) []byte {
	encoded, err := json.MarshalIndent(car, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return encoded
}

func saveJsonToFile(encoded []byte, name string) {
	file, err := os.Create(fmt.Sprintf("json/%s.json", name))
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(encoded)
	if err != nil {
		fmt.Println(err)
	}

}

func process(channel chan Car, wg *sync.WaitGroup) {
	for data := range channel {
		encoded := convertToJson(data)
		saveJsonToFile(encoded, data.ID)
	}

	wg.Done()
}
