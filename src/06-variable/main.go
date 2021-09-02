package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var a string
	_ = a
	var b float32
	_ = b
	fmt.Println(time.Now().Format(time.Kitchen))
	c := 4
	_ = c
	d := 4
	i, _ := strconv.Atoi("10")
	d += i
	fmt.Print(d)

	var (
		e int  = 1
		f bool = false
	)
	_, _ = e, f
}
