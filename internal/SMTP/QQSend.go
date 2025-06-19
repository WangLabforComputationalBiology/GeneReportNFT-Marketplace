package SMTP

import (
	"bytes"
	"fmt"
	"gopkg.in/gomail.v2"
	"html/template"
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
		Address string
		Code    string
	}{
		Address: userAddress,
		Code:    code,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "1356088661@qq.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "your Verification Code")

	var buf bytes.Buffer
	tmpl, err := template.ParseFiles("/home/blockchain/GeneReportNFT-Marketplace/internal/SMTP/email_template.html")
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return err
	}

	m.SetBody("text/html", buf.String())
	// 发送邮件
	d := gomail.NewDialer("smtp.qq.com", 587, "1356088661@qq.com", AuthCode)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
