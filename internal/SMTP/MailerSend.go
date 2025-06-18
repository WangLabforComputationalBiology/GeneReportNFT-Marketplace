package SMTP

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

var APIKey = "mlsn.000038e6bccd11ad4741f9dd0357a441c63e3903a8282c3a2887dd5e17cc0db8mlsn.000038e6bccd11ad4741f9dd0357a441c63e3903a8282c3a2887dd5e17cc0db8"
var SMTPServ *mailersend.Mailersend

func init() {
	SMTPServ = mailersend.NewMailersend(APIKey)
}

func GenerateVerifyCode() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

func SendEmailCode(code, name, userAddress, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	subject := "your Verification Code"

	from := mailersend.From{
		Name:  "biochainer",
		Email: "biochainer@demo.com",
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
	_, _ = SMTPServ.Email.Send(ctx, message)

}
