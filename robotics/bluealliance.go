package robotics

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"sort"
	"strings"
)

func GetTeams(c *fiber.Ctx) error {

	database := make([]map[string]interface{}, 0)

	appendDB := func(index int, red1, red2, red3, blue1, blue2, blue3 string) {
		teams := []map[string]interface{}{
			{"text": blue1, "index": 0, "color": "#307fe8"},
			{"text": blue2, "index": 1, "color": "#307fe8"},
			{"text": blue3, "index": 2, "color": "#307fe8"},
			{"text": red1, "index": 3, "color": "#B50303"},
			{"text": red2, "index": 1, "color": "#B50303"},
			{"text": red3, "index": 2, "color": "#B50303"},
		}

		data := map[string]interface{}{
			"match": index + 1,
			"index": index,
			"teams": teams,
		}

		database = append(database, data)
	}

	getBlueAllianceData := func() {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://www.thebluealliance.com/api/v3/event/2023hop/matches", nil)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"data":   "Internal Server Error",
			})
			return
		}

		req.Header.Set("X-TBA-Auth-Key", os.Getenv("BL_API_KEY"))
		res, err := client.Do(req)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"data":   "Internal Server Error",
			})
			return
		}
		defer res.Body.Close()

		var matches []map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&matches)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": 500,
				"data":   "Internal Server Error",
			})
			return
		}

		for _, match := range matches {
			blueTeams := match["alliances"].(map[string]interface{})["blue"].(map[string]interface{})["team_keys"].([]interface{})
			redTeams := match["alliances"].(map[string]interface{})["red"].(map[string]interface{})["team_keys"].([]interface{})

			appendDB(
				int(match["match_number"].(float64))-1,
				strings.TrimPrefix(redTeams[0].(string), "frc"),
				strings.TrimPrefix(redTeams[1].(string), "frc"),
				strings.TrimPrefix(redTeams[2].(string), "frc"),
				strings.TrimPrefix(blueTeams[0].(string), "frc"),
				strings.TrimPrefix(blueTeams[1].(string), "frc"),
				strings.TrimPrefix(blueTeams[2].(string), "frc"),
			)
		}
	}

	getBlueAllianceData()

	sort.SliceStable(database, func(i, j int) bool {
		return database[i]["index"].(int) < database[j]["index"].(int)
	})

	return c.JSON(fiber.Map{
		"status": 200,
		"data":   database,
	})
}
