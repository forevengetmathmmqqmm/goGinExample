package models

type Role struct { // 结构体首字母大写, 和数据库表名对应, 默认访问数据表users, 可以设置访问数据表的方法
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}
