package SMTP

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "mlsn.d82c5a91a0da6195a66e31183b48c3e4bab1f4838046f92893b5046bdffb94ca"
var SMTPServ *mailersend.Mailersend

func init() {
	SMTPServ = mailersend.NewMailersend(APIKey)
}

func GenerateVerifyCode() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

func SendEmailCode(code, name, userAddress, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	subject := "your Verification Code"

	from := mailersend.From{
		Name:  "biochainer",
		Email: "test-r9084zvw9m8gw63d.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  name,
			Email: email,
		},
	}

	personalization := []mailersend.Personalization{
		{
			Email: email,
			Data: map[string]interface{}{
				"code":    code,
				"address": userAddress,
			},
		},
	}

	tags := []string{"foo", "bar"}

	message := SMTPServ.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetTemplateID("neqvygm9z6zl0p7w")
	message.SetPersonalization(personalization)

	message.SetTags(tags)
	_, err := SMTPServ.Email.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
