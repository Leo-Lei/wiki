---
layout: post
title: Gradle Bugs
date: 2016-09-19 17:10:00
tags:
- Gradle
categories: 
- Java
- Gradle
---


# 在IntellijIdea中使用mavenLocal
mavenLoca，会让Gradle从maven的本地仓库去寻找组件。      
mavenLocal只会从$USER_HOME/.m2/setting.xml文件中去读取maven的配置。如果你设置了M2_HOME的环境变量的。
