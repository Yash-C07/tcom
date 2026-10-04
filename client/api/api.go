package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom/types"
)

func Login(request types.LoginModel) int {

	readerData, er := json.Marshal(request)
	reader := bytes.NewReader(readerData)

	if er != nil {
		fmt.Println(er)
	}
	response, err := http.Post("http://localhost:3000/login", "application/json", reader)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	statuscode := response.StatusCode
	return statuscode
}

func Register(request types.RegisterModel) int {

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
