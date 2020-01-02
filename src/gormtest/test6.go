package main

import (
	"database/sql"
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

	age := 25

	/* save:如果指定主键，更新主键行数据；如果没有指定主键，将会新插入一行 */
	// insert into users(name,age,birthday) values("tuweiguang",25,nowtime);
	//var user1 User
	//user1.Name = "tuweiguang"
	//user1.Age = &age
	//user1.Birthday = time.Now()
	//db.Save(&user1)

	// update users set name="tuweiguang",age=25,birthday=nowtime where id=2;
	var user2 User
	user2.ID = 2
	user2.Name = "tuweiguang"
	user2.Age = &age
	user2.Birthday = time.Now()
	db.Save(&user2)

	/* 修改指定字段 */
	// 更新单个属性，如果它有变化
	// update users set name="cuichang" where id=87;
	var user3 User
	user3.ID = 87
	db.Model(&user3).Update("name", "cuichang")

	// update users set name="cuichang" where id=87 and age=18;
	db.Model(&user3).Where("age = ?", 18).Update("name", "tuweiguang2")

	// 更新多个属性
	var user4 User
	// update
	user4.ID = 887

	// update users set name="tuweiguang3",age=26 where id=887;
	db.Model(&user4).Updates(map[string]interface{}{"name": "tuweiguang3", "age": 26})

	user4.ID = 888
	age = 27
	// update users set name="tuweiguang4",age=27 where id=888;
	db.Model(&user4).Updates(User{Name: "tuweiguang4", Age: &age})

	user4.ID = 889
	// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	db.Model(&user4).Updates(User{Name: ""})

	// 更新或忽略某些字段
	var user5 User
	user5.ID = 889
	// update users set name=tuweiguang5 where id=889;
	db.Model(&user5).Select("name").Updates(map[string]interface{}{"name": "tuweiguang5", "age": 28})

	user5.ID = 890
	// update users set age=29 where id=890;
	db.Model(&user5).Omit("name").Updates(map[string]interface{}{"name": "tuweiguang6", "age": 29})

	// !!!! 以上所有更新都会触发钩子函数BeforeUpdate, AfterUpdate，会更新updated_at字段，若model没有updated_at字段就算了，
	// 若有的话，updated_at被更新为当前时间。若不想在update时候，不触发钩子函数，使用UpdateColumn
	var base1 Base
	base1.ID = 1
	db.Model(&base1).Update("age", 21)

	db.Model(&base1).UpdateColumn("age", 23)

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
