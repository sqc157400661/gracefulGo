# 一定要执行的规范
## 1、格式化代码
所有代码在提交代码库之前均使用`gofmt`进行格式化修正。

注意：部分IDE可以配置。



## 2、注释
注释必须是**完整的句子**，句子的结尾应该用句号作为结尾（英文句号）【这样做，能使注释在转化成`godoc`时有一个不错的格式】。

注释推荐用英文。

## 3、命名
### （1）、包名
- **全部小写**。没有大写或下划线。
- 大多数使用命名导入的情况下，不需要重命名。
- **简短而简洁**。请记住，在每个使用的地方都完整标识了该名称。
- **不用复数**。例如net/url，而不是net/urls。
- 不要用“util”，“shared”或“lib”。这些是不好的，信息量不足的名称。
- 在引包的时候，需要注意不要使用相对路径，而应该使用**绝对路径**。
### （2）、变量名
 - 驼峰式命名,首字母小写,如`mixedCaps`
 - 变量命名应该尽可能短，尤其是局部变量。
 - 特殊的变量以及全局变量，我们可能需要对它有更多的描述，使用长命名是个不错的建议。
### （3）、函数名
- 驼峰式命名，名字可以长但是得把功能，必要的参数描述清楚，函数名应当是动词或动词短语，如 `postPayment、deletePage、save`。
- 例外:为了对相关的测试用例进行分组，函数名可能包含下划线，如：`TestMyFunction_WhatIsBeingTested`。
### （4）、函数返回值命名
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func (n *Node) Parent1() *Node
func (n *Node) Parent2() (*Node, error)
```
这是一个不好的代码风格，我们只知道函数返回的类型，
但不知道每个返回值的名字
</td><td>

```go
func (n *Node) Parent1() (node *Node)
func (n *Node) Parent2() (node *Node, err error)
```
这条建议几乎不需要过多的解释。尤其对于一种场景，
当你需要在函数结束的defer中对返回值做一些事情，
给返回值名字实在是太必要了。
</td></tr>
</tbody></table>

### （5）、结构体名
- 结构体名应该是名词或名词短语，如 `Custome、WikiPage、Account、AddressParser`。
- 类名不应当是动词,避免使用 `Manager、Processor、Data、Info`这样的类名。
- 属性和接收者方法，大写开头表示public，小写开头表示private。
### （6）、接口命名
- 单个函数的接口名以”er”作为后缀，如 Reader,Writer。接口的实现则去掉“er”。

## 4、对于未导出的顶层常量和变量，使用_作为前缀
在未导出的顶级`vars`和`consts`， 前面加上前缀_，以使它们在使用时明确表示它们是全局符号。

例外：未导出的错误值，应以`err`开头。

基本依据：顶级变量和常量具有包范围作用域。使用通用名称可能很容易在其他文件中意外使用错误的值。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
```

</td><td>

```go
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

</td></tr>
</tbody></table>

## 5、结构体中的嵌入

嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
type Client struct {
  version int
  http.Client
}
```

</td><td>

```go
type Client struct {
  http.Client

  version int
}
```

</td></tr>
</tbody></table>

内嵌应该提供切实的好处，比如以语义上合适的方式添加或增强功能。
它应该在对当前系统没有不利影响（另请参见：`避免在公共结构中嵌入类型`[Avoid Embedding Types in Public Structs]）。

  [Avoid Embedding Types in Public Structs]: #avoid-embedding-types-in-public-structs

嵌入 **不应该**:

- 纯粹是为了美观或方便。
- 使外部类型更难构造或使用。
- 影响外部类型的零值。如果外部类型有一个有用的零值，则在嵌入内部类型之后应该仍然有一个有用的零值。
- 作为嵌入内部类型的副作用，从外部类型公开不相关的函数或字段。
- 公开未导出的类型。
- 影响外部类型的复制形式。
- 更改外部类型的API或类型语义。
- 嵌入内部类型的非规范形式。
- 公开外部类型的实现详细信息。
- 允许用户观察或控制类型内部。
- 通过包装的方式改变内部函数的一般行为，这种包装方式会给用户带来一些意料之外情况。

