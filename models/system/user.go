package system

import (
	"github.com/forevengetmathmmqqmm/goGinExample/global"
)

type User struct {
	global.Model
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
	Address       string `json:"address"`
	Intro         string `json:"intro"`
	Avatar        string `json:"avatar"`
	RoleId        int    `json:"role_id"`
	IsSuper       string `json:"is_super"`
}

type LoginApi struct {
	Nickname string `json:"nickname" form:"nickname"`
	Password string `json:"password" form:"password"`
}

type LoginResp struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type EditUserApi struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Intro      string `json:"intro"`
	Avatar     string `json:"avatar"`
	ModifiedOn string `json:"modified_on"`
	RoleId     int    `json:"role_id"`
	Password   string `json:"password"`
}

type AddUserApi struct {
	Nickname  string `json:"nickname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Intro     string `json:"intro"`
	Avatar    string `json:"avatar"`
	CreatedOn string `json:"modified_on"`
	RoleId    int    `json:"role_id"`
	Password  string `json:"password"`
}

type UserListResp struct {
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
	RoleId        int    `json:"role_id"`
	IsSuper       string `json:"is_super"`
}
