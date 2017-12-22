---
layout: post
title: Linux用户
date: 2017-03-21 15:10:00
tags:
- Linux
categories: Linux
---



# 查看所有用户
```bash
cat /etc/passwd
```
输出结果如下图所示        
```bash
$ cat /etc/passwd
root:x:0:0:root:/root:/bin/bash
jenkins:x:498:499:Jenkins Continuous Integration Server:/var/lib/jenkins:/bin/bash
apache:x:48:48:Apache:/var/www:/sbin/nologin
mysql:x:27:27:MySQL Server:/var/lib/mysql:/bin/bash
jira:x:500:500:Atlassian JIRA:/home/jira:/bin/bash
```

# 创建用户
```bash
useradd [选项] [用户名]
```

* `-d`:指定用户所属主目录。如果此目录不存在，如果同时使用-m选项，可以创建主目录。
* `-g`:指定用户所属的用户组
* `-G`:指定用户所属的附加组。一个用户只能有一个主组，但是可以有多个附加组
* `-s`:指定用户的登陆shell

```bash
useradd -d /usr/admin -m admin
```
此命令创建了一个用户admin，主目录是/usr/admin。

```bash
useradd -s /bin/sh -g mygroup –G adm,root gem
```
此命令创建了一个用户gem，shell是/bin/sh,主组是mygroup,附加组是adm和root组。


# 删除用户
```bash
userdel [用户名]
```

# 修改用户
```bash
usermod [选项] [用户名]
```

```bash
usermod -s /bin/ksh -d /home/z –g developer sam
```
此命令将用户sam的登录Shell修改为ksh，主目录改为/home/z，用户组改为developer。


# 修改用户口令
```bash
passwd [选项] 用户名
```

```bash
sudo passwd root
```


