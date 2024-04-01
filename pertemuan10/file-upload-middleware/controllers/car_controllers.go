package controllers

import (
	"encoding/csv"
	"fmt"
	"sekolahbeta/pertemuan9/model"
	"sekolahbeta/pertemuan9/utils"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouteCars(app *fiber.App) {
	carsGroup := app.Group("/cars", CheckClient)
	app.Get("/", CheckClient)
	carsGroup.Get("/by-id/:id", GetCarByID)
	carsGroup.Post("/", InsertCarData)
	carsGroup.Post("/import-csv", CheckRole, ImportCsvFile)
}

func CheckClient(c *fiber.Ctx) error {
	client := string(c.Request().Header.Peek("Client"))
	if client == "Mobile" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "User Unauthorized",
	})
}

func CheckRole(c *fiber.Ctx) error {
	client := string(c.Request().Header.Peek("Role"))
	if client == "Admin" {
		return c.Next()
	}
	return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
		"message": "User Unauthorized",
	})
}

// file csv yang diload, adalah file cars_500.csv
// diambil dari pertemuan4/goroutine/

func ImportCsvFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": fmt.Sprintf("An Error occured, %s", err.Error()),
		})
	}

	files := form.File["csv_file"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": "A csv_file is required",
		})
	}

	for _, file := range files {
		fileContents, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "csv file is not valid",
			})
		}
		defer fileContents.Close()

		reader := csv.NewReader(fileContents)

		records, err := reader.ReadAll()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
				"message": "csv file cannot be loaded",
			})
		}

		err = utils.ImportCSVFile(records)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]any{
				"message": fmt.Sprintf("failed to import, error : %s", err.Error()),
			})
		}

		// var fileBuffer bytes.Buffer
		// _, err = io.Copy(&fileBuffer, fileContents)
		// if err != nil {
		// 	return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
		// 		"message": "csv file cannot be loaded",
		// 	})
		// }

		// os.WriteFile(file.Filename, fileBuffer.Bytes(), 0644)
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Success Import",
	})

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
