package main

import "fmt"

var a = &[]int{1, 2, 3}
var i int
func f() int {
	i = 1
	a = &[]int{7, 8, 9}
	return 0
}

func main() {
	// 表达式"a"、"i"和"f()"的估值顺序未定义。
	(*a)[i] = f()
	fmt.Println(*a)
}