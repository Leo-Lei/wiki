---
layout: post
title: Go语言规范
date: 2017-03-08 15:30:00
tags:
- go
categories: go
---





# 注释
Go中有2种注释：
1. 单行注释。以`//`开始。
2. 多行注释。以`/*`开始，以`*/`结束。

# Go中的分号
大多数语言使用`;`来结束一条语句。但Go中大多数情况下，可以省略掉`;`。

# Go中的标识符
Go中的标识符可以给变量，类型等命名。一个标识符可以包括字母，数字和`_`，但必须是以字母开头。
下面的标识符都是合法的:
```go
a
_x9
ThisVariableIsExported
αβ
```

# 关键字
Go语言的关键字比较少，下面是全部的关键字：
```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
关键字不能作为标识符。

# 运算符
```go
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

# Integer字面常量
```go
42                  // 10进制
0600                // 8进制
0xBadFace           // 16进制
170141183460469231731687303715884105727
```

# 浮点常量
```go
72.40
2.718
6.67e-11
```
# 字符串常量
```go
“abc”
```

