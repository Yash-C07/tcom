// All models are defined here!
package models

type FriendReqModel struct {
	Sender   string
	Reciever string
}

type AllFriends struct {
	Items map[string]string
}
type AllFriendReq struct {
	Items []string
}
