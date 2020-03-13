package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Stu struct {
	Name  string `json:"name"`
	Age   int
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	grade int `json:"Grade"`
}

func main() {
	stu := Stu{
		Name: "张三",
		Age:  18,
	}

	cla := new(Class)
	cla.Name = "1班"
	cla.grade = 3
	stu.Class = cla

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStu))

	//除了struct类型，直接返回该类型
	a := []string{"hello", "world"}
	aa, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(aa))
}
