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

func Register_User(c *gin.Context) {
	//获取参数
	db_User := common.GetDB_User()
	Email := c.PostForm("email")
	Name := c.PostForm("name")
	PassWord := c.PostForm("password")
	Phone := c.PostForm("phone")
	Code_Email := c.PostForm("code")
	//数据验证
	if len(Phone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 403, nil, "手机号必须是11位数")
		return
	}
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
		log.Println("phone:", Phone, "name:", Name, "email:", Email)
		//判断手机号是否存在
		if isPhoneExited_User(db_User, Phone).ID != 0 {
			response.Response(c, http.StatusOK, 403, nil, "该手机号已注册")
			return
		}
		if isPhoneExited_User(db_User, Email).ID != 0 {
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
			newUser := model.User{
				Name:       Name,
				PassWord:   string(HashPassword),
				Phone:      Phone,
				ErrorTimes: 0,
				Email:      Email,
			}
			db_User.Create(&newUser)
			response.SuccessRe(c, "注册成功", gin.H{"code": 200, "msg": "注册成功"})
		} else {
			response.Response(c, http.StatusOK, 402, gin.H{"msg": "验证码填写错误"}, "err")
		}
	}
}
func isPhoneExited_User(db *gorm.DB, Phone string) model.User {
	var User model.User
	db.Where("Phone = ?", Phone).Order("id asc").Limit(1).Find(&User)
	return User
}
func Login_User(c *gin.Context) {
	db := common.GetDB_User()
	//获取参数
	Phone := c.PostForm("phone")
	PassWord := c.PostForm("password")
	//数据验证
	if len(Phone) != 11 {
		response.Response(c, http.StatusOK, 403, nil, "手机号必须是11位数")
		return
	}
	if len(PassWord) < 6 {
		response.Response(c, http.StatusOK, 400, nil, "密码不能低于6位数")
		return
	}
	log.Println("phone:", Phone, "PassWord:", PassWord, " is logining")
	User := isPhoneExited_User(db, Phone)
	log.Println(User)
	if User.ID == 0 {
		response.FalseRe(c, "用户不存在", nil)
		return
	} else if User.ErrorTimes >= 3 {
		response.FalseRe(c, "密码错误次数过多，账号已冻结", nil)
		return
	}
	//验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(User.PassWord), []byte(PassWord)); err != nil {
		response.FalseRe(c, "密码错误", nil)
		User.ErrorTimes++
		db.Save(&User)
		return
	}
	token, err := common.ReleaseToken_User(User)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "token生成失败，系统异常")
		log.Printf("token generate error : %v", err)
		return
	}
	DB_code := common.GetDB_Email()
	DB_code.Where("email=?", User.Email).Delete(&model.EmailCode{})
	response.SuccessRe(c, "登陆成功", gin.H{"token": token})
	User.ErrorTimes = 0
	db.Save(&User)
}
func Info_User(ctx *gin.Context) {
	User, _ := ctx.Get("User")
	response.Response(ctx, http.StatusOK, 200, gin.H{"User": dto.UserInfo(User.(model.User))}, "成功获取信息")
}
