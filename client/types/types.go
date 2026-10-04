package types

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterModel struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
