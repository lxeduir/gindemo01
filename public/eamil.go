package public

import (
	"fmt"
	"gopkg.in/eapache/queue.v1"
	"net/smtp"
	"regexp"
	"time"
)

var tos = queue.New()
var adds = queue.New()
var flag = false

func Email(to []string, content string) bool {
	subject := fmt.Sprintf("Subject: %s\r\n", "国信安发信邮件")
	send := fmt.Sprintf("From: %s 发件邮箱\r\n", "国信安")
	receiver := fmt.Sprintf("To: %s\r\n", to[0])
	contentType := "Content-Type: text/html" + "; charset=UTF-8\r\n\r\n"
	msg := []byte(subject + send + receiver + contentType + content)
	addr := "smtp.qq.com:25"
	auth := smtp.PlainAuth("2508339002@qq.com", "2508339002@qq.com", "seuyirptahcudjjb", "smtp.qq.com")
	from := "2508339002@qq.com"
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		return false
	}
	return true
}
func CaptchaEmail(addr string) string {
	headText := "尊敬的用户："
	footText := "如果您没有获取过验证码，请忽略此邮件"
	starttime := "验证码过期时间为：" + time.Now().Add(time.Minute*5).Format("2006-01-02 15:04:05")
	content := "<div><p>" + headText + "</p><p>" + "您的验证码为: <h1>" + addr + "</h1></p><p>" + starttime + "</p></div>" + "<div><p>" + footText + "</p></div>"

	return content
}
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
} //正则表达式匹配邮箱地址
