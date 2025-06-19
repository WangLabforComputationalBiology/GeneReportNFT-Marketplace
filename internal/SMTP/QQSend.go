package SMTP

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"time"
)

var AuthCode = "jxghzvatitckggfj"

func init() {

}

func GenerateVerifyCode() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

func SendEmailCode(code, name, userAddress, email string) error {
	data := struct {
		Code    string
		Address string
	}{
		Code:    code,
		Address: userAddress,
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	m := gomail.NewMessage()
	m.SetHeader("From", "1356088661@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "your Verification Code")

	var buf bytes.Buffer
	tmpl, err := template.ParseFiles("email_template.html")

	if err := tmpl.Execute(&buf, data); err != nil {
		log.Fatal("Execute template failed:", err)
	}
	if err != nil {
		log.Fatal("Parse template failed:", err)
	}
	m.SetBody("text/html", buf.String())
	// 发送邮件
	d := gomail.NewDialer("smtp.qq.com", 587, "1356088661@qq.com", "your_auth_code")
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
