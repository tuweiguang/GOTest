package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	ID        uint      // column name is `id`
	Name      string    // column name is `name`
	Age       *int      `gorm:"default:18"`
	Birthday  time.Time // column name is `birthday`
	CreatedAt time.Time // column name is `created_at`
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.198.133:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 根据主键查询第一条记录
	// SELECT * FROM users ORDER BY id LIMIT 1;
	var user1 User
	db.First(&user1)
	fmt.Println("user1:", user1)

	// 随机获取一条记录
	// SELECT * FROM users LIMIT 1;
	var user2 User
	db.Take(&user2)
	fmt.Println("user2:", user2)

	// 根据主键查询最后一条记录
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	var user3 User
	db.Last(&user3)
	fmt.Println("user3:", user3)

	// 查询所有的记录
	// SELECT * FROM users;
	var user4 []User
	db.Find(&user4)
	fmt.Println("user4:", user4)

	// 查询指定的某条记录(仅当主键为整型时可用)
	// SELECT * FROM users WHERE id = 10;
	var user5 User
	db.First(&user5, 87)
	fmt.Println("user5:", user5)

	// Get first matched record
	// select * from users where name = Jinzhu limit 1;
	var user6 User
	db.Where("name = ?", "Jinzhu").First(&user6)
	fmt.Println("user6:", user6)

	// Get all matched records
	// select * from users where age = 18;
	var user7 []User
	db.Where("age = ?", 18).Find(&user7)
	fmt.Println("user7:", user7)

	// <>
	var user8 []User
	db.Where("name <> ?", "Jinzhu").Find(&user8)
	fmt.Println("user8:", user8)

	// struct
	// SELECT * FROM users WHERE name = "jinzhu";
	var user9 []User
	db.Where(&User{Name: "Jinzhu"}).Find(&user9)
	fmt.Println("user9:", user9)

	// map
	// SELECT * FROM users WHERE name = "jinzhu";
	var user10 []User
	db.Where(map[string]interface{}{
		"name": "Jinzhu",
	}).Find(&user10)
	fmt.Println("user10:", user10)

	// slice
	// SELECT * FROM users WHERE id IN (1, 87, 887);
	var user11 []User
	db.Where([]int64{1, 87, 887}).Find(&user11)
	fmt.Println("user11:", user11)

	// 类似where
	var user12 []User
	db.First(&user12, 87)
	fmt.Println("user12:", user12)

	//
	var user13 []User
	db.Find(&user13, "name = ?", "Jinzhu")
	fmt.Println("user13:", user13)

	// FirstOrInit (无作用)
	// 根据给定的条件查询第一条记录 (仅支持 struct 和 map 条件)
	var user14 []User
	age := 18
	db.FirstOrInit(&user14, User{Name: "non_existing"})
	fmt.Println("user14:", user14)

	// Attrs
	//
	var user15 []User
	age = 20
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: &age}).FirstOrInit(&user15)
	fmt.Println("user15:", user15)

	var user16 []User
	db.Where(User{Name: "non_existing"}).Assign(User{Age: &age}).First(&user16)
	fmt.Println("user16:", user16)

	// FirstOrCreate (无作用)
	var user17 []User
	db.FirstOrCreate(&user17, User{Name: "non_existing"})
	fmt.Println("user17:", user17)
}
