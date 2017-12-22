---
layout: post
title: Docker Compose
date: 2017-03-25 11:10:00
tags:
- docker
categories: Java
---


# Docker compose命令

|               command              |                                    |
| ---------------------------------- | ---------------------------------- |
| `docker-compose up -d`             | 后台启动docker compose              |
| `docker-compose logs`              | 查看日志                            |
| `docker-compose port`              | 打印绑定的开放端口                    |
| `docker-compose ps`                | 显示容器                            |
| `docker-compose run`               | 运行一个一次性命令                    |
| `docker-compose build`             | 构建或重建服务                       |
| `docker-compoe help `              | 命令帮助                            |
| `docker-compose kill`              | 杀掉容器                            |
| `docker-compose pull`              | 拉取镜像服务                         |
| `docker-compose restart`           | 重启服务                            |
| `docker-compose rm`                | 删除服务                            |
| `docker-compose scale`             | 设置服务的容器数目                    |
| `docker-compose start`             | 开启服务                            |
| `docker-compose stop`              | 停止服务                            |


# Docker compose
安装docker-compose:            
https://github.com/docker/compose/releases    

```bash
curl -L https://github.com/docker/compose/releases/download/1.12.0-rc1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```
运行docker-compose:
```bash
docker-compose up 
```
-f   指定docker-compose.xml文件，默认是 docker-compose.xml  ，  当一条命令有多个-f参数时，会做替换操作
-p  指定docker-compose的项目目录，也就是docker-compose.xml文件的存储目录

```bash
docker-compose -f opt/disconf-compose/docker-compose.yml up
```


