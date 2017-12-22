---
layout: post
title: How to Use PIP
date: 2016-01-15 14:00:00
tags:
- Python
categories: Python
---

# PIP命令

|                    Command                              |                 Desc                |
| ------------------------------------------------------- | ----------------------------------- |
| `pip install Flask`                                     | 安装包                               |
| `pip search yaml`                                       | 搜索包                               |
| `pip install Flask --proxy="http://127.0.0.1:8080"`     | 配置代理                             | 

> 注意：执行`pip install`命令时，可能会出现没有权限的错误，这时候就需要使用`sudo pip install`了。    

# pip安装的第三方包的安装路径
安装路径是在`site-packages`中的。比如`C:\\Python27\\lib\\site-packages`.    
```bash
$ python
>>> import site;
>>> site.getsitepackages()
['C:\\Python27', 'C:\\Python27\\lib\\site-packages']
>>> exit()
```

# pip离线安装包
大致步骤如下:
1. 创建一个目录作为pip repository，来保存第三方的包
2. 使用PIP下载包到上一步中创建的pip repository。这时候需要机器是联网的。
3. 将整个pip repository目录都拷贝到无法联网的机器上
4. 使用pip命令来安装，并指定使用本地的pip repository，而不是远程的官方的pip仓库

第一步：创建一个本地的PIP repository:
```bash    
$ mkdir C:\foo\my-pip-repository
```
第二步：下载第三方包到本地的PIP repository：
```bash   
$ pip install --download C:\foo\my-pip-repository Flask --proxy="http://127.0.0.1:8080"
```
第三步：指定本地的pip repository，安装第三方包:
```bash    
$ pip install --no-index --find-links=C:\foo\my-pip-repository Flask
```

