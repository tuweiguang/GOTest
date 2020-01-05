package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model //id：联合外键
	Username   string
	Orders     Order
}
type Order struct {
	gorm.Model
	UserID uint // 外键
	Price  float64
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

	db.DropTableIfExists(&Order{})
	db.CreateTable(&Order{})

	db.Create(&User{
		Username: "tuweiguang1",
		Orders:   Order{Price: 10},
	})

	db.Create(&User{
		Username: "tuweiguang2",
		Orders:   Order{Price: 20},
	})

	db.Create(&User{
		Username: "tuweiguang3",
		Orders:   Order{Price: 30},
	})

	db.Create(&User{
		Username: "tuweiguang4",
		Orders:   Order{Price: 40},
	})

	// SELECT * FROM users;
	// 不会填充Orders从表数据
	// 默认是不会自动预加载 需要设置db.Set("gorm:auto_preload", true)
	var users1 []User
	db.Find(&users1)
	fmt.Println("user1:", users1)

	// SELECT * FROM users;
	// 但是会填充Orders从表数据
	var users2 []User
	db.Preload("Orders").Find(&users2)
	fmt.Println("user2:", users2)

	// SELECT * FROM users where orders.price not in (10,20);
	var user3 []User
	db.Preload("Orders", "price NOT IN (?)", []float64{10, 20}).Find(&user3)
	fmt.Println("user3:", user3)

	// SELECT * FROM users where username="tuweiguang3" and orders.price not in (10,20);
	var user4 []User
	db.Where("Username = ?", "tuweiguang2").Preload("Orders", "price NOT IN (?)", []float64{10, 20}).Find(&user4)
	fmt.Println("user4:", user4)
}
