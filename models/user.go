package models

import (
	"time"
)

type User struct {
	Model
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
	Address       string `json:"address"`
	Intro         string `json:"intro"`
	Avatar        string `json:"avatar"`
}
type UserInfo struct {
	User
	Keystore string `json:"keystore"`
}
type Web3 struct {
	Model
	UserId int `json:"userId"`
	WalletResp
}
type ResUser struct {
	CreatedOn string `json:"created_on"`
	RespUser
}
type RespUser struct {
	ID            int    `json:"id"`
	Nickname      string `json:"nickname"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
	Address       string `json:"address"`
	Intro         string `json:"intro"`
	Avatar        string `json:"avatar"`
	ModifiedOn    string `json:"modified_on"`
	CreatedOn     string `json:"create_on"`
}
type Req struct {
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
}

type UserListParams struct {
	Phone    string
	ID       int
	Nickname string
	Email    string
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
	db.Table("user").Where(key+" = ?", val).Limit(10).First(&user)
	return
}

func GetUser(key string, val interface{}) (userInfo UserInfo) {
	db.Table("user").Select(
		"user.id,user.created_on,user.modified_on,user.nickname,user.password,user.phone,user.email,user.wallet_address,user.address,user.intro,user.avatar,web3.keystore").Joins(
		"left join web3 on web3.user_id = user.id").First(&userInfo)
	return
}

// 获取用户列表
func GetUserList() (userList []RespUser) {
	db.Table("user").Offset(0).Limit(10).Find(&userList)
	return
}

func UpdateUser(data RespUser) (user ResUser) {
	data.ModifiedOn = time.Now().Format("2006-01-02 15:04:05")
	db.Table("user").Save(&data)
	db.Table("user").First(&user, data.ID)
	return
}

type WalletResp struct {
	WalletAddress string `json:"wallet_address" form:"wallet_address"`
	Keystore      string `json:"keystore" form:"keystore"`
}
type UserWallet struct {
	ID            int    `gorm:"primary_key" json:"id"`
	WalletAddress string `json:"wallet_address"`
}

func UpdateWallet(data WalletResp, userId int) {
	var web3 Web3
	web3.WalletResp = data
	web3.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	web3.UserId = userId
	db.Table("web3").Create(&web3)
	user := UserWallet{
		WalletAddress: data.WalletAddress,
		ID:            userId,
	}
	db.Table("user").Save(&user)
}
