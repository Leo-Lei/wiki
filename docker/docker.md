---
layout: post
title: Docker
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---


# Docker command

docker hub相关

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker push leolei2094/mysql:latest`                                     | 上传镜像到Docker Hub                           |

docker日志命令

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker logs container_id`                                                | 查看容器的日志                                 |



docker容器命令

|                           command                                         |          usage                               | 
| ------------------------------------------------------------------------- | -------------------------------------------- | 
| `docker attach mysql`                                                     | 附加到容器,使用`Ctrl+P`,`Ctrl+Q`退出，不让会关闭容器| 
| `docker run mysql`                                                        | 运行mysql镜像，tag是latest                     | 
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

# CentOS安装Docker
1. 安装`yum-utils`，可以提供`yum-config-manager`命令：
```bash
sudo yum install -y yum-utils
```
2. 添加centos的docker repository
```bash
sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
```
3. 安装docker
```bash
sudo yum makecache fast
sudo yum install docker-ce
```
4. 启动Docker
```bash
sudo systemctl start docker
```
5. 检查Docker是否安装正确
```bash
docker run hello-world
```

# Mac安装Docker
[Docker](https://www.docker.com)

# Docker数据卷
## 将本地的一个目录挂载到容器中
命令为`docker run -v /foo/bar:/data ubuntu`。将主机上的`/foo/bar`目录挂载到容器中的`/data`目录。不管对容器中的/data操作还是主机上的`/foo/bar`进行操作，都是完全实时同步的。    
### 使用场景
假设有一个静态blog站点的镜像，比如hexo容器。本地有markdown格式的post文件。希望在容器启动的时候将本地/home/post目录挂载到容器的`/opt/blog/_posts`目录。这样就可以在容器中使用hexo来生成站点了。    
`docker run -it -d -v /home/post:/opt/blog/_posts hexo /bin/bash`    
登录到容器中可以看到容器中已经有`/opt/blog/_posts`目录了。和本地的文件夹内容一样。    

## 在容器里创建一个挂载点，供别的镜像共享使用    
命令`docker run -v /data ubuntu`。只设置了挂载点，当没指定关联的主机目录。这时docker会自动绑定主机上一个目录，通过`docker inspect`可以查看到。一般是在`/var/lib/docker/volumes/0ab0aaf0d6ef391cb68b72bd8c43216a8f8ae9205f0ae941ef16ebe32dc9fc01/_data`下。进入这个目录可以发现，该目录下并没有文件，甚至连这个目录都不存在，所以这种方式不是为了主机和容器间数据共享，而是为了容器于容器间数据共享。    
运行一个容器，该容器不运行应用，只存储数据。可以基于centos或者ubuntu等镜像来创建容器。         
`docker run -it -d ubuntu /bin/bash`          
向容器里面添加数据，可以直接在容器中创建，或通过`docker cp`从主机拷贝文件到容器中。假设添加文件到`/data`目录。        
`docker commit -m "add data" -a "leolei" container_id leolei2094/data-volume`         
启动容器volume容器            
`docker run -it --name data -v /data leolei2094/data-volume echo 'data volume only...'`                 
启动需要使用data-volume数据的镜像          
`docker run -it --volumes-from data centos /bin/bash`           
进入容器，可以发现有/data目录了，里面的文件内容和data-volume里的是一样的。    


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
Host模式下，官方文档介绍说容器的网络配置和宿主机器是一样的，IP也是一样的。经测试，大多数情况下是这样的，当也出现了，如果宿主机器的网络环境比较复杂，会出现容器的网络配置，比如IP，和宿主机器是不一样的。具体原因不清楚。
![Docker网络模型](http://wiki.jikexueyuan.com/project/docker-technology-and-combat/images/network.png)






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


