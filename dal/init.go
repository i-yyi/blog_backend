package dal

import (
	config "codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/conf"
	"codehub.devcloud.cn-north-4.huaweicloud.com/taqse2020_mihoyo00001/blog_backend/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	log.Info("Init Mysql %v", config.Conf.MySQL)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.GetConf().MySQL.DSN,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	db.AutoMigrate(&model.Comment{}, &model.Post{})
	if err != nil {
		log.Error("Gorm Init Failed")
	}
}
func GetDB() *gorm.DB {
	return db
}
