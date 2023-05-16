package dao

import (
	"github.com/sirupsen/logrus"
	"github.com/zouxingyuks/tools/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var daoLog *logrus.Entry

func InitDao() {
	dsn := config.Configs.GetString("dao.username") + ":" + config.Configs.GetString("dao.password") + "@tcp(" + config.Configs.GetString("dao.host") + ")/" + config.Configs.GetString("dao.dbname") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		daoLog.Errorln("数据库连接失败")
		panic(err)
	}
	DB = db
}
