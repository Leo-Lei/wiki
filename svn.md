---
layout: post
title: Base useage of SVN
date: 2015-12-16 11:00:00
tags:
- SVN
categories: VCS
description: The tutoria will describe the useage of SVN.
---


# CentOS 安装SVN
```bash
yum install subversion
```
> 截止到我写这个文档的时候，通过yum安装的svn版本并不是最新的，但使用上并没出现什么问题。

# 新建版本库
### 新建一个目录，以保存svn仓库的文件
```bash
mkdir -p /opt/svn/repos/sample
```
### 使用`create`新建一个svn仓库
```bash
svnadmin create /opt/svn/repos/sample
```
create命令会将`/opt/svn/repos/sample`目录初始化为一个svn仓库，会在该目录下自动生成一些文件
```bash
ls -lh

drwxr-xr-x 2 root root 4.0K 6月  15 13:30 conf
drwxr-sr-x 6 root root 4.0K 6月  15 11:14 db
-r--r--r-- 1 root root    2 6月  15 11:14 format
drwxr-xr-x 2 root root 4.0K 6月  15 11:14 hooks
drwxr-xr-x 2 root root 4.0K 6月  15 11:14 locks
-rw-r--r-- 1 root root  229 6月  15 11:14 README.txt
```
### 配置`conf/svnserve.conf`文件
```bash
# vi /opt/svn/repos/sample/conf/svnserve.conf   
[general]   
anon-access = none   
auth-access = write   
password-db = passwd   
authz-db = authz   
realm = My Sample Repository #这是个提示信息  
```
### 配置`conf/passwd`
```bash
# vi /opt/svn/repos/sample/conf/passwd   
[users]   
user1 = 123456
user2 = 888888
```
### 编辑`conf/authz`文件
```bash
# vi /svn/project/conf/authz   
[groups]   
developer = user1,user2   
readonly = test2   
[/]   
@developer = rw   
@readonly = r   
* =  
```
### 重启svn服务
```bash
svnserv -d -r /opt/svn/repos/sample
```
### 访问svn服务
```bash
svn://192.168.11.101
```

# SVN仓库迁移
### 将准备迁移的仓库导出
```bash
svnadmin dump /opt/svn/my_repo > /opt/svn_dump
```
### 在新的服务器上创建目录，以保存仓库
```bash
svnadmin create /opt/svn/my_repo
```
### 在新服务器上导入dump文件
```bash
svnadmin load /opt/svn/my_repo < /opt/svn_dump
```

