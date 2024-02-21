package mail

import (
	"PaSer/common"
	"PaSer/model"
	"PaSer/response"
	"PaSer/util"
	"errors"
	"log"
	"net/http"
	"net/smtp"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 主函数
func Code_email(c *gin.Context) {
	//获取前端传来的请求中的email
	User_email := c.PostForm("email")
	//获取验证码和验证码截至时间
	SendMail(User_email, c)
}
func SendMail(Toemail string, c *gin.Context) {
	DB := common.GetDB_Email()
	smtpHost := "smtp.qq.com"                                 // SMTP服务器地址
	smtpPort := "587"                                         // SMTP服务器端口
	smtpUser := viper.GetString("emailCode.email")            // SMTP用户名
	smtpPassword := viper.GetString("emailCode.smtpPassword") // SMTP密码（授权码）
	toUserEmail := Toemail                                    // 接收者邮箱地址
	code := util.RandomString(6)                              // 验证码
	infTime := time.Now().Add(1 * time.Minute)
	//验证码可用
	if !IsEmailLegal(toUserEmail, DB) {
		response.Response(c, http.StatusBadRequest, 402, gin.H{"msg": "err"}, "请1分钟后重试")
		return
	}
	//验证码不可用
	if CodeTimeAble(DB, toUserEmail) {
		response.Response(c, 200, 402, gin.H{"msg": "请稍后获取验证码"}, "验证码已申请")
		return
	}
	e := email.NewEmail()
	e.From = smtpUser                                                                                                                                     // 发件人邮箱账号
	e.To = append(e.To, toUserEmail)                                                                                                                      // 收件人邮箱地址                                                                               // 收件人邮箱地址
	e.Subject = "欢迎使用AsPer问卷系统"                                                                                                                           // 邮件主题
	e.Text = []byte("验证码:" + code)                                                                                                                        // 邮件正文内容（纯文本）
	e.HTML = []byte("<strong>" + string(e.Text) + "</strong><br><p>有效时长5分钟</p><p>  本项目由mahaonan001在GitHub上开源的问卷系统go项目,如果有兴趣参加,欢迎联系1649801526@qq.com</p>") // 邮件正文内容（HTML格式）
	newCode := model.EmailCode{
		Email:      toUserEmail,
		Code_email: code,
		InfTime:    infTime,
	}
	err := e.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)) // 发送邮件
	if err == nil {
		DB.Create(&newCode)
		response.SuccessRe(c, "成功获取验证码", nil)
		return
	}

	// 处理发送邮件失败的情况
	// 比如打印日志或返回错误信息
	log.Println("邮件发送失败", err)
}

// 判断邮箱是否合法或在一分钟内是否申请过验证码
func IsEmailLegal(email string, DB *gorm.DB) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	// 使用MatchString()函数来判断电子邮件地址是否匹配正则表达式
	return emailRegex.MatchString(email)
}

func CodeTimeAble(DB *gorm.DB, email string) bool { //存在一分钟内可用的验证码
	var EmailCode model.EmailCode
	TimeNow := time.Now()
	if err := DB.Model(&model.EmailCode{}).Error; err != nil {
		log.Println("fail to get table count")
		return false
	}

	result := DB.Where("email=? and inf_time >= ?", email, TimeNow).Limit(1).Find(&EmailCode)
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	if EmailCode.ID != 0 {
		return true //已经获取获取验证码
	} else {
		// 查询到结果
		return false
	}
}
