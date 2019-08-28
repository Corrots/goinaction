package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main()  {
	var b bytes.Buffer

	// 将字符串写入buffer
	b.Write([]byte("Hello "))
	//使用Fprintf将字符串拼接到buffer
	fmt.Fprintf(&b, "Golang!")
	//将Buffer的内容写到Stdout
	io.Copy(os.Stdout, &b)

}
