package model

import "gorm.io/gorm"

type PaperNew struct {
	gorm.Model
	Auther    string `gorm:"type:varchar(20);not null;primaryKey"`
	Name      string `gorm:"type:varchar(20);not null;primaryKey"`
	Titles    string `gorm:"type:varchar(20);not null"`
	Questions string `gorm:"type:TEXT(65535);not null"`
}
type Operations struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
}
type QuestionnaireT struct {
	Branch       string     `json:"branch"`
	Id           int        `json:"id"`
	Questionaire string     `json:"question"`
	Operation    Operations `json:"operation"`
}
type Paper struct {
	Auther        string           `json:"auther"`
	Name          string           `json:"name"`
	Title         []string         `json:"title"`
	Questionnaire []QuestionnaireT `json:"questionnaire"`
	Kind          string           `json:"kind"`
}
