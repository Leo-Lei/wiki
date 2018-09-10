---
layout: post
title: Docker
date: 2017-07-01 12:10:00
tags:
- docker
categories: Java
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
| `cat /etc/redhat-release`    | 查看redhat,centOS版本            |



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
yum install java-1.8.0-openjdk-debug.x86_64
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



# yum epel
如果yum官方的源里找不到我们需要的包，可以去该软件的官方文档查看，是否官方已提供repo，如果官方没有repo，我们还可以使用第三方的yum源。比如epel。
```bash
yum install epel-release
```


# 配置环境变量
方法一：直接运行命令export PATH=$PATH:/usr/local/webserver/php/bin 和 export PATH=$PATH:/usr/local/webserver/mysql/bin

使用这种方法，只会对当前会话有效，也就是说每当登出或注销系统以后，PATH 设置就会失效，只是临时生效。

方法二：执行vi ~/.bash_profile修改文件中PATH一行，将/usr/local/webserver/php/bin 和 /usr/local/webserver/mysql/bin 加入到PATH=$PATH:$HOME/bin一行之后

这种方法只对当前登录用户生效

方法三：修改`/etc/profile`文件使其永久性生效，并对所有系统用户生效，在文件末尾加上如下两行代码
```bash
PATH=$PATH:/usr/local/webserver/php/bin:/usr/local/webserver/mysql/bin
export PATH
```
最后：执行 命令source /etc/profile或 执行点命令 ./profile使其修改生效，执行完可通过echo $PATH命令查看是否添加成功。
