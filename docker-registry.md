---
layout: post
title: Docker Registry
date: 2017-03-15 12:10:00
tags:
- docker
categories: Java
---


|              name            |                                     |
| ---------------------------- | ----------------------------------- |
| Docker Registry              | 开源的                               |
| Docker Hub                   | Docker官方的镜像仓库,免费              |
| Docker Trusted Registry      | 商用的Docker镜像仓库                  |                       


# Docker镜像命名规则
***registry.domain.com/mycom/imagename:latest***，其中：
1. registry.domain.com: 镜像所在服务地。如果没有指定，默认是dockerhub官方仓库的地址。
2. mycom：镜像命名空间，比如可以是项目名或用户名。
3. imagename：镜像名字
4. latest：版本号


# 部署Docker Registry
```bash
docker run --name docker-registry -d -p 5000:5000 -v /opt/registry/data:/var/lib/registry registry:2
```

# Push镜像到Registry

如果Registry没有做相应的配置，直接push镜像到registry会遇到错误**server gave HTTP response to HTTPS client**    
1. 新建文件`/etc/docker/daemon.json`
```bash
cd /etc/docker
touch daemon.json
```
>  网上很多文章说可以修改`/etc/sysconfig/docer`文件来解决，但这是旧版本的Docker会有这文件，新版本的Docker并无这文件。所以新版本的需要创建`/etc/docker/daemon.json`文件。
2. 添加如下配置
```json
{ "insecure-registries":["192.168.1.100:5000"] }
```
3. 重启docker
```bash
systemctl restart docker
```
4. push镜像：
```bash
docker pull ubuntu:latest
docker tag ubuntu:latest 192.168.1.100:5000/mycom/ubuntu:latest
docker push 192.168.1.100:5000/mycom/ubuntu:latest
```

# Registry API
|                      api                              |                                                   |
| ----------------------------------------------------- | ------------------------------------------------- |
| `curl 192.168.1.100:5000/v2/_catalog`                 |  输入格式为:{"repositories":["hello-world"]}        |
| `curl 192.168.1.100:5000/v2/<name>/tags/list`         |  name是仓库名,比如ubuntu，mysql等                    |





