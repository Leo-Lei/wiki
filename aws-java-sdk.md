---
layout: post
title: Amazon Java SDK
date: 2017-05-26 12:10:00
tags:
- docker
categories: Java
description: docker
---

# 将开发工具包和Gradle一起使用
1. 在`build.gradle`中添加依赖管理插件
```groovy
buildscript {
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath "io.spring.gradle:dependency-management-plugin:1.0.0.RC2"
    }
}

apply plugin: "io.spring.dependency-management"
```
2. 将BOM添加到dependencyManagement部分
```groovy
dependencyManagement {
    imports {
        mavenBom 'com.amazonaws:aws-java-sdk-bom:1.10.77'
    }
}
```
3. 在依赖部分指定使用的开发工具包模块
```groovy
dependencies {
    compile 'com.amazonaws:aws-java-sdk-s3'
    testCompile group: 'junit', name: 'junit', version: '4.11'
}
```
Gradle 将自动使用 BOM 中的信息来解析开发工具包依赖项的正确版本。
# 链接
* [setup-project-gradle.html](http://docs.aws.amazon.com/zh_cn/sdk-for-java/v1/developer-guide/setup-project-gradle.html)

