package robotics

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
)

func CollectData(c *fiber.Ctx) error {

	requestArgs := new(struct {
		Name string `json:"name"`
	})

	err := c.BodyParser(requestArgs)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	payload := map[string]interface{}{
		"content":     "LOG: " + requestArgs.Name + " Opened BoarScout",
		"embeds":      nil,
		"attachments": []string{},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	resp, err := http.Post(os.Getenv("API_URL"),
		"application/json",
		bytes.NewBuffer(payloadBytes),
	)

	if err != nil || resp.StatusCode != 204 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Sus...reported to the FBI",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "eW91IGhhdmUgMCByaXp6LCBzdG9wIGRlY29kaW5nIGJhc2U2NCBmb3IgZnVuLg==",
	})
}
