package main
import (
	"fmt"
	"log"
)
import "runtime"
var a string
var done bool
func setup() {
	a = "hello, world"
	done = true
	if done {
		log.Println(len(a)) // 如果被打印出来，它总是12
	}
}
func main() {
	go setup()
	for !done {
		runtime.Gosched()
	}
	log.Println(a) // 期待的打印结果：hello, world

	aa := "1234"
	fmt.Println(aa[0:2])
}