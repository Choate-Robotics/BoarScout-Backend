package robotics

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func AuthenticateUser(c *fiber.Ctx) error {

	requestArgs := new(struct {
		Code string `json:"code"`
	})

	err := c.BodyParser(requestArgs)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	if requestArgs.Code != os.Getenv("ACCESS_CODE") {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "sus o-o",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Authentication successful",
	})

}
