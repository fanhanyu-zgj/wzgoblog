package controllers

import (
	"fmt"
	"net/http"
	"wz/app/models/user"
	"wz/app/requests"
	"wz/pkg/auth"
	"wz/pkg/flash"
	"wz/pkg/view"
)

type AuthController struct {
}

type userForm struct {
	Name             string `gorm:"type:varchar(255);notnull;unique" valid:"name"`
	Email            string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password         string `gorm:"type:varchar(255)" valid:"password"`
	Password_confrim string `gorm:"-" valid:"password_confirm"`
}

func (auth *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}
	errs := requests.VlidateRegistrationForm(_user)
	if len(errs) > 0 {
		view.RenderSimple(w, view.D{"Errors": errs, "User": _user}, "auth.register")
	} else {
		_user.Create()
		if _user.ID > 0 {
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败,请联系管理员")
		}
	}
}
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	// 2. 尝试登录
	if err := auth.Attempt(email, password); err == nil {
		// 登录成功
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 3. 失败，显示错误提示
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}

}
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登录！")
	http.Redirect(w, r, "/", http.StatusFound)
}
