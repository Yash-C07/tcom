package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom/models"
	"tcom/utils"
)

func Login(request models.LoginModel) int {

	var user models.User

	readerData, er := json.Marshal(request)
	reader := bytes.NewReader(readerData)

	if er != nil {
		fmt.Println(er)
	}
	response, err := http.Post("http://localhost:3000/login", "application/json", reader)
	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	statuscode := response.StatusCode
	utils.SetCurrentUser(user)
	return statuscode
}

func Register(request models.RegisterModel) int {

	readerData, er := json.Marshal(request)
	reader := bytes.NewReader(readerData)

	if er != nil {
		fmt.Println(er)
	}

	response, err := http.Post("http://localhost:3000/register", "application/json", reader)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	statuscode := response.StatusCode
	return statuscode

}
