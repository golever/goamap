package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   string `gorm:"PRIMARY"`
	Name string
	Age  int8
}

func (User) TableName() string {
	return "user"
}

func Insert() {
	user := User{"2","zhangsan", 17}

	db, _ := connect()
	//if err == nil {
	//	fmt.Println("failed to connect database")
	//	return
	//}

	db.Create(&user)
	fmt.Println("数据插入完成")
}

func connect() (db *gorm.DB, err error) {
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
