---
layout: post
title: Docker Commands
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---


# docker容器命令

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker run mysql`                                                        | 运行mysql镜像，tag是latest                     | 
| `docker attach mysql`                                                     | 附加到容器,使用`Ctrl+P`,`Ctrl+Q`退出，不让会关闭容器| 
| `docker run mysql:tag`                                                    | 运行mysql镜像,并指定tag                        | 
| `docker run -d mysql`                                                     | 后台运行mysql镜像                              | 
| `docker run -i -t container_id /bin/bash`                                 | 运行容器并打开一个交互终端                       |
| `docker run -it container_id /bin/bash`                                   | 运行容器并打开一个交互终端.-i -t 和-it等效的      |
| `docker run --name my_container /bin/bash`                                | 运行容器并为容器取名为my_container               |
| `docker ps`                                                               | 查看容器(正在运行的)                            | 
| `docker ps -a`                                                            | 查看容器(包括停止的，正在运行的                   | 
| `docker images`                                                           | 查看镜像                                      | 
| `docker inspect container_id`                                             | 查看容器详情                                   | 
| `docker stop container_id`                                                | 停止容器                                      | 
| `docker start container_id`                                               | 启动容器                                      | 
| `docker rm container_id`                                                  | 删除容器                                      | 
| `docker rmi image_id`                                                     | 删除镜像                                      | 
| `docker run -p 9000:8080 tomcat`                                          | 运行镜像，并将容器的8080端口映射到本机的9000端口    | 
| `docker commit -m "this is comment" -a "leolei" container_id leolei2094/hexo:v2` | 将容器里的改动提交到镜像中                 | 
| `docker tag image_id leolei2094/hexo:v2`                                  | 给镜像取一个tag                                | 
| `docker run -v /home/doc:/opt/doc centos /bin/bash`                       | 将主机目录/home/doc挂载到容器/opt/doc目录        | 
| `docker cp container:/opt/doc /home/doc`                                  | 拷贝文件                                      |
| `docker cp /home/doc container:/opt/doc`                                  | 拷贝文件                                      |
| `docker exec -it mysql:5.7.16 apt-get install -y vim`                     | 进入容器执行命令                               |
| `docker exec -it my_container /bin/bash`                                  | 进入容器                                      |
| `docker exec some-mysql sh -c 'mysql -uroot < /opt/hello.sql'`            | 执行容器内命令,适用于命令中包含特殊字符，比如`<`    |
| `docker run -it --entrypoint=/bin/bash container_id`                      | 启动容器，覆盖容器的entrypoint                  |
| `docker run --rm java`                                                    | `--rm`:容器停止后自动删除容器                 |

# Dockerfile command

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `FROM ubuntu:3.0`                                                         | 基于ubuntu:3.0创建镜像                         | 
| `MAINTAINER leiwei <leiwei2094@gmail.com>`                                | 指定维护者                                     | 
| `ENV REFRESHED_AT 2016-01-24`                                             | 设置环境变量                                   |  
| `RUN curl http://mycompany.com/archive/2.6.30.zip`                        | 执行命令                                      |
| `RUN unzip 2.6.30.zip`                                                    | 执行命令                                      |
| `RUN cd disconf-2.6.30/disconf-web`                                       | 执行命令                                      |
| `VOLUME /home/data`                                                       | 创建数据卷，用于和别的容器数据共享                 |
| `WORKDIR disconf-2.6.30/disconf-web`                                      | 设置当前工作路径                                |
| `CMD ["sh","deploy/deploy.sh"]`                                           | 容器启动后执行命令                              |
| `COPY hello.txt /root/data/hello.txt`                                     | 复制本地主机文件到容器中                         |


# Docker网络
### Bridge
```bash
docker run --net=bridge ubuntu
```
Bridge模式是默认的。Bridge模式下，每个容器都有自己的网络，IP是`172.17.0.X`。
### Host
```bash
docker run --net=host ubuntu
```

# Docker删除镜像
```bash
docker rmi image-name:tag
```

```bash
docker rmi -f image-name:tag
```

```bash
docker tag image-id image-name:tag
docker rmi -f image-name:tag
```

# docker hub相关

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker push leolei2094/mysql:latest`                                     | 上传镜像到Docker Hub                           |


# docker日志命令

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker logs container_id`                                                | 查看容器的日志                                 |
