package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"qwit.dev/qwitravel-api/apis"
	"qwit.dev/qwitravel-api/models"
)

//
// Budapesti Közlekedési Központ (BKK) API
//

func BKKGetStopSchedule(stopIds []string, date string) (int, models.BKKStopSchedule) {
	if len(stopIds) == 0 {
		return fiber.ErrBadRequest.Code, models.BKKStopSchedule{}
	}

	if date == "" {
		currentTime := time.Now()
		date = strings.ReplaceAll(currentTime.Format("2006#01#02"), "#", "")
	}

	var finalData models.BKKStopSchedule

	client := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		apis.BKKStopSchedule+
			"?stopId="+stopIds[0]+
			"&date="+date+
			"&onlyDepartures=false&appVersion=1.0.0&version=4&includeReferences=routes,alerts"+
			"&key="+os.Getenv("BKK_API_KEY"),
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError.Code, models.BKKStopSchedule{}
	}

	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError.Code, models.BKKStopSchedule{}
	}

	if resp.StatusCode != 200 {
		return resp.StatusCode, models.BKKStopSchedule{}
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError.Code, models.BKKStopSchedule{}
	}

	err = json.Unmarshal(body, &finalData)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError.Code, models.BKKStopSchedule{}
	}

	return fiber.StatusOK, finalData
}
