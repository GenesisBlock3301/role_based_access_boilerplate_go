package user_services

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/serializers"
	"net/smtp"
	"os"
)

func SendEmail(user serializers.User, toEmail, subject, token string) error {
	relativeLink := "/email-verify/"
	absURL := "http://" + configurations.CurrentSite + configurations.BaseUrl + relativeLink + "?token=" + token
	emailBody := "Hi,  " + user.Email + " Use the link below to verify your email\n\n" + absURL
	//err = utils.SendEmail(user.Email, "Verify Your Email", emailBody)
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	msg := "From: " + from + "\n" +
		"To: " + toEmail + "\n" +
		"Subject: " + subject + "\n\n" + emailBody
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		smtp.PlainAuth("", from, password, smtpHost),
		from, []string{toEmail}, []byte(msg))
	return err
}
