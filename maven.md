---
layout: post
title: Base useage of Maven
date: 2015-06-26 15:30:00
tags:
- Maven
categories: Java
---


# Install component and all dependencied from command line
```bash
mvn org.apache.maven.plugins:maven-dependency-plugin:2.4:get -DartifactId=jersey-servlet -DgroupId=com.sun.jersey -Dversion=1.19
```

# Maven commands

| command                                                 |             Description                                       |                   
| ------------------------------------------------------- | ------------------------------------------------------------- |
| `mvn clean`                                             | Clean the *target* folder.                                    |
| `mvn clean install`                                     | Install artifact(jar or war) to local maven repository.       |
| `mvn dependency:tree -Dverbose -Dincludes=asm:asm`      | 检查maven依赖树                                                 |
| `mvn install -Dmaven.test.skip=true`                    | 跳过单元测试                                                    |

# 4. Maven dependency
In the pom file, you can specify the dependency.
## 4.1 The dependency coordinate
* groupId：org.springframework
* artifactId: spring-core
* version: 2.5.6
* packing: default value is jar.
* classifier: additional property.

## 4.2 The dependency scope
* compile： default value.
* test:
* provided: Is available for compiling and testing, but unavailable for running. Maven will not add it to war while packing the web project.
For example, the servlet-api is *provided* scope.
* runtime: Is available for test and running, but unavailable for compiling main class. For example, the JDBC Driver implementation is *runtime* scope.
* system:

> Note: Maven use different *class path* while compiling main code and test code.

| scope        |   compiling?  |  testing?   | running  |      example                                 |
| ------------ | ------------- | ----------- | -------- | -------------------------------------------- |
| compile      | Y             | Y           | Y        |   spring-core                                |
| test         | --            | Y           | --       |   JUnit                                      |
| provided     | Y             | Y           | --       |   servlet-api                                |
| runtime      | --            | Y           | Y        |   JDBC driver implementation                 |
| system       | Y             | Y           | --       |   Local jar files. Not in maven repository.  |
