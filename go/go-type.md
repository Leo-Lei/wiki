---
layout: post
title: Go type
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# Go类型系统



|   数据类型     | 引用/值/复合          |   零值(默认值)                                                                 |
| ------------- | --------------------- | ----------------------------------------------------------------------------- |
| int           | Value Type            | 0                                                                             |
| struct        | 值类型，复合类型       | Struct是复合类型，go语言会自动递归地初始化每一个元素为其类型对应的零值。            |
| bool          | 值类型                | false                                                                         |
| float         | 值类型                | 0.0                                                                            |
| string        | 值类型                | ""                                                                             |
| pointer       | 引用类型              | nil                                                                            |
| function      |                       | nil                                                                            |
| interface     |                       | nil                                                                            |
| slice         | 引用类型               | nil                                                                            |
| channel       | 引用类型               | nil                                                                            |
| map           | 引用类型               | nil                                                                            |
| array         | 值类型, 复合类型       | Array是复合类型，go语言会自动递归地初始化每一个元素为其类型对应的零值                |





# Go type
```go
type 类型名字 底层类型
```


```go
type Celsium float64        // 摄氏温度
type Fahrenheit fload64     // 华氏温度

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }


var c Celsius
var f Fahrenheit
fmt.Println(c == 0)          // "true"
fmt.Println(f >= 0)          // "true"
fmt.Println(c == f)          // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!
```




