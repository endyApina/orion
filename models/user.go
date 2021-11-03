package models

import (
	"errors"
	"orion/mailer"
	"os"

	"github.com/joho/godotenv"
)

func (data Subscribe) RegisterNewSubscriber() error {
	if data.FullName == "" {
		return errors.New("empty full name")
	}

	if data.Email == "" {
		return errors.New("empty email field")
	}

	if data.PhoneNumber == "" {
		return errors.New("empty phone number field")
	}

	if err := conn.Create(&data).Error; err != nil {
		LogError(err)
		return err
	}
	return nil
}

func (data Subscribe) NotifyNewSubscriber() error {
	_ = godotenv.Load("conf.env")
	templatePath := os.Getenv("template_path") + "subscription.html"

	mailSubject := "ORION APP - Thank you for your subscription"
	newRequestData := mailer.NewRequest(data.Email, mailSubject)
	go newRequestData.SendOTPRegistrationEmail(templatePath, data)
	return nil
}
