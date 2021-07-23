package http

type AddUserRegistrationResponse struct {
	ID string `json:"id"`
}

//NewCreateNewItemResponse construct CreateNewItemResponse
func NewAddUserRegistrationResponse(id string) *AddUserRegistrationResponse {
	return &AddUserRegistrationResponse{
		id,
	}
}
