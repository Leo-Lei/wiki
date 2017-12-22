---
layout: post
title: Docker MySql
date: 2016-08-16 14:25:00
tags:
- Gradle
categories: 
- Java
---

# 在Docker中运行Mysql
1. 从Docker Hub拉取mysql镜像
:`docker pull mysql`    
2. 运行mysql镜像:`docker run -it -d -p 3307:3306 mysql /bin/bash`。    
`-it`:创建一个终端，以供后面从主机ssh到mysql容器。    
`-d`:让myslq容器在后台运行。    
`-p 3307:3306`:将容器的3306端口绑定到本机的3307端口。这样就可以通过本机的3307端口来访问容器中的mysql。    
3. attach到mysql容器:`docker attach container_id`。会ssh到mysql容器。            
4. 启动mysql
容器中mysql默认是不启动的，启动它:`> service mysql start`。一旦容器中的mysql启动后，就可以在本机上通过mysql client来连接到容器中的mysql。
5. 创建mysql用户，用于远程连接
`create user 'admin'@'%' identified by 'admin'`    
`grant all on *.* to 'admin'@'%'`    
6. 在本机连接容器中的mysql
在本地通过mysql client连接到容器中的mysql:`mysql -h host_ip -P 3307 -uadmin -padmin`。注意这里不能使用localhost，或127.0.0.1，要使用本地的真实ip。Mac可通过`ifconfig`,Windows可用`ipconfig`,Linux可用`ip addr`来查看ip。    
7. 创建数据库或表
一旦连接上来mysql，就可以创建数据库或表了。    

使用JDBC来连接容器中的mysql:
只需要通过本机ip和端口3307就可以通过JDBC连接到容器中的mysql。
`jdbc:mysql://192.168.4.14:3307/my_db`

## Docker中mysql数据的持久化    
注意：在mysql容器中创建的数据库或表等数据，在容器被删掉之后，数据是会丢失的。将这个容器commit到镜像，mysql的数据也会丢失。可以在启动mysql容器的时候，添加-v来挂载mysql数据文件，将容器中的mysql数据进行持久化。

在host上创建一个目录，用于存储mysql的数据，我们会将该目录挂载到mysql容器中，这样mysql容器中的数据就会写入到host的这个目录。同时，容器中的mysql也可以从该目录读取数据。
`mkdir -p root/docker.data/mysql`        
`docker run -it -d --name some-mysql -v /root/docker.data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql /bin/bash`来启动容器，其中：
`-it`:创建一个终端用于在host可以ssh到容器中        
`-d`:后台运行    
`--name`:给容器指定一个名字，不然docker会随机分配一个name    
`-v /root/docker.data/mysql:/var/lib/mysql`:将本地的`/root/docker.data/mysql`目录mount到容器的`/var/lib/mysql`目录。`/var/lib/mysql`是mysql存储数据的目录。



