package statistica

import (
	"fmt"
	"net/http"
	"testreq/models"
	"testreq/predmets"
)

func GetStatistic(client *http.Client, studentId string, termId string) {
	result := []models.LessonPredmets{}
	predmetList := predmets.GetCourses(client, studentId, termId)
	for _, predmet := range predmetList {
		for _, item := range predmet.LessonsTypes {
			lessons := predmets.GetLessons(client, studentId, item.LessonTypeId)
			for _, lesson := range lessons {
				lesson.LessonName = predmet.Name
				result = append(result, lesson)
			}
		}
	}
	nList := []models.LessonPredmets{}
	yList := []models.LessonPredmets{}
	for _, res := range result {
		if res.Attendance == 0 {
			nList = append(nList, res)
		}
		if res.Attendance == 2 {
			yList = append(yList, res)
		}
	}
	fmt.Println(nList)
	fmt.Println("num N ", len(nList))
	fmt.Println("num Y ", len(yList))
}
