package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom-server/db"
	"tcom-server/types"
)

func RegisterHandler(writer http.ResponseWriter, reqptr *http.Request) {
	if reqptr.Method != http.MethodPost {
		return
	}
	var registerReq types.RegisterReq
	err := json.NewDecoder(reqptr.Body).Decode(&registerReq)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	Conn := db.GetConn()
	_, err = Conn.Exec(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)", registerReq.Username, registerReq.Password)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Registered successfully! Welcome ", registerReq.Username)
	writer.WriteHeader(http.StatusOK)

}
