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
	subject := fmt.Sprintf("Subject: %s\r\n", "注册邮件")
	send := fmt.Sprintf("From: %s 测试发件邮箱\r\n", "国信安")
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
	fmt.Println("发信时间" + time.Now().String() + "收件人:" + to[0])
	return true
}
func SignUp(addr string) string {
	headText := "请点击链接完成注册"
	curl1 := "https://api.edulx.xyz/front/signup/emailverification"
	curl1 = "http://127.0.0.1:8000/front/signup/emailverification" //本地测试用
	curl1 += addr
	content := "<div><p>" + headText + "</p><a href=\"" + curl1 + "\">验证链接</a></div>"
	return content
}
func QueueEmail(to string, add string) {
	tos.Add(to)
	adds.Add(add)
	if !flag {
		go func() {
			for true {
				if tos.Length() > 0 {
					t := fmt.Sprintf("%v", tos.Remove())
					tos := []string{t}
					a := fmt.Sprintf("%v", adds.Remove())
					Email(tos, SignUp(a))
				}
				time.Sleep(time.Second)
			}
			flag = false
		}()
	} else {
		flag = true
	}

}
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
} //正则表达式匹配邮箱地址
