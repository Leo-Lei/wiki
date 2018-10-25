---
layout: post
title: Base useage of Maven
date: 2015-06-26 15:30:00
tags:
- Maven
categories: Java
---


# 从命令行安装maven依赖
```bash
mvn org.apache.maven.plugins:maven-dependency-plugin:2.4:get -DartifactId=jersey-servlet -DgroupId=com.sun.jersey -Dversion=1.19
```

# Maven命令

| command                                                 |             Description                                       |                   
| ------------------------------------------------------- | ------------------------------------------------------------- |
| `mvn clean`                                             | Clean the *target* folder.                                    |
| `mvn clean install`                                     | Install artifact(jar or war) to local maven repository.       |
| `mvn dependency:tree -Dverbose -Dincludes=asm:asm`      | 检查maven依赖树                                                 |
| `mvn install -Dmaven.test.skip=true`                    | 跳过单元测试                                                    |
| `mvn clean install -U`                                  | 使用`-U`强制更新依赖                                            |

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

# settings.xml例子
```xml
<?xml version="1.0" encoding="UTF-8"?>

<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
          xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
          xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">
  <!-- localRepository
   | The path to the local repository maven will use to store artifacts.
   |
   | Default: ${user.home}/.m2/repository
  <localRepository>/path/to/local/repo</localRepository>
  -->
  <pluginGroups>

  </pluginGroups>

  <proxies>
  </proxies>
  
  <servers>
      <server>
			<id>releases</id>
			<username>admin</username>
			<password>admin246</password>
		</server>
		<server>
			<id>snapshots</id>
			<username>admin</username>
			<password>admin246</password>
		</server>
  </servers>

   <mirrors>
    <!--  <mirror>
      <id>alimaven</id>
      <name>aliyun maven</name>
      <url>http://maven.aliyun.com/nexus/content/groups/public/</url>
      <mirrorOf>central</mirrorOf>
    </mirror> -->
 </mirrors>

   <profiles>
     <profile>
         <id>pnt</id>
         <activation>
             <jdk>1.8</jdk>
         </activation>
         <properties>
             <releases.repo>http://192.168.4.10:8087/nexus/content/repositories/releases</releases.repo>
             <snapshots.repo>http://192.168.4.10:8087/nexus/content/repositories/snapshots</snapshots.repo>
         </properties>
         <repositories>
             <repository>
                 <id>nexus</id>
                 <name>local private nexus</name>
                 <url>http://192.168.4.10:8087/nexus/content/groups/public/</url>
                 <releases>
                     <enabled>true</enabled>
                 </releases>
                 <snapshots>
                     <enabled>true</enabled>
                 </snapshots>
             </repository>
         </repositories>
         <pluginRepositories>
             <pluginRepository>
                 <id>nexus</id>
                 <name>local private nexus</name>
                 <url>http://192.168.4.10:8087/nexus/content/groups/public/</url>
                 <releases>
                     <enabled>true</enabled>
                 </releases>
                 <snapshots>
                     <enabled>true</enabled>
                 </snapshots>
             </pluginRepository>
         </pluginRepositories>
     </profile>
 </profiles>

 <activeProfiles>
   <activeProfile>pnt</activeProfile>
</activeProfiles>
</settings>
```


# Profile配置
```
<project>
    <properties>...</properties>
    <dependencyManagement>...</dependencyManagement>
    <dependencies>...</dependencies>
    <build>...</build>
    <profiles>
        <profile>
            <id>foo</id>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
            <dependencies>
                <dependency>
                    <groupId>org.springframework</groupId>
                    <artifactId>spring</artifactId>
                    <version>1.0.0</version>
                </dependency>
            </dependencies>
        </profile>
        <profile>
            <id>bar</id>
            <activation>
                <activeByDefault>false</activeByDefault>
            </activation>
            <dependencies>
                <dependency>
                    <groupId>log4j</groupId>
                    <artifactId>log4j</artifactId>
                    <version>1.0.0</version>
                </dependency>
            </dependencies>
        </profile>
    </profiles>
</project>

```

