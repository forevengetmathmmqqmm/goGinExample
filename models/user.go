package models

import (
	"fmt"
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
	db.Table("user").Where(key+" = ?", val).First(&user)
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

func UpdateWallet(data WalletResp, userId int) {
	var web3 Web3
	web3.WalletResp = data
	web3.CreatedOn = time.Now().Format("2006-01-02 15:04:05")
	web3.UserId = userId
	db.Table("web3").Create(&web3)
	user := map[string]interface{}{
		"WalletAddress": data.WalletAddress,
		"ID":            userId,
	}
	fmt.Print(">>>>user", user)
	db.Table("user").Save(&user)
}
