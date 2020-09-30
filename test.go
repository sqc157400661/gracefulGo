package main
import "fmt"
func main() {
	// 假设此切片的长度很大，以至于它的元素
	// 将被开辟在堆上。
	bs := make([]byte, 1 << 31)
	// 一个聪明的编译器将觉察到bs的底层元素
	// 部分已经不会再被使用，而正确地认为bs的
	// 底层元素部分在此刻可以被安全地回收了。
	fmt.Println(len(bs))
}