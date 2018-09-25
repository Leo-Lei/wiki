---
layout: post
title: Go struct
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# Go struct


# 空结构体
```go
type Q struct{}         // 定义一个类型Q，基于struct{}
var q struct{}          // 定义一个struct{}类型的变量q
s := struct{}{}         // 新建一个struct{}类型的变量
```

```go
type Student struct {
    id        int
    name      string
    sddress   string
    age       int
}
```



```go
var s *Student = new(Student)
s.id = 101
s.name = "tom"
s.address = "中关村"
s.age = 18
```

```go
var s1 Student = Student{103, "tom", "中关村", 20}
```

```go
var s1 *Student = &Student{102, "tom", "中关村", 20}
```

```go
var s1 *Student = &Student{id: 102, name: "tom", address: "中关村", age: 20}
```
