package models

import "time"

type Activity struct {
	Id int `json:"id"`
	AddActivityRequest
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type AddActivityRequest struct {
	Name           string `json:"name"`
	ImgUrl         string `json:"img_url"`
	Status         int    `json:"status"`
	IsOkApply      int    `json:"is_ok_apply"`
	ActivityStart  string `json:"activity_start"`
	ActivityEnd    string `json:"activity_end"`
	Type           int    `json:"type"`
	Num            int    `json:"num"`
	InitialVoucher int    `json:"initial_voucher"`
}

func AddActivity(data AddActivityRequest) (activity Activity, err error) {
	activity.AddActivityRequest = data
	activity.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err = db.Table("activity").Create(&activity).Error
	return activity, err
}

type EditActivityRequest struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ImgUrl         string `json:"img_url"`
	Status         int    `json:"status"`
	IsOkApply      int    `json:"is_ok_apply"`
	ActivityStart  string `json:"activity_start"`
	ActivityEnd    string `json:"activity_end"`
	Type           int    `json:"type"`
	Num            int    `json:"num"`
	InitialVoucher int    `json:"initial_voucher"`
	UpdateAt       string `json:"update_at"`
}

func EditActivity(data EditActivityRequest) (activity Activity, err error) {
	data.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	db.Table("activity").Save(&data).Where("id = ?", data.Id).First(&activity)
	return activity, err
}

func GetActivityDetail(id int) (activity Activity, err error) {
	db.Table("activity").Where("id = ?", id).First(&activity)
	return
}

type ActivityListRequest struct {
	List   []Activity `json:"list"`
	Count  int        `json:"count"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
}

func GetActivityList() (activity []Activity, count int64, err error) {
	db.Table("activity").Count(&count)
	db.Table("activity").Offset(0).Limit(10).Find(&activity)
	return
}

type DelActivityResponse struct {
	Id int `json:"id"`
}

func DelActivity(id int) (err error) {
	delData := DelActivityResponse{
		Id: id,
	}
	db.Table("activity").Delete(&delData)
	return
}
