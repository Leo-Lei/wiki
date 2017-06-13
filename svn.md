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


# SVN仓库迁移
### 将准备迁移的仓库导出
```bash
svnadmin dump my_repo > /opt/svn_dump
```
### 在新的服务器上创建目录，以保存仓库
```bash
svnadmin create /opt/svn
```
### 导入dump文件
```bash
svnadmin load /opt/svn < /opt/svn_dump
```
### 更改conf文件    
将迁移的仓库的conf文件复制到新的服务器
