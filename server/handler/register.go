package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom-server/db"
	"tcom-server/models"
)

func RegisterHandler(writer http.ResponseWriter, reqptr *http.Request) {
	if reqptr.Method != http.MethodPost {
		return
	}
	var registerReq models.RegisterReq
	err := json.NewDecoder(reqptr.Body).Decode(&registerReq)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	Conn := db.GetConn()
	_, err = Conn.Exec(context.Background(),
		"INSERT INTO users (username, password,displayname) VALUES ($1, $2, $3)", registerReq.Username, registerReq.Password, registerReq.DisplayName)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Registered successfully! Welcome ", registerReq.Username)
	writer.WriteHeader(http.StatusOK)

}
