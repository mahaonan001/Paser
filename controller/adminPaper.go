package controller

import (
	"PaSer/common"
	"PaSer/model"
	"PaSer/response"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Paper struct {
	Name          string              `json:"name"`
	Title         []string            `json:"model"`
	Questionnaire []map[string]string `json:"questionnaire"`
}

func AddPaper(c *gin.Context) {
	var paper Paper
	c.ShouldBindJSON(&paper)
	db := common.Con_Db_asp()
	Name := paper.Name
	Title := strings.Join(paper.Title, ",")
	questionnaireJson, err := json.Marshal(paper.Questionnaire)
	if err != nil {
		response.FalseRe(c, fmt.Sprintf("Error encoding questionnaire: %s", err), nil)
		return
	}
	Questions := string(questionnaireJson)
	PaperNewUp := model.PaperNew{
		Name:      Name,
		Titles:    Title,
		Questions: Questions,
	}
	fmt.Println("paper:", paper)
	fmt.Println(PaperNewUp)
	db.Create(&PaperNewUp)
	response.SuccessRe(c, "问卷已上传", nil)
}
