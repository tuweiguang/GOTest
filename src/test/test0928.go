package main

import (
	"encoding/json"
	"fmt"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	b2 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var m2 FamilyMember
	err := json.Unmarshal(b2, &m2)
	fmt.Println(err, m2)
}
