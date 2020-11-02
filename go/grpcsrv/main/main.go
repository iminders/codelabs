package main

import (
	"fmt"

	"github.com/iminders/codelabs/go/grpcsrv/serv"
)

func main() {
	fmt.Println("Hello, 世界")
	msg := serv.Hello()
	fmt.Println(msg)
}
