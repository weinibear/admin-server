package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/weinibear/admin-server/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		fmt.Printf("fail to read file: %v", err)
	}


}