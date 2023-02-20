package main

import (
	"fmt"

	"github.com/Enzuigiri/GoFiberAPITest/test"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Calc(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println(test.Print("enzu"))
	fmt.Println(Calc(10))
	app := fiber.New()
	app.Use(logger.New())

	// Give Response When At /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	// Handle error
	if err != nil {
		panic(err)
	}
}
