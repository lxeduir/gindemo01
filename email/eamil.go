package email

import (
	"fmt"
	"net/smtp"
)

func Email(to []string, content string) bool {
	subject := fmt.Sprintf("Subject: %s\r\n", "注册邮件")
	send := fmt.Sprintf("From: %s 测试发件邮箱\r\n", "国信安")
	receiver := fmt.Sprintf("To: %s\r\n", to[0])
	contentType := "Content-Type: text/plain" + "; charset=UTF-8\r\n\r\n"
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
func SignUp(addr string) string {
	headText := "请点击链接完成注册"
	//curl1 := "https://api.edulx.xyz/front/signup/emailverification"
	curl1 := "http://127.0.0.1:8000/front/signup/emailverification" //本地测试用
	curl1 += addr
	content := headText + curl1
	return content
}
