---
layout: post
title: Gradle - The Best Automation Build Tool of 21 Century
date: 2015-06-28 18:58:17
tags:
- Gradle
categories: 
- Java
- Gradle
description: The tutoria will show you how to use Gradle to build your project.
---


|                    command                                       |   desp                           | 
| ---------------------------------------------------------------- | -------------------------------- | 
| `gradle -q webapp:dependencyInsight --dependency groovy`         | 查看某一个特定依赖                  | 
| `gradle build -x test`                                           | skip test                        |
| `gradle build -Pfoo=bar`                                         | 指定自定义参数                     |




# Gradle发布构件
```gradle
apply plugin: 'java'
apply plugin: 'maven'
group='com.zqf'
version='1.0.0'
repositories {
	//依赖maven远程仓库
	maven { url "http://nexus.zhaiqianfeng.com:8081/nexus/content/groups/public" }
}
dependencies {
    testCompile 'junit:junit:4.12'
}
//发布构件
uploadArchives{
	repositories{
		mavenDeployer{
			//发布到maven本地文件
			repository(url:"file://localhost/tmp/maven-rpo/")
			
			//发布到maven远程仓库
			repository(url: "http://nexus.zhaiqianfeng.com:8081/nexus/content/repositories/thirdparty/") {
			    authentication(userName: "admin", password: "admin123")
			}
		}	
		
		//发布到ivy本地文件
		ivy{
			url "/tmp/ivy-rpo/"
		}
		
		//发布到maven本地仓库
		mavenLocal()
		//更多的仓库.....
	}
}

```








# Reference
[https://docs.gradle.org/current/userguide/userguide.html](https://docs.gradle.org/current/userguide/userguide.html)










