package validate

import (
	"fmt"
	"net/smtp"
	"strings"
)

func UserRegisterMiddleware() {
	auth := smtp.PlainAuth("", "1113181943@qq.com", "", "smtp.qq.com")
	to := []string{"1113181943@qq.com"}
	nickname := "test"
	user := "1113181943@qq.com"
	subject := "test mail"
	contentType := "Content-Type: text/plain; charset=UTF-8"
	body := "This is the email body."
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}
