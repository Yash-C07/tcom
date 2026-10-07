// All http handlers are put under this package!
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"tcom-server/db"
	"tcom-server/models"
)

func AddFriendReq(writer http.ResponseWriter, req *http.Request) {

	Conn := db.GetConn()

	if req.Method != http.MethodPost {
		return
	}

	var model models.FriendReqModel

	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer req.Body.Close()

	_, err = Conn.Exec(context.Background(), "insert into friend_reqs(sender, reciever, status) values($1, $2, 'p');", model.Sender, model.Reciever)

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func RemoveFriend(writer http.ResponseWriter, req *http.Request) {

	Conn := db.GetConn()
	if req.Method != http.MethodPost {
		return
	}

	var model models.FriendReqModel

	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Body.Close()

	_, err = Conn.Query(context.Background(), "delete from friend_reqs where sender = $1 and reciever = $2;", model.Sender, model.Reciever)

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func GetFriends(writer http.ResponseWriter, req *http.Request) {

	Conn := db.GetConn()
	if req.Method != http.MethodGet {
		return
	}
	var modelType struct {
		Name string
	}

	err := json.NewDecoder(req.Body).Decode(&modelType)
	model := modelType.Name
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	rows, err := Conn.Query(context.Background(),
		`SELECT CASE 
		WHEN sender = $1 THEN reciever
        	ELSE sender
    	END AS friend_id
		FROM public.friend_reqs
		WHERE (sender = $1 OR reciever = $1) 
	  	AND status = 'a';`, model)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var friends models.AllFriends
	var friendsList []string
	for rows.Next() {
		var f string
		err := rows.Scan(&f)
		if err != nil {
			fmt.Println(err)
			return
		}

		friendsList = append(friendsList, f)
	}
	friends.Items = friendsList

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	by, err := json.Marshal(friends)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Write(by)
	writer.WriteHeader(http.StatusOK)
}

func GetFriendReqs(writer http.ResponseWriter, req *http.Request) {

	Conn := db.GetConn()
	if req.Method != http.MethodGet {
		return
	}
	var modelType struct {
		Name string
	}

	err := json.NewDecoder(req.Body).Decode(&modelType)
	model := modelType.Name
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	rows, err := Conn.Query(context.Background(),
		`SELECT sender from friend_reqs
		WHERE (reciever = $1) 
	  	AND status = 'p';`, model)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var friends models.AllFriends
	var requestlist []string
	for rows.Next() {
		var f string
		err := rows.Scan(&f)
		if err != nil {
			fmt.Println(err)
			return
		}

		requestlist = append(requestlist, f)
	}
	friends.Items = requestlist

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	by, err := json.Marshal(friends)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Write(by)
	writer.WriteHeader(http.StatusOK)
}

func AcceptFriendReq(writer http.ResponseWriter, req *http.Request) {
	Conn := db.GetConn()
	if req.Method != http.MethodPost {
		return
	}

	var model models.FriendReqModel

	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Body.Close()

	_, err = Conn.Query(context.Background(), "update friend_reqs set status = 'a' where sender = $1 and reciever = $2;", model.Sender, model.Reciever)

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func RejectFriendReq(writer http.ResponseWriter, req *http.Request) {
	Conn := db.GetConn()
	if req.Method != http.MethodPost {
		return
	}

	var model models.FriendReqModel

	err := json.NewDecoder(req.Body).Decode(&model)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer req.Body.Close()

	_, err = Conn.Exec(context.Background(), "delete from friend_reqs where sender = $1 and reciever = $2; ", model.Sender, model.Reciever)

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
