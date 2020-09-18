# 一定要执行的规范
## 1、格式化代码
所有代码在提交代码库之前均使用`gofmt`进行格式化修正。

注意：部分IDE可以配置。



## 2、注释
注释必须是**完整的句子**，句子的结尾应该用句号作为结尾（英文句号）【这样做，能使注释在转化成`godoc`时有一个不错的格式】。

注释推荐用英文。

## 3、命名
### 1、包名
- **全部小写**。没有大写或下划线。
- 大多数使用命名导入的情况下，不需要重命名。
- **简短而简洁**。请记住，在每个使用的地方都完整标识了该名称。
- **不用复数**。例如net/url，而不是net/urls。
- 不要用“util”，“shared”或“lib”。这些是不好的，信息量不足的名称。
- 在引包的时候，需要注意不要使用相对路径，而应该使用**绝对路径**。
### 2、变量名
 - 驼峰式命名,首字母小写,如`mixedCaps`
 - 变量命名应该尽可能短，尤其是局部变量。
 - 特殊的变量以及全局变量，我们可能需要对它有更多的描述，使用长命名是个不错的建议。
### 3、函数名
- 驼峰式命名，名字可以长但是得把功能，必要的参数描述清楚，函数名应当是动词或动词短语，如 `postPayment、deletePage、save`。
- 例外:为了对相关的测试用例进行分组，函数名可能包含下划线，如：`TestMyFunction_WhatIsBeingTested`。
### 4、函数返回值命名
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

### 5、结构体名
- 结构体名应该是名词或名词短语，如 `Custome、WikiPage、Account、AddressParser`。
- 类名不应当是动词,避免使用 `Manager、Processor、Data、Info`这样的类名。
- 属性和接收者方法，大写开头表示public，小写开头表示private。
### 6、接口命名
- 单个函数的接口名以”er”作为后缀，如 Reader,Writer。接口的实现则去掉“er”。