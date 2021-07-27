package event

type UserRegistrationRequestedMsg struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserRegistrationApprovedMsg struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserRegistrationCreatedMsg struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type EmailNotificationSentMsg struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
