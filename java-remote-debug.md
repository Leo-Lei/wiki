---
layout: post
title: Java Remote Debugging
date: 2017-06-01 14:10:00
tags:
- docker
categories: Java
---

# 远程调试
```bash
-agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005
```

```bash
java -jar -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005 /opt/hello-world.jar
```

IDEA中新建Remote
