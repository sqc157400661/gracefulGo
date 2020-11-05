package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		}
	}
	tryReceive := func()  {
		select {
		case v := <-c: fmt.Println(v)

		}
	}
	go trySend("Hello!") // 发送成功
	go trySend("Hi!")    // 发送成功
	go trySend("Bye!")   // 发送失败，但不会阻塞。
	for i:=0; i<=2;i++  {
		go tryReceive()
	}
	time.Sleep(time.Second*3)
}