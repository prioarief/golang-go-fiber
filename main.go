package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type ResponseStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users []User = []User{
	{
		Id:   1,
		Name: "Kylian Mbappe",
	},
}

func main() {
	// fiber instance
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(helmet.New())

	app.Get("/health-check", func(c *fiber.Ctx) error {
		// return c.SendString("Server running well")
		return c.Status(200).JSON(fiber.Map{
			"code":    1,
			"message": "All is well",
			"data":    nil,
		})
		// return c.Status(200).JSON(ResponseStruct{
		// 	Code:    1,
		// 	Message: "Server running well",
		// 	Data:    nil,
		// })
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		u := new(User)

		if err := c.BodyParser(u); err != nil {
			return c.Status(400).JSON(ResponseStruct{
				Code:    0,
				Message: err.Error(),
			})
		}

		newUser := User{
			Id:   len(users) + 1,
			Name: u.Name,
		}

		users = append(users, newUser)

		return c.Status(201).JSON(newUser)
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(ResponseStruct{
			Code:    1,
			Message: "List Users",
			Data:    users,
		})
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(ResponseStruct{
				Code:    0,
				Message: "Id must be number",
			})
		}

		for _, user := range users {
			if user.Id == id {
				return c.Status(200).JSON(ResponseStruct{
					Code:    1,
					Message: "Get User Detail",
					Data:    user,
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(ResponseStruct{
			Code:    0,
			Message: "Data not found",
		})
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(ResponseStruct{
				Code:    0,
				Message: "Id must be number",
			})
		}

		u := new(User)

		if err := c.BodyParser(u); err != nil {
			return c.Status(400).JSON(ResponseStruct{
				Code:    0,
				Message: err.Error(),
			})
		}

		for i, user := range users {
			if user.Id == id {
				users[i].Name = u.Name
				user.Name = u.Name

				return c.Status(200).JSON(ResponseStruct{
					Code:    1,
					Message: "Update User",
					Data:    user,
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(ResponseStruct{
			Code:    0,
			Message: "Data not found",
		})
	})

	// app.Post("/form", func(c *fiber.Ctx) error {

	// })

	log.Fatal(app.Listen(":3000"))
}
