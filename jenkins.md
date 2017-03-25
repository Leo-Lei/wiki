---
layout: post
title: Jenkins
date: 2016-07-15 15:40:00
tags:
- Java
categories: Java
description: Jenkins.
---

# Overview               
Jenkins is a continus integration tool.               
The official website is [https://jenkins.io](https://jenkins.io).                
You can fidn more documents on:[https://wiki.jenkins-ci.org/display/JENKINS/Home](https://wiki.jenkins-ci.org/display/JENKINS/Home).

# Install and access Jenkins    
The easiest way to execute Jenkins is through the built in Jetty servlet container. You can execute Jenkins like this:
```bash
$ java -jar jenkins.war
```
To see Jenkins, simply bring up a web browser and go to URL http://myServer:8080 where myServer is the name of the system running Jenkins.        

| Parameter                       |            Description                  |
| ------------------------------- | --------------------------------------- |
| --httpPort=$HTTP_PORT           | The default is port 8080.               |

Jenkins need some disk space to perform builds and keep archives. The default directory is `~/.jenkins`. But you can set "JENKINS_HOME" environment variable to the new home directory before launching the servlet container.

# Install Jenkins plugin manually      
Yes you can. Download the Plugin (*.hpi File) and put it in the following directory:    
`<jenkinsHome>/plugins/`                
Afterward you will need to restart Jenkins.



