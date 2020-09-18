# 供大家参考的规范
## 1、Import Dot（Golang特性：import . 包名）
在测试中，我们很可能会使用这个特性，该特性能让我们避免循环引用问题，思考一下下面的例子：
```go
package foo_test

import (
    "bar/testutil" // also imports "foo"
    . "foo"
)
```
以上例子，该测试文件不能定义在于foo包里面，因为它import了`bar/testutil`，而`bar/testutilimport`了foo，这将构成循环引用。

所以我们需要将该测试文件定义在`foo_test`包中。使用了`import . "foo"`后，该测试文件内代码能直接调用foo里面的函数而不需要显式地写上包名。

但`import dot`这个特性，建议只在这种场景下使用，因为它会大大增加代码的理解难度。






