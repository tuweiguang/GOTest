package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 主表
type User struct {
	gorm.Model
	MemberNumber string       // 关联外键
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string // 外键
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

	c := []CreditCard{
		{
			Number:           "card1",
			UserMemberNumber: "111",
		},
		{
			Number:           "card2",
			UserMemberNumber: "222",
		},
	}
	u := User{
		MemberNumber: "user1",
		CreditCards:  c,
	}
	db.Create(&u)

	//var card CreditCard
	//var user User
	//user.ID = 1
	//db.Model(&user).Related(&card, "CreditCard")
	//fmt.Println(user)
	//fmt.Println(card)
}
