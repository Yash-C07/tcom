package models

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	DisplayName     string `json:"displayname"`
}
