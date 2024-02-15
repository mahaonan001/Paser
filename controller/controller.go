package controller

import (
	"PaSer/common"
	"PaSer/dto"
	"PaSer/model"
	"PaSer/response"
	"PaSer/util"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	//获取参数
	db_Admin := common.GetDB_Admin()
	Email := c.PostForm("email")
	Name := c.PostForm("name")
	PassWord := c.PostForm("password")
	Code_Email := c.PostForm("code")
	//数据验证
	if len(PassWord) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 403, nil, "密码不能低于6位数")
		return
	}
	if len(Name) == 0 {
		Name = util.RandomString(10)
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	// 使用MatchString()函数来判断电子邮件地址是否匹配正则表达式
	if emailRegex.MatchString(Email) {
		log.Println("name:", Name, "email:", Email)
		if isEmailExited(db_Admin, Email).ID != 0 {
			response.Response(c, http.StatusOK, 403, nil, "该邮箱已注册")
			return
		}
		DB_code := common.GetDB_Email()
		var Code_email model.EmailCode
		DB_code.Where("email=?", Email).Order("id desc").Limit(1).Find(&Code_email)
		if Code_Email == Code_email.Code_email { //存在符合条件的验证码
			HashPassword, err := bcrypt.GenerateFromPassword([]byte(PassWord), bcrypt.DefaultCost)
			if err != nil {
				response.Response(c, http.StatusInternalServerError, 500, nil, "加密失败")
				log.Println(err)
				return
			}
			newAdmin := model.Admin{
				Name:       Name,
				PassWord:   string(HashPassword),
				ErrorTimes: 0,
				Email:      Email,
			}
			db_Admin.Create(&newAdmin)
			response.SuccessRe(c, "注册成功", gin.H{"code": 200, "msg": "注册成功"})
		} else {
			response.Response(c, http.StatusOK, 402, gin.H{"msg": "验证码填写错误"}, "err")
		}
	}
}
func isEmailExited(db *gorm.DB, Email string) model.Admin {
	var Admin model.Admin
	db.Where("Email = ?", Email).Order("id asc").Limit(1).Find(&Admin)
	return Admin
}
func Login(c *gin.Context) {
	db := common.GetDB_Admin()
	//获取参数
	Email := c.PostForm("email")
	PassWord := c.PostForm("password")
	//数据验证
	if len(PassWord) < 6 {
		response.Response(c, http.StatusOK, 400, nil, "密码不能低于6位数")
		return
	}
	log.Println("email:", Email, "PassWord:", PassWord, " is logining")
	Admin := isEmailExited(db, Email)
	log.Println(Admin)
	if Admin.ID == 0 {
		response.FalseRe(c, "用户不存在", nil)
		return
	} else if Admin.ErrorTimes >= 3 {
		response.FalseRe(c, "密码错误次数过多，账号已冻结", nil)
		return
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(Admin.PassWord), []byte(PassWord)); err != nil {
		response.FalseRe(c, "密码错误", nil)
		Admin.ErrorTimes++
		db.Save(&Admin)
		return
	}
	token, err := common.ReleaseToken(Admin)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token生成失败，系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	DB_code := common.GetDB_Email()
	DB_code.Where("email=?", Admin.Email).Delete(&model.EmailCode{})
	response.SuccessRe(c, "登陆成功", gin.H{"token": token})
	Admin.ErrorTimes = 0
	db.Save(&Admin)
}
func Info(ctx *gin.Context) {
	Admin, _ := ctx.Get("Admin")
	response.Response(ctx, http.StatusOK, 200, gin.H{"Admin": dto.AdminInfo(Admin.(model.Admin))}, "成功获取信息")
}
