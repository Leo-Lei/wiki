---
layout: post
title: Python Module
date: 2017-03-17 10:30:17
tags:
- Python
categories: Python
---


# 导入自定义模块
### 和导入文件在同一目录
目录结构：
```text
hello.py
main.py
```
hello.py的内容：
```python
def say_hello():
    print 'hello world!'
```
在main中导入foo.py
```bash
import hello
hello.say_hello()
```
### 导入文件在目录中
目录结构：
```text
mycompany
    __init__.py
    hello.py
main.py
```
main中导入hello.py
```python
import mycompany.hello
mycompany.hello.say_hello() #一定要写全名称mycompany.hello.say_hello()。不能用hello.say_hello()
```
或者
```python
from mycompany import hello
hello.say_hello()
```
或者
```python
import mycompany.hello as hello
hello.say_hello()
```


# `site-package`中加入自定义模块
在Python的安装路径的`site-package`文件夹中，新建一个`XXX.pth`文件，在里面加上自定义模块所在的路径。    
假设自定义模块在这个路径下面:
```text
/opt/python-lib
        mycompany
            __init__.py
            utils
                __init__.py
                hello.py
```
hello.py的内容:
```python
def say_hello():
    print 'hello world'
```
在site-package目录添加`.pth`文件
```bash
cd /Library/Python/2.7/site-packages
touch my-python-lib.pth
```
将`/opt/python-lib`加入到`.pth`文件中
```text
/opt/python-lib
```
这样就可以直接导入`python-lib`中的自定义模块了
```python
import mycompany.utils.hello

mycompany.utils.hello.say_hello()
```
