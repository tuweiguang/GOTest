package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "username", "姓名") //命令参数 默认值 提示
	//返回该类型的指针，此为*string
	var age int
	flag.IntVar(&age, "age", 20, "年龄") //绑定变量 命令参数 默认值 提示

	flag.Parse()

	args := flag.Args()

	fmt.Printf("name:%v\n", *name)
	fmt.Printf("age:%v\n", age)
	fmt.Printf("args:%v", args)
}

/*
执行：test.exe
结果：name:username
      age:20
	  args:[]

执行：test.exe -h
结果：Usage of test.exe:
        -age int
              年龄 <default 20>
        -name string
              姓名 <default "username">

执行：test.exe -name tuweiguang -age 22
结果：name:tuweiguang
      age:22
	  args:[]

执行：test.exe -name tuweiguang -age 22 aa bb cc
结果：name:tuweiguang
      age:20
	  args:[aa bb cc]

执行：test.exe -name tuweiguang -age -length 123
结果：flag provided but not defined:-length
      Usage of test.exe:
        -age int
              年龄 <default 20>
        -name string
              姓名 <default "username">
*/
