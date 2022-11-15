package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-StringBuilder"
)

func main() {

	builder := go_StringBuilder.New()
	s := builder.AppendString("test").AppendString(" a is: ").AppendInt(100).String()
	fmt.Println(s) // test a is: 100

}
