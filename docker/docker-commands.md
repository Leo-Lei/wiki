---
layout: post
title: Docker Commands
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---


|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker run mysql`                                                        | 运行mysql镜像，tag是latest                     | 
| `docker run mysql:tag`                                                    | 运行mysql镜像,并指定tag                        | 
| `docker run -d mysql`                                                     | 后台运行mysql镜像                              | 
| `docker run -i -t container_id /bin/bash`                                 | 运行容器并打开一个交互终端                       |
| `docker run -it container_id /bin/bash`                                   | 运行容器并打开一个交互终端.-i -t 和-it等效的      |
| `docker run --name my_container /bin/bash`                                | 运行容器并为容器取名为my_container               |
| `docker run -p 9000:8080 tomcat`                                          | 运行镜像，并将容器的8080端口映射到本机的9000端口    | 
| `docker run -v /home/doc:/opt/doc centos /bin/bash`                       | 将主机目录/home/doc挂载到容器/opt/doc目录        | 
| `docker run -it --entrypoint=/bin/bash container_id`                      | 启动容器，覆盖容器的entrypoint                  |
| `docker run --rm java`                                                    | `--rm`:容器停止后自动删除容器                  |
| `docker run --net=bridge ubuntu`                                          | 使用桥接网络模式                                |
| `docker run --net=host ubuntu`                                            | 使用host网络模式                               |
| `docker ps`                                                               | 查看容器(正在运行的)                            | 
| `docker ps -a`                                                            | 查看容器(包括停止的，正在运行的                   | 
| `docker images`                                                           | 查看镜像                                      | 
| `docker inspect container_id`                                             | 查看容器详情                                   | 
| `docker stop container_id`                                                | 停止容器                                      | 
| `docker start container_id`                                               | 启动容器                                      | 
| `docker rm container_id`                                                  | 删除容器                                      | 
| `docker rmi image_id`                                                     | 删除镜像                                      | 
| `docker commit -m "this is comment" -a "tom" container_id tom/hexo:v2`    | 将容器里的改动提交到镜像中                      | 
| `docker tag image_id leolei2094/hexo:v2`                                  | 给镜像取一个tag                                | 
| `docker cp container:/opt/doc /home/doc`                                  | 拷贝文件                                      |
| `docker cp /home/doc container:/opt/doc`                                  | 拷贝文件                                      |
| `docker attach mysql`                                                     | 附加到容器,使用`Ctrl+P`,`Ctrl+Q`退出，不让会关闭容器| 
| `docker exec -it mysql:5.7.16 apt-get install -y vim`                     | 进入容器执行命令                               |
| `docker exec -it my_container /bin/bash`                                  | 进入容器                                      |
| `docker exec some-mysql sh -c 'mysql -uroot < /opt/hello.sql'`            | 执行容器内命令,适用于命令中包含特殊字符，比如`<`    |
| `docker push leolei2094/mysql:latest`                                     | 上传镜像到Docker Hub                           |
| `docker logs container_id`                                                | 查看容器的日志                                 |
