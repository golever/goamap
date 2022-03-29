package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysql_connect() (db *gorm.DB, err error) {
	d := "root:123456@tcp(127.0.0.1:3306)/goamap?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       d,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

}
