package models

import (
	"gin-blog/app/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func (this *User) CheckAuth(username, password string) (*User, error) {
	db.Where("username = ?", username).First(&this)
	if this.ID == 0 {
		return nil, &exceptions.Exception{Message: "用户不存在"}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(password)); err != nil {
		return nil, &exceptions.Exception{Message: "密码错误"}
	}
	return this, nil
}