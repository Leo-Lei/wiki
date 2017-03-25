---
layout: post
title: Python
date: 2016-08-05 16:30:17
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---



# 让Python文件可执行
```python
#! /usr/bin/python
```

# Python支持utf-8
```python
# -*- coding: utf-8 -*-

print '你好'
```





# Python的`if __name__=='__main__'`
假设有如下的文件`hello.py`:
```python
print 'hello world'
```
如果在`app.py`文件中导入`hello.py`:    
`app.py`文件内容如下:
```python
import hello.py
```
会导致`hello.py`中的`print 'hello world'`被执行，很多时候这不是我们想要的，所以，一般在python文件中我们只会定义一些变量和方法，并且有一个类似于main函数的入口，这样就会避免在python文件被导入的时候，执行了某些代码。如果需要执行main函数，可以用下面的语法：
```python
if __name__=='__main__':
    hello()
```
比如完整的Python文件:
```python
import sys

def hello():
    print 'hello world'
    
if __name__=='__main__':
    hello()
```

# Python的string.format()
```python
'{0},{1}'.format('kzc',18)                     # 输出'kzc,18'  
'{},{}'.format('kzc',18)                       # 输出'kzc,18'  
'{1},{0},{1}'.format('kzc',18)                 # 输出'18,kzc,18'
'{name},{age}'.format(age=18,name='kzc')       # 输出'kzc,18'
'This guy is {person.name},is {person.age} old'.format(person=john)   # 输出 'This guy is john,is 26 old' 
```

# Python获取命令行参数
```python
import sys

print "script file name:", sys.argv[0]

for i in range(1, len(sys.argv)):
    print "arg", i, sys.argv[i]
```
执行该脚本
```bash
test.py hello world
```
输出
```text
script file name: test.py
arg 1 hello
arg 2 world
```
