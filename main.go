package main

import (
	auth "testreq/authentication"
	"testreq/zachetka"
)

func main() {
	login := ""
	psws := ""
	client := auth.Login(login, psws)

	// user := helpers.GetUserProperties(client)

	// myteachers.GetTeachers(client, "2024-09-02T00:00:00.000Z", "2024-09-09T00:00:00.000Z", user.RaspID, user.TargetType)
	// result := predmets.GetCourses(client, user.UserId, user.Terms[4].TermId)
	// fmt.Println(result)
	// predmets.GetLessons(client, user.UserId, result[0].LessonsTypes[0].LessonTypeId)
	// raspisanie.GetRasp(client, "2024-09-09T00:00:00.000Z", user.RaspID, user.TargetType)
	// statistica.GetStatistic(client, user.UserId, user.Terms[4].TermId)
	// wfi.GetWiFiInfo(client)
	zachetka.GetZachetka(client)

}