简单地说，有意识地和有目的地嵌入。一种很好的测试体验是，
"是否所有这些导出的内部方法/字段都将直接添加到外部类型"
如果答案是`some`或`no`，不要嵌入内部类型-而是使用字段。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
type A struct {
    // Bad: A.Lock() and A.Unlock() 现在可用
    // 不提供任何功能性好处，并允许用户控制有关A的内部细节。
    sync.Mutex
}
```

</td><td>

```go
type countingWriteCloser struct {
    // Good: Write() 在外层提供用于特定目的，
    // 并且委托工作到内部类型的Write()中。
    io.WriteCloser
    count int
}
func (w *countingWriteCloser) Write(bs []byte) (int, error) {
    w.count += len(bs)
    return w.WriteCloser.Write(bs)
}
```

</td></tr>
<tr><td>

```go
type Book struct {
    // Bad: 指针更改零值的有用性
    io.ReadWriter
    // other fields
}
// later
var b Book
b.Read(...)  // panic: nil pointer
b.String()   // panic: nil pointer
b.Write(...) // panic: nil pointer
```

</td><td>

```go
type Book struct {
    // Good: 有用的零值
    bytes.Buffer
    // other fields
}
// later
var b Book
b.Read(...)  // ok
b.String()   // ok
b.Write(...) // ok
```

</td></tr>
<tr><td>

```go
type Client struct {
    sync.Mutex
    sync.WaitGroup
    bytes.Buffer
    url.URL
}
```

</td><td>

```go
type Client struct {
    mtx sync.Mutex
    wg  sync.WaitGroup
    buf bytes.Buffer
    url url.URL
}
```

</td></tr>
</tbody></table>

## 6、初始化结构体时指定字段名
有3个字段及以上时：初始化结构体时，指定字段名称
<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
k := User{"John", "Doe", true}
```

</td><td>

```go
k := User{
    FirstName: "John",
    LastName: "Doe",
    Admin: true,
}
```

</td></tr>
</tbody></table>

## 7、nil 是一个有效的 slice

`nil` 是一个有效的长度为 0 的 slice，这意味着，

- 您不应明确返回长度为零的切片。应该返回`nil` 来代替。

  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  if x == "" {
    return []int{}
  }
  ```

  </td><td>

  ```go
  if x == "" {
    return nil
  }
  ```

  </td></tr>
  </tbody></table>

- 要检查切片是否为空，请始终使用`len(s) == 0`。而非 `nil`。

  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  func isEmpty(s []string) bool {
    return s == nil
  }
  ```

  </td><td>

  ```go
  func isEmpty(s []string) bool {
    return len(s) == 0
  }
  ```

  </td></tr>
  </tbody></table>

- 零值切片（用`var`声明的切片）可立即使用，无需调用`make()`创建。

  <table>
  <thead><tr><th>Bad</th><th>Good</th></tr></thead>
  <tbody>
  <tr><td>

  ```go
  nums := []int{}
  // or, nums := make([]int)

  if add1 {
    nums = append(nums, 1)
  }

  if add2 {
    nums = append(nums, 2)
  }
  ```

  </td><td>

  ```go
  var nums []int

  if add1 {
    nums = append(nums, 1)
  }

  if add2 {
    nums = append(nums, 2)
  }
  ```

  </td></tr>
  </tbody></table>

记住，虽然nil切片是有效的切片，但它不等于长度为0的切片（一个为nil，另一个不是），并且在不同的情况下（例如序列化），这两个切片的处理方式可能不同。

## 7、初始化 Struct 引用

在初始化结构引用时，请使用`&T{}`代替`new(T)`，以使其与结构体初始化一致。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
sval := T{Name: "foo"}

// inconsistent
sptr := new(T)
sptr.Name = "bar"
```

</td><td>

```go
sval := T{Name: "foo"}

sptr := &T{Name: "bar"}
```

</td></tr>
</tbody></table>

## 8、初始化 Maps

对于空 map 请使用 `make(..)` 初始化， 并且 map 是通过编程方式填充的。
这使得 map 初始化在表现上不同于声明，并且它还可以方便地在 make 后添加大小提示。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
var (
  // m1 读写安全;
  // m2 在写入时会 panic
  m1 = map[T1]T2{}
  m2 map[T1]T2
)
```

</td><td>

```go
var (
  // m1 读写安全;
  // m2 在写入时会 panic
  m1 = make(map[T1]T2)
  m2 map[T1]T2
)
```

</td></tr>
<tr><td>

声明和初始化看起来非常相似的。

</td><td>

声明和初始化看起来差别非常大。

</td></tr>
</tbody></table>

