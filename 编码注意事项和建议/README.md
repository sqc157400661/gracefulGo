# Go编码注意事项和建议

## 控制结构
### (1) switch流程控制代码块中的switch表达式的缺省默认值为类型确定值true（其类型为预定义类型bool）
例如，下列程序会打印出`true`。
```
package main

import "fmt"

func main() {
	switch { // <=> switch true {
	case true:  fmt.Println("true")
	case false: fmt.Println("false")
	}
}
```
### (2) switch后接函数调用要注意
下面程序的结果会打印什么？`true`还是`false`？ 答案是`true`。
```
package main

import "fmt"

func False() bool {
	return false
}

func main() {
	switch False()
	{
	case true:  fmt.Println("true")
	case false: fmt.Println("false")
	}
}
```
原因请阅读:[Go代码断行规则](./Go代码断行规则.md#duanhang)

### (3) os.Exit 和 runtime.Goexit
1. `os.Exit`:函数从任何函数里退出一个程序。 `os.Exit`函数调用接受一个int代码值做为参数并将此代码返回给操作系统。
1. `runtime.Goexit`函数退出一个goroutine。 `runtime.Goexit`函数没有参数
示例1:
```go
// exit-example.go
package main

import "os"
import "time"

func main() {
	go func() {
		time.Sleep(time.Second)
		os.Exit(1)
	}()
	select{}
}
```
示例2:
```go
package main

import "fmt"
import "runtime"

func main() {
	c := make(chan int)
	go func() {
		defer func() {c <- 1}()
		defer fmt.Println("Go")
		func() {
			defer fmt.Println("C")
			runtime.Goexit()
		}()
		fmt.Println("Java")
	}()
	<-c
}
```
## 切片/数组/map

## 函数
### (1) 嵌套的延迟函数调用可以修改外层函数的返回结果
```go
package main

import "fmt"

func F() (r int) {
	defer func() {
		r = 789
	}()

	return 123 // <=> r = 123; return
}

func main() {
	fmt.Println(F()) // 789
}
```
## 程序异常以及处理
### (1) panic和recover
下面程序的结果会打印什么？panic不会被捕获
```go
package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			fmt.Println("3:", recover())
		}()
	}()
	defer func() {
		func() {
			fmt.Println("2:", recover())
		}()
	}()
	func() {
		defer func() {
			fmt.Println("1:", recover())
		}()
	}()
	panic(121)
}

```
原因解析:[恐慌（panic）和恢复（recover）](./panic_recover.md)

##  其他

### 运算优先级


