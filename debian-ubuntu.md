---
layout: post
title: Debian / Ubuntu
date: 2017-03-12 11:10:00
tags:
- docker
categories: Java
description: docker
---

# Dockerfile command
指令忽略大小写，但建议使用大写。

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `FROM ubuntu:3.0`                                                         | 基于ubuntu:3.0创建镜像                         | 
| `MAINTAINER leiwei <leiwei2094@gmail.com>`                                | 指定维护者                                     | 
| `ENV REFRESHED_AT 2016-01-24`                                             | 设置环境变量                                   |  
| `RUN curl http://mycompany.com/archive/2.6.30.zip`                        | 执行命令                                      |
| `RUN unzip 2.6.30.zip`                                                    | 执行命令                                      |
| `RUN cd disconf-2.6.30/disconf-web`                                       | 执行命令                                      |
| `VOLUME /home/data`                                                       | 创建数据卷，用于和别的容器数据共享                 |
| `WORKDIR disconf-2.6.30/disconf-web`                                      | 设置当前工作路径                                |
| `CMD ["sh","deploy/deploy.sh"]`                                           | 容器启动后执行命令                              |
| `COPY hello.txt /root/data/hello.txt`                                     | 复制本地主机文件到容器中                         |


### FROM
```
FROM <image>
```
或
```
FROM <image>:<tag>
```
