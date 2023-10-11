package models

import "time"

type User struct {
	Model
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
type Req struct {
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
}

func Login(data Req) (user User) {
	db.Table("user").Create(&User{
		Model: Model{
			CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
		},
		Nickname: data.Nickname,
		Password: data.Password,
	})
	return
}

func HasUser(key string, val interface{}) (user User) {
	if key == "id" {

	}
	db.Table("user").Where(key+" = ?", val).First(&user)
	return
}
