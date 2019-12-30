package main

import (
	"flag"
	"fmt"
)

//自定义flag
type Nickname string

func (n *Nickname) Set(name string) error {
	*n = Nickname(name)
	return nil
}

func (n *Nickname) String() string {
	return fmt.Sprintf("my Nickname:%v\n", *n)
}

func main() {
	name := flag.String("name", "username", "姓名") //命令参数 默认值 提示
	//返回该类型的指针，此为*string
	var age int
	flag.IntVar(&age, "age", 20, "年龄") //绑定变量 命令参数 默认值 提示

	var nname Nickname
	flag.Var(&nname, "nickname", "nickname参数") //绑定变量 命令参数 提示

	flag.Parse()

	args := flag.Args()

	fmt.Printf("name:%v\n", *name)
	fmt.Printf("age:%v\n", age)
	fmt.Printf("args:%v\n", args)
	fmt.Printf("nickname:%v\n", nname)
}

/*
执行：test.exe
结果：name:username
      age:20
	  args:[]
	  nickname:

执行：test.exe -h
结果：Usage of test.exe:
        -age int
              年龄 <default 20>
        -name string
              姓名 <default "username">
        -nickname value
              nickname参数

执行：test.exe -name tuweiguang -age 22 -nickname soul
结果：name:tuweiguang
      age:22
	  args:[]
	  nickname:soul

执行：test.exe -name tuweiguang -age 22 aa bb cc -nickname soul
结果：name:tuweiguang
      age:20
	  args:[aa bb cc -nickname soul]
	  nickname:

执行：test.exe -name tuweiguang -age -length 123
结果：name:tuweiguang
      age:20
	  args:[aa bb cc]
	  nickname:soul
*/
