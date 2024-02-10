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
	response.Response(c, 200, 200, gin.H{"code": PaperNewUp.Number}, "上传成功")
}
