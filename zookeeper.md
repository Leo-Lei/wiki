---
layout: post
title: Zookeeper
date: 2016-08-11 17:10:00
tags:
- Java
categories: Java
description: web-authentication
---

# Overview               

# Install Zookeeper          
Go to [Zookeeper website](https://zookeeper.apache.org/releases.html), and find the Download link. Then you can download the release package from download page.    
The download package is in the format of `.tar.gz`, you can unpackage it on Windows, Mac or Linux platform.     
**Configure Zookeeper**             
Go to the `<zookeeper/conf>` directory of Zookeeper, there is a `zoo_sample.cfg` file. Rename it to `zoo.cfg` and Zookeeper will read it as the default configuration file.        
Zookeeper can run as single node, or cluster.    
**1. Single node configuration**:    
```text
tickTime=2000
initLimit=10
syncLimit=5
dataDir=C:/zookeeper-data
clientPort=2181
```
**2. Cluster configuration**:    
To be added......      

**Run Zookeeper**    
1. Go to `<zookeeper>/bin` folder.    
2. Run `zkServer.sh start` or `zkServer.cmd` to run Zookeeper server.    
3. run `zkCli.sh -server 127.0.0.1:2181` or `zkCli.cmd` to run Zookeeper client.    

**Zookeeper CMD interface**    
After started the Zookeeper server, then you can open the Zookeeper client and execute command via the command-line interface, to create/delete/list data in Zookeeper.    
```bash
[zk: localhost:2181(CONNECTED) 0] ls /
[zookeeper]
[zk: localhost:2181(CONNECTED) 1] create /app app
Created /app
[zk: localhost:2181(CONNECTED) 2] ls /
[app, zookeeper]
[zk: localhost:2181(CONNECTED) 3] create /app/app1 app1
Created app/app1
[zk: localhost:2181(CONNECTED) 4] rmr /app
```

# Resources

[https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper](https://www.ibm.com/developerworks/cn/opensource/os-cn-zookeeper)

