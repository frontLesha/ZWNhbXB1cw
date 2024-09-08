package models

type Predmet struct {
	Name         string       `json:"Name"`
	LessonsTypes []LessonType `json:"LessonType"`
}

type LessonType struct {
	Name         string `json:"lessonName"`
	LessonTypeId string `json:"lessonTypeId"`
}

type LessonPredmets struct {
	Room       string `json:"Room"`
	KodPr      int    `json:"Kod_pr"`
	Name       string `json:"Name"`
	Date       string `json:"Date"`
	Attendance int    `json:"Attendance"`
	GradeText  string `json:"GradeText"`
	LessonName string `json:"LessonName"`
}

type Teacher struct {
	Name    string   `json:"Name"`
	Id      int      `json:"Id"`
	Lessons []string `json:"Lessons"`
}

type LessonRasp struct {
	Discipline      string  `json:"Discipline"`
	TimeBegin       string  `json:"TimeBegin"`
	TimeEnd         string  `json:"TimeEnd"`
	PairNumberStart int     `json:"PairNumberStart"`
	Auditore        Aud     `json:"Aud"`
	LessonType      string  `json:"LessonType"`
	Teacher         Teacher `json:"Teacher"`
}

type Aud struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type DayLessons struct {
	Lessons []LessonRasp `json:"Lessons"`
	WeekDay string       `json:"WeekDay"`
}

type Term struct {
	TermId  string
	TermNum string
}

type UserProperties struct {
	UserId     string `json:"UserId"`
	Terms      []Term `json:"Terms"`
	RaspID     string `json:"RaspId"`
	TargetType string `json:"TargetType"`
}
