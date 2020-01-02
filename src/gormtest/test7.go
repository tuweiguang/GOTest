package main

import (
	"database/sql"
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

	// delete from users where id=891;
	var user1 User
	user1.ID = 891
	//db.Delete(&user1)
	//
	//db.Where("id = ?",890).Delete(User{})

	//db.Where("id in (?)",[]int{888,889}).Delete(User{})

	//db.Where("id = ?", 2).Delete(&Base{})

	// 查询记录时会忽略被软删除的记录
	var base1 []Base
	db.Where("id = 2").Find(&base1)
	fmt.Println(base1)

	// 查询被软删除的记录
	var base2 []Base
	db.Unscoped().Where("id = 2").Find(&base2)
	fmt.Println(base2)

	// Unscoped 方法可以物理删除记录
	var base3 Base
	base3.ID = 2
	db.Unscoped().Delete(&base3)
}

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
