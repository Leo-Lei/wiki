---
layout: post
title: Mongo
date: 2017-04-11 13:05:00
tags:
- Java
categories: Java
description: web-authentication
---


# yum安装mongo
```bash
/etc/yum.repos.d/mongodb-org-3.4.repo
```

```bash
[mongodb-org-3.4]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/3.4/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-3.4.asc
```

```bash
yum install mongodb-org            # 安装所有mongo的组件
yum install mongodb-org-shell      # 安装mongo客户端
```

# 连接mongodb
```bash
mongo --host mongo-server:3717 -u root -p
```


# mongo命令
|                command                  |                                             |
| --------------------------------------- | ------------------------------------------- |
| `show dbs`                              | 显示数据库列表                                 |
| `use <db_name>`                         | 切换数据库                                    |
| `show collections`                      | 显示集合，类似关系型数据库的表                    |


