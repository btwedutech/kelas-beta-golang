package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sekolahbeta/hacker/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {
	started := time.Now()
	file, err := openFile("json/cars_100000.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	defer file.Close()

	cars, err := loadFile(file)
	if err != nil {
		return err
	}

	jsonData, err := convertCarsToJson(cars)
	if err != nil {
		return err
	}

	fmt.Println("[Tanpa Goroutine]", time.Since(started))

	return c.Send(jsonData)
}

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func loadFile(file *os.File) ([]model.Car, error) {
	var cars []model.Car
	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		return cars, err
	}

	for _, car := range data {
		cars = append(cars, model.Car{
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

	return cars, nil
}

func convertCarsToJson(data []model.Car) ([]byte, error) {
	return json.Marshal(data)
}
