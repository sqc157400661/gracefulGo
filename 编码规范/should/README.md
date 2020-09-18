# 强烈建议执行的规范
## 1、代码行长度
一行最长**不超过80个字符**，超过的使用换行展示，尽量保持格式优雅

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


