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






|                  |       Value Type                             |                 Reference Type                  |
| :--------------: | :------------------------------------------: | :---------------------------------------------: |
| `a = b`          | 创建b的一个副本，给a。修改b**不会**影响a        | a和b指向内存中同一个数据。修改b**会**影响a         |




|                 Func                 |                  result                      |               Remark               |
| ------------------------------------ | -------------------------------------------- | ---------------------------------- |
| `func update(m map[int]string)`      | 指针传递。同一个map                           |                                    |
| `func update(s Student)`             | 值传递。方法接收一个副本。两个Student          |                                    |
| `func update(s *Student)`            | 指针传递。同一个Student                       |                                    |



