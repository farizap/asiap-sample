package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v8"
)

func main() {
	client := gocloak.NewClient("http://localhost:8080")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "farizapriyanto@gmail.com", "123456", "axiapp")
	if err != nil {
		fmt.Print(err)
		panic("Something wrong with the credentials or url")
	}

	attr := make(map[string][]string)
	attr["division"] = []string{"sp"}
	user := gocloak.User{
		FirstName:  gocloak.StringP("Bob"),
		LastName:   gocloak.StringP("Uncle"),
		Email:      gocloak.StringP("somethingFalse@really.wrong"),
		Enabled:    gocloak.BoolP(true),
		Username:   gocloak.StringP("somethingFalse@really.wrong"),
		Attributes: &attr,
	}

	// userPass := gocloak.SetPasswordRequest{
	// 	Temporary: gocloak.BoolP(false),
	// 	Password:  gocloak.StringP("password"),
	// }

	_, err = client.CreateUser(ctx, token.AccessToken, "axiapp", user)

	if err != nil {
		fmt.Print(err.Error())
		panic("Oh no!, failed to create user :(")
	}

	getUsersParams := gocloak.GetUsersParams{Email: user.Email}
	userInfo, err := client.GetUsers(ctx, token.AccessToken, "axiapp", getUsersParams)
	if err != nil {
		fmt.Print(err.Error())
		panic("Oh no!, failed to get user info :(")
	}

	err = client.SetPassword(ctx, token.AccessToken, *userInfo[0].ID, "axiapp", "123456", false)
	if err != nil {
		fmt.Print(err.Error())
		panic("Oh no!, failed to set password :(")
	}

	jwt, err := client.Login(ctx, "backoffice-dashboard", "0913ce29-e2fc-4244-93d8-6c08117d0eec", "axiapp", "somethingFalse@really.wrong", "123456")
	if err != nil {
		fmt.Print(err.Error())
		panic("Oh no!, failed to login :((")
	}

	fmt.Println(jwt.AccessToken)
	fmt.Println(jwt.ExpiresIn)
	fmt.Println(jwt.RefreshExpiresIn)
	fmt.Println(jwt.RefreshToken)
	fmt.Println(jwt.TokenType)

}
