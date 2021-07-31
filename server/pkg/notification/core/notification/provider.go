package notification

type Provider interface {
	SendEmail(id string) error
}
