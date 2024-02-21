package model

import "gorm.io/gorm"

type PaperNew struct {
	gorm.Model
	Auther    string `gorm:"type:varchar(20);not null;primaryKey"`
	Name      string `gorm:"type:varchar(20);not null;primaryKey"`
	Titles    string `gorm:"type:varchar(20);not null"`
	Questions string `gorm:"type:TEXT(65535);not null"`
	Number    string `gorm:"type:varchar(6);not null;unique"`
}

type PaperUser struct {
	gorm.Model
	Acount    string `gorm:"type:varchar(20);not null;primaryKey"`
	Questions string `gorm:"type:TEXT(65535);not null"`
	Number    string `gorm:"type:varchar(6);not null;primaryKey"`
	Times     int    `gorm:"type:int;not null;primaruKey"`
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
	Kind         string     `json:"kind"`
	Questionaire string     `json:"question"`
	Operation    Operations `json:"operation"`
}
type Paper struct {
	Auther        string           `json:"auther"`
	Name          string           `json:"name"`
	Title         []string         `json:"title"`
	Questionnaire []QuestionnaireT `json:"questionnaire"`
}
type UploadPaper struct {
	Acount        string           `json:"acount"`
	Number        string           `json:"number"`
	Questionnaire []QuestionnaireT `json:"questionnaire"`
	Willing       bool             `json:"willing"`
}
