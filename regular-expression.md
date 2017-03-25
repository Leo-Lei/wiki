---
layout: post
title: Base useage of Regular Expression
date: 2015-08-25 17:30:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the useage of Linux.
---

# 1. Metacharacter

| Code    |             Description                       |      Syntax                            |
| ------- | --------------------------------------------- | -------------------------------------- |
| `.`     | Any character except for new-line             |                                        |
| `\s`    | Any blank character                           |                                        |
| `\d`    | Number. i.e. 0,1,2,3,4,5,6,7,8,9              | `\d{3}` match 123                      |
| `^`     | Begining of a string                          |                                        |
| `$`     | End of a string                               |                                        |

# 2. Character Escape
If you want to search the metacharacter itself, you should use the `\` to escape the character. For example, use `\.` and `\*` to search `.` and `*` in text. Use `\\` to search `\` itself.        
For example, google\.com match "google.com"







# 3. Some examples
下面这个例子同时使用了这两种断言：(?<=\s)\d+(?=\s)匹配以空白符间隔的数字(再次强调，不包括这些空白符)。





