---
layout: post
title: Go Command tool
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 通过os库获取命令行参数
```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    fmt.Println(os.Args)
}
```

编译后执行
```bash
./cmd -user="tom"
[./cmd -user=tom]
```

