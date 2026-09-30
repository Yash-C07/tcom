package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/tcom-db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close(context.Background())
	Conn = conn

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func loginHandler(writer http.ResponseWriter, reqptr *http.Request) {
	if reqptr.Method != http.MethodPost {
		return
	}
	var loginReq LoginReq
	var u LoginReq
	err := json.NewDecoder(reqptr.Body).Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loginReq)
	err = Conn.QueryRow(context.Background(),
		"select * from users where username = $1 and password = $2", loginReq.Username, loginReq.Password).
		Scan(&u.Username, &u.Password)
	if err != nil {
		fmt.Println("Login Failed. Try again", err)
		return
	}

	if u.Username == "" && u.Password == "" {
		fmt.Println("Login Failed. Try again")
		writer.WriteHeader(http.StatusForbidden)
	} else {
		fmt.Println("Login Success " + u.Username)
		writer.WriteHeader(http.StatusOK)
	}

	fmt.Println(u)

}

func registerHandler(writer http.ResponseWriter, reqptr *http.Request) {
	if reqptr.Method != http.MethodPost {
		return
	}
	var registerReq RegisterReq
	err := json.NewDecoder(reqptr.Body).Decode(&registerReq)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		return
	}

	_, err = Conn.Exec(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2)", registerReq.Username, registerReq.Password)
	if err != nil {
		fmt.Println("Registration Unsuccessful Error: ", err)
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println("Registered successfully! Welcome ", registerReq.Username)
	writer.WriteHeader(http.StatusOK)

}
