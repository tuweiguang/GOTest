package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	ID        uint      // column name is `id`
	Name      string    // column name is `name`
	Age       *int      `gorm:"default:18"`
	Birthday  time.Time // column name is `birthday`
	CreatedAt time.Time // column name is `created_at`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", rand.Intn(1000))
	return nil
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.198.133:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	if !db.HasTable(&User{}) {
		// 为create table语句添加扩展SQL选项
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
		db.CreateTable(&User{})
	}

	user := User{Name: "Jinzhu" + strconv.Itoa(rand.Intn(1000)), Birthday: time.Now()}
	db.Create(&user)

}
