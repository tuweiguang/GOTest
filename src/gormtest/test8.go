package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 外键是在 从表 上面定义的，受 主表 关联键影响
// 1
//主表
//type User struct {
//	gorm.Model
//	Refer string
//	Name string
//}
//从表
//type Profile struct {
//	gorm.Model
//	Name      string
//	User      User `gorm:"association_foreignkey:Refer"` // 指定主表Refer字段作为关联外键
//	UserRefer string                                     // 必须有 “从表名+从表关联字段明”组合的字段，不然关联不起来
//}

// 2
type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"foreignkey:UserRefer"` // 指定从表UserRefer作为外键，默认与User表ID字段关联
	UserRefer uint
}

// 3
//type User struct {
//	gorm.Model
//	Refer string
//	Name string
//}
//
//type Profile struct {
//	gorm.Model
//	Name      string
//	User      User `gorm:"foreignkey:UserRefer2;association_foreignkey:Refer"` // 指定主表Refer字段作为关联外键；指定从表UserRefer2作为外键
//	UserRefer string
//    UserRefer2 string
//}
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

	u := User{
		Name: "tuweiguang1",
		//Refer:"defer1",
	}

	db.Create(&Profile{
		Name: "tuweiguang1",
		User: u,
		//UserRefer:"userdefer1",
		//UserRefer2:"userdefer2",
	})

	var user User
	var profile Profile
	profile.UserRefer = 1
	// SELECT * FROM profiles WHERE user_refer = 1;
	db.Model(&user).Related(&profile, "UserRefer").Find(&profile)
	fmt.Println(profile)
	fmt.Println(user)
}
