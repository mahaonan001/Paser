package controller

import (
	"PaSer/common"
	"PaSer/model"
	"PaSer/response"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPaper(c *gin.Context) {
	var Paper model.PaperNew
	Code := c.Query("code")
	db := common.GetDB_ASP()
	result := db.Where("Number=?", Code).Limit(1).Find(&Paper)
	if result.Error != nil {
		response.FalseRe(c, "数据库连接失败", nil)
		return
	}
	if Paper.ID == 0 {
		response.FalseRe(c, "无此问卷", nil)
		return
	}
	jsonBytes, err := json.Marshal(Paper)
	if err != nil {
		fmt.Println("Error marshaling struct to JSON:", err)
		return
	}
	var jsonObj map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonObj)
	if err != nil {
		fmt.Println("Error unmarshaling JSON to object:", err)
		return
	}
	response.SuccessRe(c, "", gin.H{"paper": jsonObj["Questions"]})
}
