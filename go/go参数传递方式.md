---
layout: post
title: Go中的值类型和引用类型
date: 2017-03-08 15:30:00
tags:
- go
categories: go
---

Go中参数传递的方式和C++有点像。


# C++ 中三种参数传递方式
### 值传递：
最常见的一种传参方式，函数的形参是实参的拷贝，函数中改变形参不会影响到函数外部的形参。一般是函数内部修改参数而又不希望影响到调用者的时候会采用值传递。

### 指针传递
形参是指向实参地址的一个指针，顾名思义，在函数中对形参指向的内容操作，实参本身会被修改。

### 引用传递
在 C++ 中，引用是变量的别名，实际上是同一个东西，在内存中也存在同一个地址。换句话说，不管在哪里对引用操作，都相当直接操作被引用的变量。

下面看 demo:
```cpp
#include <iostream>


//值传递
void func1(int a) {
    std::cout <<  "值传递，变量地址：" <<  &a <<  ", 变量值：" << a << std::endl;
    a ++ ;
}

//指针传递
void func2 (int* a) {
    std::cout <<  "指针传递，变量地址：" <<  a <<  ", 变量值：" << *a << std::endl;
    *a = *a + 1;
}

//引用传递
void func3 (int& a) {
    std::cout <<  "指针传递，变量地址：" <<  &a <<  ", 变量值：" << a << std::endl;
    a ++;
}

int main() {
    int a = 5;

    std::cout <<  "变量实际地址：" <<  &a <<  ", 变量值：" << a << std::endl;
    func1(a);
    std::cout << "值传递操作后,变量值：" << a << std::endl;

    std::cout <<  "变量实际地址：" <<  &a <<  ", 变量值：" << a << std::endl;
    func2(&a);
    std::cout << "指针传递操作后,变量值：" << a << std::endl;

    std::cout <<  "变量实际地址：" <<  &a <<  ", 变量值：" << a << std::endl;
    func3(a);
    std::cout << "引用传递操作后,变量值：" << a << std::endl;

    return 0;
}
```
输出结果如下：
```text
变量实际地址：0x28feac, 变量值：5
值传递，变量地址：0x28fe90, 变量值：5
值传递操作后,变量值：5
变量实际地址：0x28feac, 变量值：5
指针传递，变量地址：0x28feac, 变量值：5
指针传递操作后,变量值：6
变量实际地址：0x28feac, 变量值：6
指针传递，变量地址：0x28feac, 变量值：6
引用传递操作后,变量值：7
```

# Go 中的参数传递
上面介绍了 C++ 的三种参数传递方式，值传递和指针传递容易理解，那么 Go 是不是也有这些传参方式呢？这引起过争论，但是对比 C++ 的引用传递的概念，我们可以说，Go 没有引用传递方式。为什么这么说，因为 Go 没有变量的引用这一概念。但是 Go 有引用类型，这个稍后再解释。

先看一个 Go 传值和传指针的例子：
```go
package main

import (
    "fmt"
)


func main() {
    a := 1
    fmt.Println( "变量实际地址:", &a, "变量值:", a)
    func1 (a)
    fmt.Println( "值传递操作后,变量值:", a)
    fmt.Println( "变量实际地址:", &a, "变量值:", a)
    func2(&a)
    fmt.Println( "指针传递操作后,变量值:", a)
}

//值传递
func func1 (a int) {
    a++
    fmt.Println( "值传递，变量地址:", &a, "变量值:", a)
}

//指针传递
func func2 (a *int) {
    *a = *a + 1
    fmt.Println( "指针传递，变量地址:", a, "变量值:", *a)
}
```
输出结果如下：
```text
变量实际地址: 0xc04203c1d0 变量值: 1
值传递，变量地址: 0xc04203c210 变量值: 2
值传递操作后,变量值: 1
变量实际地址: 0xc04203c1d0 变量值: 1
指针传递，变量地址: 0xc04203c1d0 变量值: 2
指针传递操作后,变量值: 2
```
可以看出，Go 基本类型的值传递和指针传递和 C++ 并没有什么不同，但是它没有变量的引用这一概念。那 Go 的引用类型怎么理解呢？

Go 的引用类型

|      Type     |    Value Type   |    Reference Type  |
| :-----------: | :-------------: | :----------------: |
|      int      |       X         |                    |
|     array     |       X         |                    | 
|     Map       |                 |          X         |
|     slice     |                 |          X         |
|    channel    |                 |          X         |

在Go中，引用类型包含切片、字典、通道等。

```go
func main()  {
	m := map[string]int {
		"tom"     : 10,
		"jerry"   : 20,
	}
	updateMap(m)
	fmt.Println(m)
}

func updateMap(m map[string]int)  {
	m["tom"] = 15
}
```
