package main

import "fmt"

func main() {
	var x interface{} = true
	_, _ = x.(int) // 断言失败，但不会导致恐慌。
	_,a := x.(int)    // 断言失败，并导致一个恐慌。
	fmt.Println(a)
}
