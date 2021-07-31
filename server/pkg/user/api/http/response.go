package http

import (
	"asiap/pkg/user/core/user"
)

type AddUserRegistrationResponse struct {
	ID string `json:"id"`
}

//NewCreateNewItemResponse construct CreateNewItemResponse
func NewAddUserRegistrationResponse(id string) *AddUserRegistrationResponse {
	return &AddUserRegistrationResponse{
		id,
	}
}

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

type UsersResponse struct {
	Users []*UserResponse `json:"users"`
}

func NewUsersResponse(users []user.User) *UsersResponse {
	var usersResponse []*UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, &UserResponse{user.ID(), user.Name(), user.Email(), user.Location(), user.Status()})
	}

	return &UsersResponse{
		usersResponse,
	}
}
