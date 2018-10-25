---
layout: post
title: 通过Spring Profile和Maven Profile配置多环境
date: 2018-10-25 18:20:00
tags:
- Java
categories: Java
---



在Spring Boot中使用profiles

Profile类似于一个应用程序的使用的配置的集合。比如，针对开发，测试和生产环境，应用有各自的配置。
可以针对每一个环境，创建一个profile。

默认spring boot会读取application.properties或application.yaml文件。

举上面的例子，如果开发，测试和生产有3套配置，我们可以创建如下的文件：
* application-develop.properties
* application-test.properties
* application-release.properties

如果是spring boot应用，可以使用`mvn clean install`来生成一个可自行jar，比如叫 hello.jar。
然后在开发环境，可以使用如下命令来运行jar:
```
java -Dspring.profiles.active=develop -jar hello.jar
```

我们看到指定了`-Dspring.profiles.active=develop`。这个配置指定了使用develop这个profile。会让spring boot加载application-develop.properties文件。
Spring Boot会加载::application-{profile}.properties::配置文件。

比如，开发，测试和生产使用不同的数据库连接串。就可以在3个application-xxx.properties文件中分别指定。

那如果想在test环境创建一个Spring Bean，而develop和release不需要。这个怎么操作呢？

这个时候，可以使用spring boot的自动配置来完成。比如，在3个application-xxx.properties文件中都指定：`env=xxx`。
比如application-test.properties中指定`env=test`。然后通过@ConditionalOnProperty注解来根据条件创建Bean。
下面的是一个demo：

```java
@ConditionalOnProperty( name = "env", havingValue = "test")
@Configuration
public class TestEnvConfiguration(){
    
    @Bean
    public FooBean foo(){
        return new FooBean();
    }
}
```

这样，就只会在test环境创建这一个FooBean，在develop和release都不会创建该FooBean。


