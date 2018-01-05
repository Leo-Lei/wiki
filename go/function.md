---
layout: post
title: Go function
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 为类型添加方法

```go
type Celsius float64                   // 定义了一个类型
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }          // 为Celsius添加一个String方法
```





