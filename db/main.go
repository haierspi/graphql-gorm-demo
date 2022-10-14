package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
)

func GetDB() *gorm.DB {
	// todo 未读进来，待看
	//dbDriver := os.Getenv("db_driver")
	//dbDsn := os.Getenv("db_dsn")

	//log.Infof("dbDriver: %v \n", dbDriver)

	db, err := gorm.Open("mysql", "root:root@tcp(xusheng1.xuxusheng.com:3306)/hhcg?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("DB open failed: %v \n", err)
	}
	log.Infof("Open db success: %v", db)
	return db
}

func init() {
	db := GetDB()

	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(&goods{}, &tag{}, &goodsTag{})
}
