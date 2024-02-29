package models

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/global"
)

type Banner struct {
	Id int `json:"id"`
	AddBannerRequest
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type AddBannerRequest struct {
	Name       string `json:"name"`
	ImgUrl     string `json:"img_url"`
	Status     int    `json:"status"`
	Link       int    `json:"link"`
	JumpType   int    `json:"jump_type"`
	ActivityId string `json:"activity_id"`
	Type       int    `json:"type"`
}

func AddBanner(data AddBannerRequest) (banner Banner, err error) {
	banner.AddBannerRequest = data
	banner.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err = global.GAV_DB.Table("banner").Create(&banner).Error
	return banner, err
}

type EditBannerRequest struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ImgUrl     string `json:"img_url"`
	Status     int    `json:"status"`
	Link       int    `json:"link"`
	JumpType   int    `json:"jump_type"`
	ActivityId string `json:"activity_id"`
	Type       int    `json:"type"`
}

func EditBanner(data EditActivityRequest) (activity Activity, err error) {
	data.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	global.GAV_DB.Table("activity").Save(&data).Where("id = ?", data.Id).First(&activity)
	return activity, err
}

func GetBannerDetail(id int) (activity Activity, err error) {
	global.GAV_DB.Table("activity").Where("id = ?", id).First(&activity)
	return
}

type BannerListRequest struct {
	List   []Banner `json:"list"`
	Count  int      `json:"count"`
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
}

func GetBannerList() (activity []Activity, count int64, err error) {
	global.GAV_DB.Table("activity").Count(&count)
	global.GAV_DB.Table("activity").Offset(0).Limit(10).Find(&activity)
	return
}

type DelBannerResponse struct {
	Id int `json:"id"`
}

func DelBanner(id int) (err error) {
	delData := DelActivityResponse{
		Id: id,
	}
	global.GAV_DB.Table("activity").Delete(&delData)
	return
}
