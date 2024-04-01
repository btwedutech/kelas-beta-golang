package controllers

import (
	"sekolahbeta/pertemuan9/model"
	"sekolahbeta/pertemuan9/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCars(app *fiber.App) {
	carsGroup := app.Group("/cars")
	carsGroup.Get("/", GetCarsList)
	carsGroup.Get("/by-id/:id", GetCarByID)
	carsGroup.Post("/", InsertCarData)
}

func InsertCarData(c *fiber.Ctx) error {
	type AddCarRequest struct {
		Nama  string `json:"nama" valid:"required,type(string)"`
		Tipe  string `json:"tipe" valid:"required"`
		Tahun string `json:"tahun" valid:"optional"`
	}

	req := new(AddCarRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Body not valid",
			})
	}

	isValid, err := govalidator.ValidateStruct(req)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	car, errCreateCar := utils.InsertCarData(model.Car{})

	if errCreateCar != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateCar.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Success Insert Data",
		"car":     car,
	})

}

func GetCarsList(c *fiber.Ctx) error {
	carsData, err := utils.GetCarsList()
	if err != nil {
		logrus.Error("Error on get cars list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    carsData,
			"message": "Success",
		},
	)
}

func GetCarByID(c *fiber.Ctx) error {
	carId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	carData, err := utils.GetCarByID(uint(carId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get car data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    carData,
			"message": "Success",
		},
	)
}
