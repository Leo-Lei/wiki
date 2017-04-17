---
layout: post
title: Docker
date: 2017-03-13 12:10:00
tags:
- docker
categories: Java
description: docker
---

|               command                |                              |
| ------------------------------------ | ---------------------------- |
| `cat /etc/redhat-release`            | 显示系统版本                   |


|             command          |               desc             |
| ---------------------------- | ------------------------------ |
| `yum install gcc`            | 安装gcc                        |
| `yum install -y gcc`         | 安装gcc,默认yes                 |
| `yum list`                   | 列出所有可安装的软件包            |
| `yum list installed`         | 列出所有已安装的软件包            |
| `yum info installed`         | 列出所有已安装的软件包信息         |
| `yum list | grep gcc`        | 搜索可用的gcc包                 |
| `yum info gcc`               | 显示安装包gcc的信息              |
| `yum makecache`              | 更新yum缓存                     |

# CentOS7 修改hostname

### 方法1
```bash
hostname <host_name>
```
这种方式，只能修改临时的主机名，当重启机器后，主机名称又变回来了。

### 方法二
```bash
hostnamectl set-hostname <hostname>
```
使用这种方式修改，可以永久性的修改主机名称！


# yum安装openjdk
### 安装JRE:
```bash
yum install java-1.8.0-openjdk.x86_64
```
### 安装JDK:
```bash
yum install java-1.8.0-openjdk-devel.x86_64
```



# Systemd服务
Systemd的服务的配置文件目录为：`/etc/systemd/system`    

|                                         |                                  |
| --------------------------------------- | -------------------------------- |
| `systemctl daemon-reload`               | 刷新systemd服务的配置              |
| `systemctl start <service_name>`        |                                  |
| `systemctl stop <service_name>`         |                                  |
| `systemctl status <service_name>`       |                                  |


```bash
[Unit]
Description=score-rank
After=syslog.target

[Service]
User=root
ExecStart=/usr/java/jdk1.8.0_92/bin/java -jar /root/leiwei/service-1.0.jar
SuccessExitStatus=143

[Install]
WantedBy=multi-user.target
```

[http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-commands.html)

