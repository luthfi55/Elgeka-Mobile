package controllers

import (
	"fmt"
	"os"

	"elgeka-mobile/models"
)

func SendEmailWithGmail(emailTo string, OtpCode string) {
	sender := models.NewGmailSender(os.Getenv("GMAIL_SENDER"), os.Getenv("PASSWORD_SENDER"))

	subject := "Your One-Time Password (OTP) for Account Activation"
	content := fmt.Sprintf(`
	<html>
        <body>
            <p style="color: black;">Dear User,</p>
            <p style="color: black;">We're excited to have you on board! To complete your account activation, please use the following One-Time Password (OTP):</p>
            <p style="color: black;">Your One-Time Password:</p>
            <h1 style="color: black;">%s</h1>
            <p style="color: black;">Please note that this password is valid for only one minute for security reasons. Enter this code on the activation screen to proceed.</p>
            <p style="color: black;">If you didn't request this code or if you're having trouble, please contact our support team immediately.</p>
            <p style="color: black;">Thank you!</p>
        </body>
    </html>
	`, OtpCode)
	to := []string{emailTo}

	err := sender.SendEmail(subject, content, to, nil, nil)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully")
}
