package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Refer string
	Name  string
}

type Profile struct {
	gorm.Model
	Name       string
	User       User `gorm:"foreignkey:UserRefer2;association_foreignkey:Refer"` // 指定主表Refer字段作为关联外键；指定从表UserRefer2作为外键
	UserRefer  string
	UserRefer2 string
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.198.133:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 创建表
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	db.DropTableIfExists(&Profile{})
	db.CreateTable(&Profile{})

	//// 外键不受联合外键影响
	//u := User{
	//	Name:"tuweiguang1",
	//	Refer:"refer1",
	//}
	//db.Set("gorm:association_save_reference", false).Create(&Profile{
	//	Name:"tuweiguang1",
	//	User:u,
	//	UserRefer:"userdefer1",
	//	UserRefer2:"userdefer2",
	//})

	//// 跳过自动创建和更新
	//u2 := User{
	//	Name:"tuweiguang2",
	//	//Refer:"defer1",
	//}
	//db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Create(&Profile{
	//	Name:"tuweiguang2",
	//	User:u2,
	//	//UserRefer:"userdefer1",
	//	//UserRefer2:"userdefer2",
	//})
}
