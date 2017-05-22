---
layout: post
title: Linux Crontab
date: 2017-03-21 15:10:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the useage of Linux.
---

# 一个简单的Demo
需求:创建一个定时任务，每分钟向`/opt/hello.log`文件中写入一行"hello world"    
1. 新建文件`/opt/hello.log`
```bash
touch /opt/hello.log
```
2. 新建一个shell文件`/opt/hello.sh`
```bash
#! /bin/bash
echo `date` : hello >> /opt/hello.log
```
3. 创建一个crontab任务
```bash
crontab -e
```
加入一个定时任务
```
*/1 * * * * /opt/hello.sh
```
4. 重启crontab
```bash
service crond restart
```
# demo
```bash
# 每天晚上1点30分执行
30 1 * * * /opt/delete_history_logs.py
```
