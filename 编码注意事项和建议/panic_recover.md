# 恐慌（panic）和恢复（recover）

## 基础知识
Go不支持异常抛出和捕获，而是推荐使用返回值显式返回错误。 不过，Go支持一套和异常抛出/捕获类似的机制。此机制称为**恐慌/恢复（panic/recover）机制**。

我们可以调用内置函数`panic`来产生一个恐慌以使当前协程进入恐慌状况,一旦一个函数调用产生一个恐慌，此函数调用将立即进入返回阶段。

协程中的一个恐慌可以使用内置函数`recover`和延迟函数`defer`消除，从而使得当前协程重新进入正常状况。(recover 仅在延迟函数 defer 中有效)

如果一个协程在恐慌状况下退出，它将使整个程序崩溃。

内置函数panic和recover的声明原型如下：
```go
func panic(v interface{})
func recover() interface{}
```

一个`recover`函数的返回值为其所恢复的恐慌在产生时被一个`panic`函数调用所消费的参数。

下面这个例子展示了如何产生一个恐慌和如何消除一个恐慌。
```go
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}
```
它的输出结果：
```
嗨！
恐慌被恢复了： 拜拜！
正常退出
```

下面的例子在一个新协程里面产生了一个恐慌，并且此协程在恐慌状况下退出，所以整个程序崩溃了。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}
```

运行之，输出如下：
```
hi!
panic: 123

goroutine 5 [running]:
...
```

除了主动`panic`,Go运行时（runtime）会在若干情形下产生恐慌，比如一个整数被0除的时候。下面这个程序将崩溃退出。
```go
package main

func main() {
	a, b := 1, 0
	_ = a/b
}
```

它的输出：
```
panic: runtime error: integer divide by zero

goroutine 1 [running]:
...
```

一般说来，恐慌用来表示正常情况下不应该发生的逻辑错误。 如果这样的一个错误在运行时刻发生了，则它肯定是由于某个bug引起的。 另一方面，非逻辑错误是现实中难以避免的错误，它们不应该导致恐慌。 我们必须被正确地对待和处理非逻辑错误。

更多可能由Go运行时产生的恐慌将在以后其它文章中提及。

以后，我们可以了解一些恐慌/恢复用例和更多关于恐慌/恢复机制的细节。


## recover函数的细节
我们先看一个例子：
```go
package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			fmt.Println("7:", recover())
		}()
	}()
	defer func() {
		func() {
			fmt.Println("6:", recover())
		}()
	}()
	func() {
		defer func() {
			fmt.Println("1:", recover())
		}()
	}()
	func() {
		defer fmt.Println("2:", recover())
	}()
	func() {
		fmt.Println("3:", recover())
	}()
	fmt.Println("4:", recover())
	defer fmt.Println("5:", recover())
	panic(789)
	defer func() {
		fmt.Println("0:", recover())
	}()
}
```
运行之，我们将发现上例中的7个`recover`函数调用都没有恢复程序中产生的恐慌。 此程序的输出结果如下：
```
1: <nil>
2: <nil>
3: <nil>
4: <nil>
5: <nil>
6: <nil>
7: <nil>
panic: 789

goroutine 1 [running]:
...
```
显然地，标号为0的recover调用（代码中最后一个）是执行不到的。 其它7个均执行到了，但是它们的返回值都是nil。为什么呢？让我们先阅读一下[Go白皮书中列出的规则](https://golang.google.cn/ref/spec#Handling_panics)：

    在下面的情况下，recover函数调用的返回值为nil：
    1、传递给相应panic函数调用的实参为nil；
    2、当前协程并没有处于恐慌状态；
    3、recover函数并未直接在一个延迟函数调用中调用。
    
这里我们忽略第一种情况。
上例中标号为1/2/3/4的recover调用都属于第二种情况，标号为5和6的recover调用都属于第三种情况。 
但是，白皮书中列出的三种情况都没有解释为什么标号为7的recover调用也没有恢复程序中的恐慌。

我们知道下面的程序中的recover调用将捕获到panic调用抛出的恐慌。
```go
// example2.go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println( recover() ) // 1
	}()

	panic(1)
}
```
但是，这个简短的程序中的recover调用和前面的例子中的标号为7的recover调用有何本质区别呢？

首先，让我们了解一些概念和事实。


### 恐慌深度 - 恐慌只能向更浅的函数深度传播

恐慌只会从一个函数传播到此函数的调用函数，而从不会传播到深度更深的被此函数调用的函数中。
```go
package main

import "fmt"

