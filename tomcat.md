---
layout: post
title: Tomcat
date: 2017-05-11 14:20:00
tags:
- Gradle
categories: 
- Java
- Gradle
---


在docker中运行tomcat容器
```bash
docker run --name tomcat -it -d -p 8080:8080  \
-e JAVA_OPTS='-Xms800m -Xmx800m -Dlogs.dir=/opt/logs -Ddata.dir=/opt/data -Ddisconf.download.dir=/opt/data/disconf' \
-v /opt/app:/usr/local/tomcat/webapps \
-v /opt/logs:/opt/logs \
-v /opt/data:/opt/data \
tomcat
```

