package main

import "fmt"

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
}
