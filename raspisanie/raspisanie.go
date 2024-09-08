package raspisanie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testreq/models"
)

func GetRasp(client *http.Client, date string, Id, targetType string) {
	var res []models.DayLessons
	reqBody := fmt.Sprintf(`{"date":"%s","Id":%s,"targetType":%s}`, date, Id, targetType)
	req, err := http.NewRequest("POST", "https://ecampus.ncfu.ru/Schedule/GetSchedule", bytes.NewReader([]byte(reqBody)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}
	formatResp, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formatResp))
}
