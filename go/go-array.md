---
layout: post
title: Go Array
date: 2017-03-08 15:30:00
tags:
- Go
categories: Go
---

数组的一些特性:
* 内存上连续
* 由于内存上连续，访问速度快
* Go中数组是Value类型的

# 声明数组
```go
var array [5]int
```
```go
var array := [5]int{10,20,30,40,50}
```
```go
var array := [...]int{10,20,30,40,50} 
```

