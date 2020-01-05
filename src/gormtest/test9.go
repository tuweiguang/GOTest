package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 从表
type CreditCard struct {
	gorm.Model
	Number string
	UID    string
}

// 主表
// 在主表中设置从表的外键和关联外键，与test8相反
type User struct {
	gorm.Model
	Name       string     `sql:"index"`
	CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
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

	db.DropTableIfExists(&CreditCard{})
	db.CreateTable(&CreditCard{})

	u := User{
		Name:       "user1",
		CreditCard: CreditCard{Number: "111111", UID: "aaaaa"},
	}
	db.Create(&u)

	var card CreditCard
	var user User
	user.ID = 1
	db.Model(&user).Related(&card, "CreditCard")
	fmt.Println(user)
	fmt.Println(card)
}
