package user

import (
	"wz/app/models"
	"wz/pkg/password"
	"wz/pkg/route"
)

type User struct {
	models.BaseModel

	Name            string `gorm:"column:name;type:varchar(255);notnull;unique" valid:"name"`
	Email           string `gorm:"column:email;type:varchar(255);default:null;unique;" valid:"email"`
	Password        string `gorm:"colmun:password;type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

func (user User) Link() string {
	return route.RouteName2URL("users.show", "id", user.GetStringID())
}
