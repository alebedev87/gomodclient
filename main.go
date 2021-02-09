package main

import (
	"fmt"
	hello "github.com/alebedev87/gomodhello"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello.Proverb())
	fmt.Println(hello.MyHello())
	fmt.Println(hello.TagCheck())
	//fmt.Println(hello.TagFormat())
}
