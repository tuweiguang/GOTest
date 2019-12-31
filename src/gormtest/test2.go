package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Base struct {
	gorm.Model
	Name         string `gorm:"column:mingzi"`
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 自定义表名
func (b Base) TableName() string {
	return "base_meta"
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.198.133:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)

	//// 迁移数据
	//db.AutoMigrate(&Base{})

	number := "111111111"
	t, err := time.Parse("2006-01-02", "1994-11-17")
	if err != nil {
		fmt.Println("time is error")
		return
	}

	// 创建表
	if !db.HasTable(&Base{}) {
		db.CreateTable(&Base{})
	}

	// 插入数据
	db.Create(&Base{
		Name:         "崔畅",
		Age:          sql.NullInt64{Int64: 18, Valid: true},
		Birthday:     &t,
		Email:        "cuichang@foxmail.com",
		Role:         "man",
		MemberNumber: &number,
		Num:          1,
		Address:      "安徽省合肥市",
		IgnoreMe:     100,
	})

	var base []Base
	db.Find(&base)
	fmt.Printf("row:\n")
	for _, v := range base {
		fmt.Println(v)
	}
}
