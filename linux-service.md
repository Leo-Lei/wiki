---
layout: post
title: Dive into HashMap
date: 2016-06-22 17:10:00
tags:
- Java
categories: Java
description: HashMap.
---

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

