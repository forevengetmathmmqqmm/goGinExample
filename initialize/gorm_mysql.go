package initialize

import (
	"fmt"
	"log"

	"github.com/forevengetmathmmqqmm/goGinExample/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMySql() *gorm.DB {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	// tablePrefix = sec.Key("TABLE_PREFIX").String()
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	return db
}
