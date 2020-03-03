package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.198.132:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	// 其作用主要是刷新数据库中的字段，使其保持最新，
	// 即让数据库之前存储的记录的表格字段和程序中最新使用的表格字段保持一致（只增不减），增加的使用默认值添加
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	fmt.Println(product)
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Println(product)

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}
