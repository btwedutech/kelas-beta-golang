package utils

import (
	"fmt"
	"sekolahbeta/pertemuan9/config"
	"sekolahbeta/pertemuan9/model"
	"strconv"
	"time"
)

func GetCarsList() ([]model.Car, error) {
	var cars model.Car
	return cars.GetAll(config.Mysql.DB)
}

func GetCarByID(id uint) (model.Car, error) {
	cars := model.Car{
		Model: model.Model{
			ID: id,
		},
	}
	return cars.GetByID(config.Mysql.DB)
}

func InsertCarData(data model.Car) (model.Car, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func ImportCSVFile(records [][]string) error {

	cars := []model.Car{}

	for index, car := range records {
		if index == 0 {
			continue
		}

		sellingPrice, err := strconv.Atoi(car[14])
		if err != nil {
			sellingPrice = 0
		}

		cars = append(cars, model.Car{
			Nama:         car[3],
			Tipe:         car[5],
			Tahun:        car[1],
			Color:        car[10],
			Condition:    car[8],
			UUID:         car[0],
			SellingPrice: sellingPrice,
		})
	}

	for _, car := range cars {
		car.CreatedAt = time.Now()
		car.UpdatedAt = time.Now()

		// Implementasi mekanisme Upsert (Update Insert)
		// rubah GetById menjadi GetBySpecific
		res, err := car.GetBySpecific(config.Mysql.DB)
		// disini kita akan cek apakah terdapat error
		// atau tidak
		if err != nil {
			// kita cek kondisi error "record not found"
			// jika ada error diluar not found, artinya
			// ada error dengan db, kita tidak bisa lanjut
			if err.Error() != "record not found" {
				return err
				// lanjut jika ketemu error record not found
				// kita tinggal insert aja datanya
			} else {
				err = car.Create(config.Mysql.DB)
				if err != nil {
					return fmt.Errorf(
						"failed to import data, error :%s",
						err.Error())
				}
			}
			// kalau tidak ada error,
			// kita update aja datanya
		} else {

			car.ID = res.ID
			car.CreatedAt = res.CreatedAt
			err = car.UpdateOneByID(config.Mysql.DB)

			if err != nil {
				return fmt.Errorf("failed to update data, error :%s", err.Error())
			}

		}
	}
	return nil
}
