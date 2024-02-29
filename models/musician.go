package models

import (
	"time"

	"github.com/forevengetmathmmqqmm/goGinExample/global"
)

type MusicianRequest struct {
	Id             int    `json:"id" from:"id"`
	StageName      string `json:"stage_name" form:"stage_name"`
	Brand          string `json:"brand" form:"brand"`
	SigningCompany string `json:"signing_company" form:"signing_company"`
	WeChat         string `json:"we_chat" form:"we_chat"`
	Weibo          string `json:"weibo" form:"weibo"`
	NetEaseCloud   string `json:"net_ease_cloud" form:"net_ease_cloud"`
	QqMusic        string `json:"qq_music" form:"qq_music"`
	Tiktok         string `json:"tiktok" form:"tiktok"`
	Status         int    `json:"status"`
	AuditTime      string `json:"audit_time"`
}
type Musician struct {
	MusicianRequest
	CreateAt string `json:"create_at"`
}
type MusicianListRequest struct {
	List   []Musician `json:"list"`
	Count  int        `json:"count"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
}

func AddMusician(data MusicianRequest) (musician Musician, err error) {
	musician.MusicianRequest = data
	musician.Status = 1
	musician.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	err = global.GAV_DB.Table("musician").Create(&musician).Error
	return musician, err
}

func EditMusician(data MusicianRequest) (musician Musician, err error) {
	musician.MusicianRequest = data
	if data.Status != 0 {
		data.AuditTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		data.Status = 1
	}
	global.GAV_DB.Table("musician").Save(&data).Where("id = ?", data.Id).First(&musician)
	return musician, err
}

func GetMusicianList() (musician []Musician, count int64, err error) {
	global.GAV_DB.Table("musician").Count(&count)
	global.GAV_DB.Table("musician").Offset(0).Limit(10).Find(&musician)
	return
}

func GetMusicianDetail(id int) (musician Musician, err error) {
	global.GAV_DB.Table("musician").Where("id = ?", id).First(&musician)
	return
}

type DelMusicianResponse struct {
	Id int `json:"id"`
}

func DelMusician(id int) (err error) {
	delData := DelMusicianResponse{
		Id: id,
	}
	global.GAV_DB.Table("musician").Delete(&delData)
	return
}
