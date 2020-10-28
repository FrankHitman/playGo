### defer-panic-and-recover
Go 没有像 Java 和 .NET 那样的 try/catch 异常机制：不能执行抛异常操作。但是有一套 defer-panic-and-recover 机制
处理错误并且在函数发生错误的地方给用户返回错误信息：照这样处理就算真的出了问题，你的程序也能继续运行并且通知给用户。
panic and recover 是用来处理真正的异常（无法预测的错误）而不是普通的错误。

### What is go panicking：
在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，
所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者。
这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 panicking。

### what the using of recover
正如名字一样，recover内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。