package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func showNumber (i int) {
	defer fmt.Println(1111)
	//runtime.Goexit()
	os.Exit(1)
	fmt.Println(i)
}

func main() {

	for i := 0; i < 10; i++ {
		go showNumber(i)
	}

	runtime.Gosched()
	fmt.Println("Haha")
}

func longRunning(messages <-chan string) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()
	for {
		select {
		case <-timer.C: // 过期了
			return
		case msg := <-messages:
			fmt.Println(msg)
			// 此if代码块很重要。
			if !timer.Stop() {
				<-timer.C
			}
		}
		// 必须重置以复用。
		timer.Reset(time.Minute)
	}
}