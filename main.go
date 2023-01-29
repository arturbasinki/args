package main

import (
	"fmt"
	"os"
)

type Argument struct {
	Index int
	Value string
}

var Arguments = []Argument{}

func main() {
	args := os.Args
	for i, v := range args {
		if i == 0 {
			continue
		}
		temp := Argument{Index: i, Value: v}
		Arguments = append(Arguments, temp)
	}
	fmt.Println(Arguments)
}
