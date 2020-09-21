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
## 切片/数组/map

## 函数

## 程序异常以及处理






