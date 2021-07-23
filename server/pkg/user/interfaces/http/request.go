package http

type AddUserRegistration struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	ManagerID string `json:"managerID"`
}
