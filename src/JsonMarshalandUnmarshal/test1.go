package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Foo string `json:"foo"`
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

type Bar struct {
	Aaa int
	Bbb string
}

type Foo struct {
	Bar *Bar
}

func main() {
	b2 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var m2 FamilyMember
	err := json.Unmarshal(b2, &m2)
	fmt.Println(err, m2)
	//
	//b3 := []byte(`{"Bar":{"Aaa":1,"Bbb":"Bbb"}}`)
	//var m3 Foo
	//err = json.Unmarshal(b3, &m3)
	//fmt.Println(err,m3.Bar)
}
