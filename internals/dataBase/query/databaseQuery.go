package query

import (
	"OfflineSearchEngine/internals/dataBase"
	"OfflineSearchEngine/internals/dataBase/DBmodels"
	"OfflineSearchEngine/internals/utility"
)

var db = dataBase.Connection()

func FindUsername(name string) bool {
	result := db.Where("user_name =? ", name).First(&DBmodels.User{})
	if result.RowsAffected == 1 {
		return true
	} else {
		return false
	}
}

func CheckPassword(username string, password string) bool {
	var user DBmodels.User
	result := db.Where("user_name=?", username).Find(&user)
	if result.Error != nil {
		return false
	}
	return user.Password == utility.HashPassword(password)
}

func GetRoles(username string) bool {
	var user DBmodels.User
	db.Where("user_name=?", username).First(&user)
	response := user.Roles
	return response
}
