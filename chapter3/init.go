package main

import (
	"fmt"
)

var num int64

func init()  {
	num = 100
}

func main() {
	fmt.Printf("init func, num = %v\n", num)
}
