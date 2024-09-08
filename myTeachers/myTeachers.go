package myteachers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testreq/models"
)

func GetTeachers(client *http.Client, dateFirst, dateSecond, Id, targetType string) {
	var res []models.DayLessons
	var response []models.Teacher
	reqBody := fmt.Sprintf(`{"date":"%s","Id":%s,"targetType":%s}`, dateFirst, Id, targetType)
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
	for _, item := range res {
		for _, lesson := range item.Lessons {
			response = append(response, models.Teacher{Name: lesson.Teacher.Name, Lessons: []string{lesson.Discipline}})
		}
	}

	reqBody = fmt.Sprintf(`{"date":"%s","Id":%s,"targetType":%s}`, dateSecond, Id, targetType)
	req, err = http.NewRequest("POST", "https://ecampus.ncfu.ru/Schedule/GetSchedule", bytes.NewReader([]byte(reqBody)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range res {
		for _, lesson := range item.Lessons {
			response = append(response, models.Teacher{Name: lesson.Teacher.Name, Lessons: []string{lesson.Discipline}})
		}
	}
	result := map[string][]string{}
	for _, teacher := range response {
		result[teacher.Name] = append(result[teacher.Name], teacher.Lessons...)
	}

	for teacher, lessons := range result {
		result[teacher] = unique(lessons)
	}
	response = []models.Teacher{}
	for teacher, lessons := range result {
		response = append(response, models.Teacher{Name: teacher, Lessons: lessons})
	}

	formatResp, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formatResp))

}

func unique(slice []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueSlice []string

	for _, value := range slice {
		if _, exists := uniqueMap[value]; !exists {
			uniqueMap[value] = struct{}{}
			uniqueSlice = append(uniqueSlice, value)
		}
	}
	return uniqueSlice
}
