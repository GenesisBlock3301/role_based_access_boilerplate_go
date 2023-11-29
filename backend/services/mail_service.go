package services

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/serializers"
	"net/smtp"
	"os"
)

func SendEmail(toEmail, subject, token string, data serializers.OTPSerializer) error {
	var relativeLink, absURL, emailBody string
	if data.IsOTP == true {
		emailBody = "Hi, Your OTP code is: " + data.Code
	} else {
		relativeLink = "/email-verify/"
		absURL = "http://" + configurations.CurrentSite + configurations.BaseUrl + relativeLink + "?token=" + token
		emailBody = "Hi,  " + toEmail + " Use the link below to verify your email\n\n" + absURL
	}

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
