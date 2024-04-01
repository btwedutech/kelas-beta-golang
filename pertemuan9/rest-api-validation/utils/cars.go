package utils

import (
	"sekolahbeta/pertemuan9/config"
	"sekolahbeta/pertemuan9/model"
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
