package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom-server/db"
	"tcom-server/types"
)

func LoginHandler(writer http.ResponseWriter, reqptr *http.Request) {
	if reqptr.Method != http.MethodPost {
		return
	}
	var loginReq types.LoginReq
	var u types.LoginReq
	err := json.NewDecoder(reqptr.Body).Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(loginReq)
	Conn := db.GetConn()
	err = Conn.QueryRow(context.Background(),
		"select * from users where username = $1 and password = $2", loginReq.Username, loginReq.Password).
		Scan(&u.Username, &u.Password)
	if err != nil {
		fmt.Println("Login Failed. Try again", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.Username == "" && u.Password == "" {
		fmt.Println("Login Failed. Try again")
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		fmt.Println("Login Success " + u.Username)
		writer.WriteHeader(http.StatusOK)
	}

	fmt.Println(u)

}
