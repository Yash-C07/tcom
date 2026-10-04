package main

import (
	"fmt"
	"net/http"
	"tcom-server/db"
	"tcom-server/handler"
)

func main() {

	db.Connect()
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Disconnect()
}
