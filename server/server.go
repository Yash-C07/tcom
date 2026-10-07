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
	http.HandleFunc("/friend-req", handler.AddFriendReq)
	http.HandleFunc("/friend-list", handler.GetFriends)
	http.HandleFunc("/friend-remove", handler.RemoveFriend)
	http.HandleFunc("/friend-accept", handler.AcceptFriendReq)
	http.HandleFunc("/friend-reject", handler.RejectFriendReq)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Disconnect()
}
