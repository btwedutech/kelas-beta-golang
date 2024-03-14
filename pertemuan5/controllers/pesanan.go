package controllers

import (
	"encoding/json"
	"fmt"
	"sekolahbeta/hacker/model"

	"github.com/gofiber/fiber/v2"
)

func GetPesanan(c *fiber.Ctx) error {

	data := []model.Pesanan{
		{
			ID:   "1",
			Name: "nama-1",
			Meja: "meja-1",
		},
		{
			ID:   "2",
			Name: "nama-2",
			Meja: "meja-2",
		},
		{
			ID:   "3",
			Name: "nama-3",
			Meja: "meja-3",
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.Send(jsonData)
}
