# Inline
---
### 普通inline
- 不包含复杂过程，可以inline 
- 可以通过注释：//go:noinline 禁止inline

### mid stack
- 允许inline调用过不可inline函数（拥有复杂过程）的函数

### 接口调用
- 直接接口方法调用会存在开销，因为会发生内存逃逸，并且不能inline，而使用具体类型调用或者使用类型断言则不会。 [inline 分析](https://mp.weixin.qq.com/s/1nqpVzitGdVVvHIuk8iAeQ)

> 内存逃逸：大概的意思是在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法，可以分析在程序的哪些地方可以访问到指针。它涉及到指针分析和形状分析。 当一个变量(或对象)在子程序中被分配时，一个指向变量的指针可能逃逸到其它执行线程中，或者去调用子程序。如果使用尾递归优化（通常在函数编程语言中是需要的），对象也可能逃逸到被调用的子程序中。 如果一个子程序分配一个对象并返回一个该对象的指针，该对象可能在程序中的任何一个地方被访问到——这样指针就成功“逃逸”了。如果指针存储在全局变量或者其它数据结构中，它们也可能发生逃逸，这种情况是当前程序中的指针逃逸。 逃逸分析需要确定指针所有可以存储的地方，保证指针的生命周期只在当前进程或线程中。

- go语言build过程中会进行逃逸分析来确定变量的分配