package controllers

import (
	"sekolahbeta/pertemuan8-2/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCars(app *fiber.App) {
	carsGroup := app.Group("/cars")
	carsGroup.Get("/", GetCarsList)
	carsGroup.Get("/by-id/:id", GetCarByID)
	// carsGroup.Post("/")
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
