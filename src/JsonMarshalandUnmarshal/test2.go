package main

import (
	"encoding/json"
	"fmt"
)

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//方式1：只声明，不分配内存
	var stus1 []*StuRead

	//方式2：分配初始值为0的内存
	//make返回的引用类型
	// stus2 --> []*StuRead
	stus2 := make([]*StuRead, 0)

	//方式3：分配初始值为默认值的内存
	//new返回的指针类型
	//stus --> *[]*StuRead 这个结构
	stus := new([]*StuRead)

	stu1 := StuRead{Name: "张三", Age: 18, HIgh: true, sex: "男", Class: Class{Name: "1班", Grade: 1}, Test: 0}
	stu2 := StuRead{Name: "李四", Age: 19, HIgh: false, sex: "女", Class: Class{Name: "2班", Grade: 10}, Test: 0}

	//由方式1和2创建的切片，都能成功追加数据
	//方式2最好分配0长度，append时会自动增长。反之指定初始长度，长度不够时不会自动增长，导致数据丢失
	stus1 = append(stus1, &stu1)
	stus1 = append(stus1, &stu2)

	stus2 = append(stus2, &stu1)
	stus2 = append(stus2, &stu2)

	*stus = append(*stus, &stu1)
	//成功编码
	json1, _ := json.Marshal(stus1)
	json2, _ := json.Marshal(stus2)
	json3, _ := json.Marshal(*stus)

	fmt.Println(json1)
	fmt.Println(json2)
	fmt.Println(json3)

	//解码
	stud1 := make([]*StuRead, 0)
	err := json.Unmarshal(json1, &stud1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range stud1 {
		fmt.Println(v)
	}
}
