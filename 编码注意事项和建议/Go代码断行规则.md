# Go代码断行规则
本文将列出Go代码中的详细断行规则。

## 分号插入规则
我们在Go编程中常遵循的一个规则是：一个显式代码块的起始左大括号`{`不放在下一行。 比如，下面这个`for`循环代码块编译将失败。
```go
for i := 5; i > 0; i--
{ // error: 未预料到的新行
}
```
为了让上面这个for循环代码块编译成功，我们不能在起始左大括号`{`前断行，而应该像下面这样进行修改：
```go
for i := 5; i > 0; i-- {
}
```

然而，有时候起始左大括号{却可以放在一个新行上，比如下面这个for循环代编译时没有问题的。
```go
for
{
    // do something ...
}
```
那么，Go代码中的根本性换行规则究竟是如何定义的呢？ 在回答这个问题之前，我们应该知道一个事实：正式的Go语法是使用分号`;`做为结尾标识符的。 但是，我们很少在Go代码中使用和看到分号。为什么呢？原因是大多数分号都是可选的，因此它们常常被省略。 在编译时刻，Go编译器会自动插入这些省略的分号。

比如，下面这个程序中的十个分号都是可以被省略掉的。
```go
package main;

import "fmt";

func main() {
	var (
		i   int;
		sum int;
	);
	for i < 6 {
		sum += i;
		i++;
	};
	fmt.Println(sum);
};
```

假设上面这个程序存储在一个`semicolons.go`文件中，我们可以运行`go fmt semicolons.go`将此程序中的不必要的分号去除掉。 在编译时刻，编译器会自动此插入这些去除掉的分号（至此文件的内存中的版本）。

<a name="duanhang">自动插入分号的规则是什么呢？</a> [Go白皮书这样描述](https://golang.google.cn/ref/spec#Semicolons)：
```
1、在Go代码中，注释除外，如果一个代码行的最后一个语法词段（token）为下列所示之一，
则一个分号将自动插入在此字段后（即行尾）：
    1、一个标识符；
    2、 一个整数、浮点数、虚部、码点或者字符串字面表示形式；
    3、这几个跳转关键字之一：break、continue、fallthrough和return；
    4、自增运算符++或者自减运算符--；
    5、一个右括号：)、]或}。
2、为了允许一条复杂语句完全显示在一个代码行中，分号可能被插入在一个右小括号)或者右大括号}之前。
```

对于上述第一条规则描述的情形，我们当然也可以手动插入这些分号，就像此前的例子中所示。换句话说，这些分号在编程时是可选的。


上述第二条规则允许我们写出如下的代码：
```go
import (_ "math"; "fmt")
var (a int; b string)
const (M = iota; N)
type (MyInt int; T struct{x bool; y int32})
type I interface{m1(int) int; m2() string}
func f() {print("a"); panic(nil)}
```
编译器在编译时刻将自动插入所需的分号，如下所示：
```go
var (a int; b string;);
const (M = iota; N;);
type (MyInt int; T struct{x bool; y int32;};);
type I interface{m1(int) int; m2() string;};
func f() {print("a"); panic(nil);};
```
编译器不会为其它任何情形插入分号。如果其它任何情形需要一个分号，我们必须手动插入此分号。 比如，上例中的每行中的第一个分号必须手动插入。下例中的分号也都需要手动插入。
```go
var a = 1; var b = true
a++; b = !b
print(a); print(b)
```

从以上两条规则可以看出，一个分号永远不会插入在`for`关键字后，这就是为什么上面的裸`for`循环例子是合法的原因。

分号自动插入规则导致的一个结果是：自增和自减运算必须呈现为单独的语句，它们不能被当作表达式使用。 比如，下面的代码是编译不通过的：
```go
func f() {
	a := 0
	println(a++)
	println(a--)
}
```
上面代码编译不通过的原因是它等价于下面的代码：

```go
func f() {
	a := 0
	println(a++;)
	println(a--;)
}
```
分号自动插入规则导致的另一个结果是：我们不能在选择器中的句点`.`之前断行。 在选择器中的句点之后断行是允许的，比如：
```go
anObject.
    MethodA().
    MethodB().
    MethodC()
```
而下面这样是非法的：
```go
anObject
    .MethodA()
    .MethodB()
    .MethodC()
```
此代码片段是非法的原因是编译器将自动在每个右小括号)后插入一个分号，如下面所示：
```go
anObject;
    .MethodA();
    .MethodB();
    .MethodC();
```
上述分号自动插入规则可以让我们写出更简洁的代码，同时也允许我们写出一些合法的但看上去有些怪异的代码，比如：
```go
package main

import "fmt"

func alwaysFalse() bool {return false}

func main() {
	for
	i := 0 // 第9行
	i < 6  // 第10行
	i++ {
		// 使用i ...
	}

	if x := alwaysFalse()   // 第15行 
	!x {
		// ...
	}

	switch alwaysFalse()   // // 第20行
	{
	case true: fmt.Println("true")
	case false: fmt.Println("false")
	}
}
```
上例中所有的流程控制代码块都是合法的。编译器将在这些行的行尾自动插入一个分号：第9行、第10行、第15行和第20行。

注意，上例中的`switch-case`代码块将输出`true`，而不是`false`。 此代码块和下面这个是不同的：
```go
switch alwaysFalse() {
case true: fmt.Println("true")
case false: fmt.Println("false")
}
```
如果我们使用`go fmt`命令格式化前者，一个分号将自动添加到`alwaysFalse()`函数调用之后，如下所示：
```go
switch alwaysFalse();
{
 case true: fmt.Println("true")
 case false: fmt.Println("false")
 }
```

插入此分号后，此代码块将和下者等价：
```go
switch alwaysFalse(); true {
case true: fmt.Println("true")
case false: fmt.Println("false")
}
```
这就是它为什么输出true的原因。

常使用`go fmt`和`go vet`命令来格式化和发现可能的逻辑错误是一个好习惯。

下面是一个很少见的情形，此情形中所示的代码看上去是合法的，但是实际上是编译不通过的。
```go
func f() {
	switch x {
	case 1:
	{
		goto A
		A: // 这里编译没问题
	}
	case 2:
		goto B
		B: // syntax error: 跳转标签后缺少语句
	case 0:
		goto C
		C: // 这里编译没问题
	}
}
```

编译错误信息表明跳转标签的声明之后必须跟一条语句。 但是，看上去，上例中的三个标签声明没什么不同，
它们都没有跟随一条语句。 那为什么只有`B:`标签声明是不合法的呢？ 原因是，根据上述第二条分号自动插入规则，
编译器将在`A:`和`C:`标签声明之后的右大括号}字符之前插入一个分号，如下所示：
```go
func f(x int) {
	switch x {
	case 1:
	{
		goto A
		A:
	;} // 一个分号插入到了这里
	case 2:
		goto B
		B: // syntax error: 跳转标签后缺少语句
	case 0:
		goto C
		C:
	;} // 一个分号插入到了这里
}
```
一个单独的分号实际上表示一条空语句。 
这就是为什么`A:`和`C:`标签声明之后确实跟随了一条语句的原因，所以它们是合法的。 
而B:标签声明跟随的`case 0:`不是一条语句，所以它是不合法的。

