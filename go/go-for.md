---
layout: post
title: Go中的for
date: 2017-03-08 15:30:00
tags:
- go
categories: go
---

# Go死循环
```go
fund main(){
    for{
    }
}
```
如果省略了循环条件，循环就不会结束，达到死循环的效果。也可以在死循环中使用break退出循环。
