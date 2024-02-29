package models

import "github.com/forevengetmathmmqqmm/goGinExample/global"

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	global.GAV_DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int64) {
	global.GAV_DB.Model(&Tag{}).Where(maps).Count(&count)

	return
}
func ExistTagByName(name string) bool {
	var tag Tag
	global.GAV_DB.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool {
	global.GAV_DB.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}
func ExistTagByID(id int) bool {
	var tag Tag
	global.GAV_DB.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	global.GAV_DB.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {
	global.GAV_DB.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
