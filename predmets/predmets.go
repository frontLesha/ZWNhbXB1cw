package predmets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testreq/models"
)

func GetCourses(client *http.Client, studentId, termId string) []models.Predmet {
	var res interface{}
	response := []models.Predmet{}
	reqBody := fmt.Sprintf(`{"studentId":%s,"termId":%s}`, studentId, termId)
	req, err := http.NewRequest("POST", "https://ecampus.ncfu.ru/studies/GetCourses", bytes.NewReader([]byte(reqBody)))
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
	formatRes := res.(map[string]interface{})["courses"].([]interface{})
	for _, lesson := range formatRes {
		predmet := models.Predmet{Name: lesson.(map[string]interface{})["Name"].(string)}
		for _, properties := range lesson.(map[string]interface{})["LessonTypes"].([]interface{}) {
			predmet.LessonsTypes = append(predmet.LessonsTypes, models.LessonType{Name: properties.(map[string]interface{})["Name"].(string), LessonTypeId: fmt.Sprintf("%0.f", properties.(map[string]interface{})["Id"].(float64))})
		}
		response = append(response, predmet)
	}
	// formatResp, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(formatResp))
	return response
}

func GetLessons(client *http.Client, studentId, lessonTypeId string) []models.LessonPredmets {
	response := []models.LessonPredmets{}
	reqBody := fmt.Sprintf(`{"studentId":%s,"lessonTypeId":%s}`, studentId, lessonTypeId)
	req, err := http.NewRequest("POST", "https://ecampus.ncfu.ru/studies/GetLessons", bytes.NewReader([]byte(reqBody)))
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
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	formatResp, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formatResp))

	return response
}
