package main

import (
	"fmt"
	// 为什么就是解析不了？？？
	"myGo/src/vsGo/pack"
)

func noHope() {
	const c = "k"
	hello := "hellp"

	_ = hello
}

var hello = "hello"

func nested() {

	var hello = "no"
	fmt.Println(hello)

}

func main() {
	greet()
	pack.pack()
}
