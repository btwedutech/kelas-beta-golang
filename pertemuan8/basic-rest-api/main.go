package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	carsRoute := app.Group("/cars")
	carsRoute.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(map[string]any{
				"data": []any{
					"mobil 1", "mobil 2",
				},
				"message": "Success",
			})
	})

	carsRoute.Get("/by-id/:id",
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).
				JSON(map[string]any{
					"data": map[string]any{
						"id": c.Params("id"),
					},
					"message": "Success",
				})
		})

	carsRoute.Get("/search",
		func(c *fiber.Ctx) error {

			return c.Status(fiber.StatusOK).
				JSON(map[string]any{
					"data": map[string]any{
						"id": c.Query("id"),
					},
					"message": "Success",
				})
		})

	carsRoute.Post("/", func(c *fiber.Ctx) error {
		type CreateCarRequest struct {
			Brand string `json:"brand"`
			Type  string `json:"type"`
			Color string `json:"color"`
		}

		req := new(CreateCarRequest)

		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(map[string]any{
					"message": "Body request not valid",
				})
		}

		return c.Status(fiber.StatusCreated).
			JSON(map[string]any{
				"data": map[string]any{
					"brand": req.Brand, "type": req.Type,
					"color": req.Color,
				},
				"message": "Success Insert Data",
			})
	})

	carsRoute.Put("/brand/:name", func(c *fiber.Ctx) error {
		type UpdateCarRequest struct {
			ID    string `json:"id"`
			Color string `json:"color"`
			Type  string `query:"type"`
		}

		req := new(UpdateCarRequest)

		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(map[string]any{
					"message": "Body request not valid",
				})
		}

		if err := c.QueryParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(map[string]any{
					"message": "Body request not valid",
				})
		}

		return c.Status(fiber.StatusCreated).
			JSON(map[string]any{
				"data": map[string]any{
					"brand": c.Params("name"), "type": req.Type,
					"color": req.Color, "id": req.ID,
				},
				"message": "Success Insert Data",
			})
	})

	app.Listen(":3000")
}
