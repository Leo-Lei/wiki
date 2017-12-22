---
layout: post
title: Kubernetes Volume
date: 2017-07-08 11:20:00
tags:
- docker
categories: Java
---

# Kubernetes Volume

* Kubernetes的Volume定义在Pod上。


# Pod中声明Volume
```yaml
template:
  metadata:
    labels:
      app: app-demo
      tier: frontend
  spec:
    volumes:
    - name: datavol
      emptyDir: {}
    containers:
    - name: tomcat-demo
      image: tomcat
      volumeMounts:
        - mountPath: /mydata-data
          name: datavol
      imagePullPolicy: IfNotPresent
```

* mounntPath: /mydata-data:将datavol的volume挂载到容器的`/mydata-data`目录

# Kubernetes的Volume类型
### emptyDir
在Pod被分配到Node时创建的，初始内容为空，也不用指定宿主机器上的文件或目录，这是Kubernetes自动分配的一个目录，当Pod从Node上移除时，emptyDir中的数据会被删除。    
所以emptyDir的一些用途：    
1. 临时空间，无须永久保留
2. 一个容器需要从另一个容器中获取数据的目录，用于多容器共享目录    
### HostPath
宿主机器上的文件或目录。可以用于保存数据，比如日志文件等。
```yaml
volumes:
- name: "log-storage"
  hostPath:
    path: "/data"
```
### gcePersistentDisk
google的公有云提供的文件存储服务。
### awsElasticBlockStore
AWS的公有云提供的文件存储服务

