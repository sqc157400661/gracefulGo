# Go编码注意事项和建议(学习分享版)

## 控制结构的疑惑
### (1) 
例如，下列程序会打印什么?。
```
package main

import "fmt"

func main() {
	switch { 
	case true:  fmt.Println("true")
	case false: fmt.Println("false")
	}
}
```
### (2) 
下面程序的结果会打印什么？`true`还是`false`？ 
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
原因解析:[Go代码断行规则](./Go代码断行规则.md#duanhang)

## 关于panic和recover的疑惑

### (1) 下面程序的结果会打印什么？
```go
package main

import "fmt"

func main() {
	fmt.Println("嗨！")
	defer func() {
		r := recover()
		fmt.Println("恐慌恢复 2：", r)
	}()
	defer func() {
		errorHandle()
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("哈哈")
}

func errorHandle(){
	if r:=recover();r!=nil{
		fmt.Println("恐慌恢复 1：", r)
	}
}
```

### (2) 
下面程序的结果会打印什么？
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



