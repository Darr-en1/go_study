package main

import (
	"fmt"
)

var aa = 3

//  bb:=true 包内作用域不适用，局部变量可以这么定义

var (
	bb = 4
	cc = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
}
