package main

import "fmt"

func main()  {
	var p1 *int
	i := 10
	p1 = &i

	fmt.Println(p1)	//p1中保存的是一段内存地址：0xc000016068
	fmt.Println(*p1)	//通过*来解引用，找到这段内存地址所指向的值

	*p1 = 1000
	fmt.Println(i)

	var p2 *string
	str1 := "test string"
	p2 = &str1

	fmt.Println(p2)
	fmt.Println(*p2)
	*p2 = "new string"
	fmt.Println(str1)

	var p3 *float64

	fmt.Println(p3)

	var p4 *bool
	b := true
	p4 = &b

	fmt.Println(p4)
	fmt.Println(*p4)
	fmt.Println(*(&b))
}