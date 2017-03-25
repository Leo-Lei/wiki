---
layout: post
title: Class in Python
date: 2017-03-25 14:40:00
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---

# 构造函数和析构函数

```python
class test(object):
    def __init__(self):
        print 'AAAA'
    def __del__(self):
        print 'BBBB'
    def my(self):
        print 'CCCC'
        
>>> obj = test()
AAAA
>>>> obj.my()
CCCC
>>>> del obj
BBBB
```
