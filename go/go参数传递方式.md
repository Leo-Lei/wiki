---
layout: post
title: Go中的值类型和引用类型
date: 2017-03-08 15:30:00
tags:
- go
categories: go
---


|      Type     |    Value Type   |    Reference Type  |
| :-----------: | :-------------: | :----------------: |
|      int      |       X         |                    |
|     array     |       X         |                    | 
|     Map       |                 |          X         |
|     slice     |                 |          X         |
|    channel    |                 |          X         |


|                 Func                 |                  result                      |               Remark               |
| ------------------------------------ | -------------------------------------------- | ---------------------------------- |
| `func update(m map[int]string)`      | 指针传递。同一个map                           |                                    |
| `func update(s Student)`             | 值传递。方法接收一个副本。两个Student          |                                    |
| `func update(s *Student)`            | 指针传递。同一个Student                       |                                    |
| `func update(i int)`                 | 值传递。两个int                              |                                    |
| `func update(i *int)`                | 指针传递。一个int                             |                                    |


