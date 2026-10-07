package utils

import "tcom/models"

var currentUser models.User

func GetCurrentUser() models.User {
	return currentUser
}
func SetCurrentUser(user models.User) {
	currentUser = user
}