我们可以在`B:`标签声明之后手动插入一个分号使之变得合法。


## 逗号,从不会被自动插入

一些包含多个类似项目的语法形式多用逗号,来做为这些项目之间的分割符，比如组合字面形式和函数参数列表等。 在这样的一个语法形式中，最后一个项目后总可以跟一个可选的逗号。 如果此逗号为它所在代码行的最后一个有效字符，则此逗号是必需的；否则，此逗号可以省略。 编译器在任何情况下都不会自动插入逗号。

比如，下面的代码是合法的：
```go
func f1(a int, b string,) (x bool, y int,) {
	return true, 789
}
var f2 func (a int, b string) (x bool, y int)
var f3 func (a int, b string, // 最后一个逗号是必需的
) (x bool, y int,             // 最后一个逗号是必需的
)
var _ = []int{2, 3, 5, 7, 9,} // 最后一个逗号是可选的
var _ = []int{2, 3, 5, 7, 9,  // 最后一个逗号是必需的
}
var _ = []int{2, 3, 5, 7, 9}
var _, _ = f1(123, "Go",) // 最后一个逗号是可选的
var _, _ = f1(123, "Go",  // 最后一个逗号是必需的
)
var _, _ = f1(123, "Go")
```

而下面这段代码是不合法的，因为编译器将自动在每一行的行尾插入一个分号（除了第二行）。 其中三行在插入分号后将导致编译错误。
```go
func f1(a int, b string,) (x bool, y int // error
) {
	return true, 789
}
var _ = []int{2, 3, 5, 7, 9 // error: unexpected newline
}
var _, _ = f1(123, "Go" // error: unexpected newline
)
```
## 结束语
和很多Go中的其它设计细节一样，Go代码断行规则设计的评价也是褒贬不一。 有些程序员不太喜欢这样的断行规则，因为这样的规则限制了代码风格的自由度。 但是这些规则不但使得代码编译速度大大提高，另一方面也使得不同Go程序员写出的代码风格大体一致，从而相互可以比较轻松地读懂对方的代码。