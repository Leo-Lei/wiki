---
layout: post
title: A Good Python IDE PyCharm
date: 2015-08-18 14:20:00
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---

# 1. Install PyCharm
You can get the PyCharm distribution from the [official web site](https://www.jetbrains.com/pycharm/download/).

# 2. PyCharm KeyMap
This is my own PyCharm KeyMap configuration, maybe it is different from the default settings.

|         Key             |                                 Action                                                  |
| ----------------------- | --------------------------------------------------------------------------------------- |
| `Ctrl`+`Q`              | Quick Document: Show the document.                                                      |
| `Ctrl`+`Alt`+`F`        | Reformat the code.                                                                      |
| `Ctrl`+`B`              | Go to definition.                                                                       |

# 4. Pycharm base usage
## 4.1 Specify the Python Interpreter
Maybe multiple versions of Python are installed on your maching, you can swith the Python version on PyCharm by following steps:    
1. File -> Settings    
2. Project -> Project Interpreter    
3. Select the Project Interpreter you want    

## 4.2 Specify the Method argument type
As you are well aware, Python is an dynamic language, there is no type for the method argument in the method singnature. But this cause some inconvenience. Because there is no type, the IDE can not support intelli-sense help. But this can be fixed by adding doc-string for the method.

```python
def f(param):
   """
   :type param: MyClass
   :param param:
   :return:
   :rtype: int
   """

   print("Hello")
   return 10

```

You can get more detains from below link:
[Type Hinting in PyCharm](https://www.jetbrains.com/pycharm/help/type-hinting-in-pycharm.html)



# 添加python doc注释以支持pycharm的智能提示
```python
def f(param):
   """
   :type param: MyClass
   :param param:
   :return:
   :rtype: int
   """

   return 10
```

添加方法返回值类型:
```python
def f(param):
   """
   :rtype: list[int]
   """

   return 10
```
> 注意， 注释在方法签名的下面，和Java,C#等静态语言把注释放方法签名签名是不一样的。    

方法返回值类型是list:
```python
def f(param):
   """
   :rtype: list[User]
   """

   users = []
   users.append(User("tom"),User("jack"))
   return users 
```