func main() { // 调用深度为0
	defer func() { // 调用深度为1
		fmt.Println("当前恐慌深度为0（执行深度为1）")
		func() { // 调用深度为2
			fmt.Println("当前恐慌深度为0（执行深度为2）")
			func() { // 调用深度为3
				fmt.Println("当前恐慌深度为0")
			}()
		}()
	}()

	defer fmt.Println("当前恐慌深度为0（执行深度为0）")

	func() { // 调用深度为1
		defer fmt.Println("当前恐慌深度为1（执行深度为1）")
		func() { // 调用深度为2
			defer fmt.Println("当前恐慌深度为2")
			func() { // 调用深度为3
				defer fmt.Println("当前恐慌深度为3")
				panic(1)
			}()
		}()
	}()
}
```
所以，一个恐慌的深度总是单调递减的，它从不增加。另外，一个协程的恐慌深度从不会小于它的执行深度。

#### ps:
- 调用深度:相对于当前协程的入口调用的调用深度。 如:对于主协程中的一个函数调用，它的调用深度是相对于main入口函数。

### 新生成的恐慌将压制同一深度的老的恐慌
一个例子：
```go
package main

import "fmt"

func main() {
	defer fmt.Println("程序退出时未崩溃")

	defer func() {
		fmt.Println( recover() ) // 3
	}()

	defer fmt.Println("恐慌3将压制恐慌2")
	defer panic(3)
	defer fmt.Println("恐慌2将压制恐慌1")
	defer panic(2)
	panic(1)
}
```

输出：
```
恐慌2将压制恐慌1
恐慌3将压制恐慌2
3
程序退出时未崩溃
```

在此例中，恐慌1被恐慌2压制了，恐慌2又被恐慌3压制了。所以，**最后被捕获的恐慌值为3**。

在**一个协程中**，任何调用**同一深度**上**最多只能有一个活动的恐慌共存**。 特别地，当一个协程的当前执行深度为0时，此协程中只能存在一个活动的恐慌。

### 一个协程中可以有多个活动的恐慌共存
一个例子：
```go
package main

import "fmt"

func main() { // 调用深度为0
	defer fmt.Println("程序崩溃了，因为退出时恐慌3依然未恢复")

	defer func() { // 调用深度为1
		defer func() { // 调用深度为2
			// 恐慌6被消除了。
			fmt.Println( recover() ) // 6
		}()

		// 恐慌3的深度为0，恐慌6的深度为1。
		defer fmt.Println("现在，恐慌3和恐慌6共存")
		defer panic(6) // 将压制恐慌5
		defer panic(5) // 将压制恐慌4
		panic(4) // 不会压制恐慌3，因为恐慌4和恐慌3的深度
		         // 不同。恐慌3为0，而恐慌4的深度为1。
	}()

	defer fmt.Println("现在，只存在恐慌3")
	defer panic(3) // 将压制恐慌2
	defer panic(2) // 将压制恐慌1
	panic(1)
}
```

在这个例子中，两个曾经共存过的恐慌之一（恐慌6）被恢复了。 但是恐慌3在程序退出时仍然没有被恢复，所以此程序在退出时崩溃了。

输出：
```
现在，只存在恐慌3
现在，恐慌3和恐慌6共存
6
程序崩溃了，因为退出时恐慌3依然未恢复
panic: 1
	panic: 2
	panic: 3

goroutine 1 [running]:
...
```

### 事实：更深的恐慌可能被先恢复也可能被后恢复
一个例子：
```go
package main

import "fmt"

func demo(recoverHighestPanicAtFirst bool) {
	fmt.Println("====================")
	defer func() {
		if !recoverHighestPanicAtFirst{
			// 恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer func() {
			//  恢复恐慌2
			fmt.Println("恐慌", recover(), "被恢复了")
		}()
		if recoverHighestPanicAtFirst {
			//  恢复恐慌1
			defer fmt.Println("恐慌", recover(), "被恢复了")
		}
		defer fmt.Println("现在有两个恐慌共存")
		panic(2)
	}()
	panic(1)
}

func main() {
	demo(true)
	demo(false)
}
```

输出:
```
====================
现在有两个恐慌共存
恐慌 1 被恢复了
恐慌 2 被恢复了
====================
现在有两个恐慌共存
恐慌 2 被恢复了
恐慌 1 被恢复了
```

那么，一个recover调用发挥作用的基本规则是什么呢？

基本规则很简单：
1. recover在调用深度为d的延迟函数里
2. 有一个深度为d-1的恐慌存在
