# 强烈建议执行的规范
## 1、代码行长度
一行最长**不超过80个字符**，超过的使用换行展示，尽量保持格式优雅。

## 2、注释
Go提供两种注释风格：
- 块注释风格`/**/`
- 行注释风格`//`

要求：
- 每一个包都应该有包注释，位于文件的顶部，在包名出现之前。
- 如果一个包有多个文件，包注释只需要出现在一个文件的顶部即可。
- 包注释建议使用块注释风格，如果这个包特别简单，需要的注释很少，也可以选择使用行注释风格。
- 每个public函数都应该有注释，注释句子应该以该函数名开头 [这样做的好处是，但你要查找某个public函数的注释时，grep函数名即可]

示例1
```
/*
	用于播放AV电影的package,包含XXOO等主要的功能方法.
*/
package av
```
示例2
```
// PlayCangJingKongMovie is a function to watch the movices of CangJingKong.
// retrun skill that you can learn.
func PlayCangJingKongMovie(str string) (*skill, error) {
```

## 3、一致性原则
一致性的代码更容易维护、是更合理的、需要更少的学习成本、并且随着新的约定出现或者出现错误后更容易迁移、更新、修复 bug。一句话：让大家认知一致。

<table>
<thead><tr>举例<th></th><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr>
<td>Import一致性</td>
<td>

```go
import (
    "fmt"
    "code.google.com/p/x/y"
    "hash/adler32"
    "appengine/foo"
    "appengine/user"
    "os"
    "github.com/foo/bar"
)
```

</td><td>

```go
import (
    //标准库
    "fmt"
    "hash/adler32"
    "os"

    // 自己库
    "appengine/foo"
    "appengine/user"

    // 第三方库
    "code.google.com/p/x/y"
    "github.com/foo/bar"
)
```

</td></tr>
<tr>
<td>声明一致性</td>
<td>

```go
const a = 1
const b = 2

var a = 1
var b = 2

type Area float64
type Volume float64
```

</td><td>

```go
const (
  a = 1
  b = 2
)

var (
  a = 1
  b = 2
)

type (
  Area float64
  Volume float64
)
```

</td></tr>
<tr>
<td>所属一致性</td>
<td>

```go
func (s *something) Cost() {
  return calcCost(s.weights)
}

type something struct{ ... }

func calcCost(n []int) int {...}

func (s *something) Stop() {...}
// 函数是按接收者分组的，因此普通工具函数应在文件末尾出现
func newSomething() *something {
    return &something{}
}
```

</td><td>

```go
type something struct{ ... }

func newSomething() *something {
    return &something{}
}

func (s *something) Cost() {
  return calcCost(s.weights)
}

func (s *something) Stop() {...}

func calcCost(n []int) int {...}
```

</td></tr>
</tbody></table>

## 4、减少逻辑代码的缩进
代码应通过尽可能先处理错误情况/特殊情况并尽早返回或继续循环来减少缩进或者嵌套
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
if err != nil {
    // error handling.
} else {
    // normal code.
}
```

</td><td>

```go
if err != nil {
    // error handling.
    return 
}
// normal code.
```

</td></tr>
<tr><td>

```go
if x, err := f(); err != nil {
    // error handling.
    return
} else {
    // use x.
}
```

</td><td>

```go
x, err := f()
if err != nil {
    // error handling
    return
}
// use x
```

</td></tr>

</tbody></table>

## 5、Don’t Panic
尽量不要使用panic处理错误。函数应该设计成多返回值，其中包括返回相应的error类型

## 6、处理错误
不要将error赋值给匿名变量_

## 7、可以清除因未重置丢失的切片元素中的指针而造成的临时性内存泄露
```go
func h() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	// 使用此s切片 ...

	s[0], s[len(s)-1] = nil, nil // 重置首尾元素指针
	return s[1:3:3]
}
```

