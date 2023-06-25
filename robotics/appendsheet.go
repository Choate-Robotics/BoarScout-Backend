package robotics

import (
	"context"
	"github.com/gofiber/fiber/v2"
	_ "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"net/http"
)

func AppendSheet(c *fiber.Ctx) error {

	var requestArgs struct {
		Data []string `json:"data"`
	}

	if err := c.BodyParser(&requestArgs); err != nil {
		log.Print(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Internal Server Error"})
	}

	data, err := ioutil.ReadFile("./cloud/creds.json")
	if err != nil {
		log.Print(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Internal Server Error"})
	}

	config, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Print(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Internal Server Error"})
	}

	client := config.Client(context.Background())
	service, err := sheets.New(client)
	if err != nil {
		log.Print(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "Internal Server Error"})
	}

	spreadsheetId := "1emEK4uK_xTWOG7vOHYPvl_5TrpsbbT10Bm93kxiKVD4"
	range2 := "Sheet1!A1"

	var values [][]interface{}
	row := []interface{}{}

	for _, value := range requestArgs.Data {
		row = append(row, value)
	}

	values = append(values, row)

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	resp, err := service.Spreadsheets.Values.Append(spreadsheetId, range2, valueRange).ValueInputOption("RAW").Do()
	if err != nil || resp.HTTPStatusCode != 200 {
		log.Print(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": 400, "message": "Bad data, u need some copium"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": 200, "message": "Success"})
}
