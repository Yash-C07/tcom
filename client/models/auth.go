package models

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterModel struct {
	DisplayName     string `json:"displayname"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
