package controller

import (
	"PaSer/common"
	"PaSer/model"
	"PaSer/response"
	"PaSer/util"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddPaper(c *gin.Context) {
	var paper model.Paper
	c.ShouldBindJSON(&paper)
	db := common.Con_Db_asp()
	Name := paper.Name
	Auther := paper.Auther
	Title := strings.Join(paper.Title, ",")
	questionnaireJson, err := json.Marshal(paper.Questionnaire)
	if err != nil {
		response.FalseRe(c, fmt.Sprintf("Error encoding questionnaire: %s", err), nil)
		return
	}
	Questions := string(questionnaireJson)
	PaperNewUp := model.PaperNew{
		Auther:    Auther,
		Name:      Name,
		Titles:    Title,
		Questions: Questions,
		Number:    util.RandomString(6),
	}
	result := db.Where(&model.PaperNew{Auther: Auther, Name: Name}).Limit(1).Find(&PaperNewUp)
	if result.Error != nil {
		response.FalseRe(c, fmt.Sprintf("err:%v", result.Error), nil)
		return
	}
	Cp := db.Create(&PaperNewUp)
	if Cp.Error != nil {
		response.FalseRe(c, fmt.Sprintf("err:%v", Cp.Error), nil)
		return
	}
	response.Response(c, 200, 200, gin.H{"code": PaperNewUp.Number}, "upload successfully")
}
func UploadPaper(c *gin.Context) {
	var paper_upload model.UploadPaper
	c.ShouldBindJSON(&paper_upload)
	acount := paper_upload.Acount
	number := paper_upload.Number
	questions := paper_upload.Questionnaire
	willing := paper_upload.Willing
	db, err_db := common.Create_Uploader_DB()
	times := findTimes(db, acount, number)
	fmt.Println(willing, times)
	if !willing && times != 0 {
		response.FalseRe(c, "you had upload once", nil)
		return
	}
	if err_db != nil {
		response.FalseRe(c, "upload failed", gin.H{"err_msg": err_db})
	}
	jsonbyte, err := json.Marshal(questions)
	if err != nil {
		response.FalseRe(c, "字符串转换错误，问卷上传出错", nil)
		return
	}
	json_question := string(jsonbyte)
	paper_db := model.PaperUser{
		Acount:    acount,
		Number:    number,
		Questions: json_question,
		Times:     times + 1,
	}
	db.Create(&paper_db)
	response.SuccessRe(c, "upload success", gin.H{"forback": fmt.Sprintf("%v上传%v问卷成功", acount, number)})
}
func findTimes(db *gorm.DB, acount, number string) int {
	var Uper model.PaperUser
	db.Where("Acount = ? and Number = ?", acount, number).Limit(1).Find(&Uper)
	if Uper.ID != 0 {
		return Uper.Times
	}
	return 0
}
