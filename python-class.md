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



# 静态方法和类方法
```python

class TestClassMethod(object):

    METHOD = 'method hoho'

    def __init__(self):
        self.name = 'leon'

    def test1(self):
        print 'test1'
        print self

    @classmethod
    def test2(cls):
        print cls
        print 'test2'
        print TestClassMethod.METHOD
        print '----------------'

    @staticmethod
    def test3():
        print TestClassMethod.METHOD
        print 'test3'

if __name__ == '__main__':
    a = TestClassMethod()
    a.test1()
    a.test2()
    a.test3()
    TestClassMethod.test3()

```

* test1为实例方法，第一个参数是self，实例本身                
* test2为类方法，第一个参数是cls，类本身         
* test3为静态方法，可以不接受参数          


|             |    是否可以访问静态变量(类变量)     |    是否可以访问实例变量    |    是否可以使用Class.method()     |
| ----------- | ------------------------------- | ---------------------- | ------------------------------- |
|   实例方法   |       no                        |     yes                 |  no                             |
|   类方法     |       yes                       |     no                 |  no                             |
|   静态方法    |       yes                       |     no                 |  yes                             |