在尽可能的情况下，请在初始化时提供 map 容量大小，详细请看 [指定Map容量提示](#指定Map容量提示)。


另外，如果 map 包含固定的元素列表，则使用 map literals(map 初始化列表) 初始化映射。


<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
m := make(map[T1]T2, 3)
m[k1] = v1
m[k2] = v2
m[k3] = v3
```

</td><td>

```go
m := map[T1]T2{
  k1: v1,
  k2: v2,
  k3: v3,
}
```

</td></tr>
</tbody></table>

基本准则是：在初始化时使用 map 初始化列表 来添加一组固定的元素。否则使用 `make` (如果可以，请尽量指定 map 容量)。

## 9、优先使用 strconv 而不是 fmt

将原语转换为字符串或从字符串转换时，`strconv`速度比`fmt`快。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
for i := 0; i < b.N; i++ {
  s := fmt.Sprint(rand.Int())
}
```

</td><td>

```go
for i := 0; i < b.N; i++ {
  s := strconv.Itoa(rand.Int())
}
```

</td></tr>
<tr><td>

```
BenchmarkFmtSprint-4    143 ns/op    2 allocs/op
```

</td><td>

```
BenchmarkStrconv-4    64.2 ns/op    1 allocs/op
```

</td></tr>
</tbody></table>



### 10、避免字符串到字节的转换

不要反复从固定字符串创建字节 slice。相反，请执行一次转换并捕获结果。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
for i := 0; i < b.N; i++ {
  w.Write([]byte("Hello world"))
}
```

</td><td>

```go
data := []byte("Hello world")
for i := 0; i < b.N; i++ {
  w.Write(data)
}
```

</tr>
<tr><td>

```
BenchmarkBad-4   50000000   22.2 ns/op
```

</td><td>

```
BenchmarkGood-4  500000000   3.25 ns/op
```

</td></tr>
</tbody></table>


## 11、指定Map容量提示

在尽可能的情况下，在使用 `make()` 初始化的时候提供容量信息

```go
make(map[T1]T2, hint)
```

向`make()`提供容量提示会在初始化时尝试调整map的大小，这将减少在将元素添加到map时为map重新分配内存。


注意，与slices不同。map capacity提示并不保证完全的抢占式分配，而是用于估计所需的hashmap bucket的数量。
因此，在将元素添加到map时，甚至在指定map容量时，仍可能发生分配。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
m := make(map[string]os.FileInfo)

files, _ := ioutil.ReadDir("./files")
for _, f := range files {
    m[f.Name()] = f
}
```

</td><td>

```go

files, _ := ioutil.ReadDir("./files")

m := make(map[string]os.FileInfo, len(files))
for _, f := range files {
    m[f.Name()] = f
}
```

</td></tr>
<tr><td>

`m` 是在没有大小提示的情况下创建的； 在运行时可能会有更多分配。

</td><td>

`m` 是有大小提示创建的；在运行时可能会有更少的分配。

</td></tr>
</tbody></table>

## 12、指定切片容量

在尽可能的情况下，在使用`make()`初始化切片时提供容量信息，特别是在追加切片时。

```go
make([]T, length, capacity)
```

与maps不同，slice capacity不是一个提示：编译器将为提供给`make()`的slice的容量分配足够的内存，
这意味着后续的append()`操作将导致零分配（直到slice的长度与容量匹配，在此之后，任何append都可能调整大小以容纳其他元素）。

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
for n := 0; n < b.N; n++ {
  data := make([]int, 0)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}
```

</td><td>

```go
for n := 0; n < b.N; n++ {
  data := make([]int, 0, size)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}
```

</td></tr>
<tr><td>

```
BenchmarkBad-4    100000000    2.48s
```

</td><td>

```
BenchmarkGood-4   100000000    0.21s
```

</td></tr>
</tbody></table>

## 13、断言要用2个参数接受,否则失败会导致恐慌
```go
package main

func main() {
	var x interface{} = true
	_, _ = x.(int) // 断言失败，但不会导致恐慌。
	_ = x.(int)    // 断言失败，并导致一个恐慌。
}
```

## 14、主动停止time.Ticker
当一个time.Timer值不再被使用，一段时间后它将被自动垃圾回收掉。 但对于一个不再使用的time.Ticker值，我们必须调用它的Stop方法结束它，否则它将永远不会得到回收。

## 15、切片赋值或者获取范围大的在前
Go运行时将检查操作中使用的下标是否越界。 如果下标越界，一个恐慌将产生，以防止这样的操作破坏内存安全。这样的检查称为边界检查。 边界检查使得我们的代码能够安全地运行；但是另一方面，也使得我们的代码运行效率略微降低。

通过编码的可以消除边界检查来提高性能，如下
```go
func f1(s []int) {
	_ = s[0] // 第5行： 需要边界检查
	_ = s[1] // 第6行： 需要边界检查
	_ = s[2] // 第7行： 需要边界检查
}

func f2(s []int) {
	_ = s[2] // 第11行： 需要边界检查
	_ = s[1] // 第12行： 边界检查消除了！
	_ = s[0] // 第13行： 边界检查消除了！
}
```
