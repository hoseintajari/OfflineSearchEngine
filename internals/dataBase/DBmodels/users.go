package DBmodels

import (
	"OfflineSearchEngine/internals/apiServer/server/requestResponsemodels"
	"OfflineSearchEngine/internals/dataBase"
	"OfflineSearchEngine/internals/utility"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Password string
	Roles    bool `gorm:"default:false"`
	Email    string
}

var db = dataBase.Connection()

func CreateUserTable() User {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("Failed to Create Users table!")
	}
	if db.Where("user_name=?", "Admin").Find(&User{}).RowsAffected == 0 {
		db.Save(&User{UserName: "Admin", Password: utility.HashPassword("admin"), Roles: true})
	}
	return User{}
}

func (u *User) Add(input requestResponsemodels.SignUpRequest) {
	db.Save(&User{
		UserName: input.UserName,
		Password: utility.HashPassword(input.Password),
		Email:    input.Email,
	})

}

func (u *User) Remove(request requestResponsemodels.RemoveUserRequest) {
	db.Where("user_name=?", request.UserName).Find(&User{}).Delete(&User{})
}
