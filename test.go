package main

import "runtime"

func DoSomething() {
	for {
		// 做点什么...

		runtime.Gosched() // 防止本协程霸占CPU不放
	}
}

func main() {
	go DoSomething()
	go DoSomething()
	select{}
}