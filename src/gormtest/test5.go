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
		fmt.Println("failed to connect database，err = %v", err)
		return
	}
	defer db.Close()

	// SELECT * FROM "users"  WHERE (age > (SELECT AVG(age) FROM "users"  WHERE (name = 'Jinzhu81')));
	//var user1 []User
	//db.Where("age > ?", db.Table("users").Select("AVG(age)").Where("name = ?", "Jinzhu81").QueryExpr()).Find(&user1)
	//fmt.Println("user1:",user1)

	/* 选择字段 */
	// select name,age from users;
	var user2 []User
	db.Select("name, age").Find(&user2)
	fmt.Println("user2:")
	for _, v := range user2 {
		fmt.Println(v)
	}

	// select name,age from users;
	var user3 []User
	db.Select([]string{"name", "age"}).Find(&user3)
	fmt.Println("user3:")
	for _, v := range user3 {
		fmt.Println(v)
	}

	rows, err := db.Table("users").Select("COALESCE(age,?)", 18).Rows()
	if err != nil {
		return
	}
	ages := make([]int, 0)
	// 操作原生sql
	for rows.Next() {
		age := 0
		err = rows.Scan(&age)
		if err != nil {
			return
		}
		ages = append(ages, age)
	}
	fmt.Println(ages)

	/* 排序 */
	// select * from users order by age desc,name;
	var user4 []User
	//db.Order("age desc, name").Find(&user4)
	db.Order("age desc").Order("name").Find(&user4)
	fmt.Println("user4:")
	for _, v := range user4 {
		fmt.Println(v)
	}

	// 覆盖排序
	var user5, user6 []User
	db.Order("age desc").Find(&user5).Order("age", true).Find(&user6)
	fmt.Println("user5:")
	for _, v := range user5 {
		fmt.Println(v)
	}
	fmt.Println("user6:")
	for _, v := range user6 {
		fmt.Println(v)
	}

	/* 数量 */
	var user7 []User
	db.Limit(2).Find(&user7)
	fmt.Println("user7:")
	for _, v := range user7 {
		fmt.Println(v)
	}

	// -1 取消 Limit 条件
	var user8, user9 []User
	db.Limit(2).Find(&user8).Limit(-1).Find(&user9)
	fmt.Println("user8:")
	for _, v := range user8 {
		fmt.Println(v)
	}
	fmt.Println("user9:")
	for _, v := range user9 {
		fmt.Println(v)
	}

	/* 偏移 */
	//(无效果)
	var user10 []User
	db.Offset(1).Find(&user10)
	fmt.Println("user10:")
	for _, v := range user10 {
		fmt.Println(v)
	}

	/* 计数*/
	var user11 []User
	count := 0
	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu81").Find(&user11).Count(&count)
	fmt.Println("user11:", user11, count)

	// 在没有具体对象情况下，Model和Table指定表名
	// 在有具体对象情况下，根据对象类型，找到对应的表名
	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	fmt.Println(count)

	db.Table("base_meta").Count(&count)
	fmt.Println(count)
	db.Table("base_meta").Select("count(distinct(age))").Count(&count)
	fmt.Println(count)

	/* group&having */
	// select age,count(name) from users group by age;
	rows, err = db.Table("users").Select("age, count(name)").Group("age").Rows()
	for rows.Next() {
		age := 0
		n := 0
		err = rows.Scan(&age, &n)
		if err != nil {
			return
		}
		fmt.Println(age, n)
	}

	// select age,count(name) from users group by age having count(name) > 1;
	type Result struct {
		Age    int64
		Number int64
	}
	var r []Result
	db.Table("users").Select("age, count(name)").Group("age").Having("count(name) > ?", 1).Scan(&r)
	fmt.Println(r)
	// 原生sql
	type Result2 struct {
		Age  int64
		Name string
	}
	var r2 []Result2
	db.Raw("SELECT name, age FROM users WHERE name = ?", "tuweiguang").Scan(&r2)
	fmt.Println(r2)

	rows, _ = db.Table("users").Select("age, count(name)").Group("age").Having("count(name) > ?", 1).Rows()
	for rows.Next() {
		age := 0
		n := 0
		err = rows.Scan(&age, &n)
		if err != nil {
			return
		}
		fmt.Println(age, n)
	}

	/* join */
	// select users.name,base_meta.mingzi from users left join base_meta on base_meta.id = users.id;
	// 左连接 右表会出现NULL
	rows, _ = db.Table("users").Select("users.name, base_meta.mingzi").Joins("left join base_meta on base_meta.id = users.id").Rows()
	for rows.Next() {
		var name string
		var mingzi string
		err = rows.Scan(&name, &mingzi)
		//if err != nil {
		//	fmt.Println("err = ",err)
		//	return
		//}
		fmt.Println(name, mingzi)
	}

	// select * from users
	// join base_meta ON base_meta.id = users.id AND base_meta.email = tuweiguang@foxmail.com
	// join credit_cards ON credit_cards.user_id = users.id
	// where credit_cards.number = 4111111111111111111
	//var user12 []User
	//db.Joins("JOIN base_meta ON base_meta.id = users.id AND base_meta.email = ?", "tuweiguang@foxmail.com").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user12)

	/* Pluck */
	// 查询表中某列数据
	var names []string
	db.Model(&User{}).Pluck("name", &names)
	fmt.Println("names:", names)
}
