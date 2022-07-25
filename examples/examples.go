package main

import (
	"fmt"
	"github.com/CC11001100/go-StringBuilder/pkg/string_builder"
)

func main() {

	builder := string_builder.New()
	s := builder.AppendString("test").AppendString(" a is: ").AppendInt(100).String()
	fmt.Println(s) // test a is: 100

}
