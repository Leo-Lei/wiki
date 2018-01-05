---
layout: post
title: Go new method
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


```go
p := new(int)     //p, *int 类型，指向匿名的int变量
fmt.Println(*p)
*p = 2
fmt.Println(*p)
```


