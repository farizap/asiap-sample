package email

import (
	"log"
)

type EmailMockProvider struct {
}

func NewEmailMockProvider() *EmailMockProvider {
	return &EmailMockProvider{}
}

func (m *EmailMockProvider) SendEmail(email string) error {
	log.Printf("Email sent to", email)

	return nil
}
