package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello" + " go!")
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.Version())
}
