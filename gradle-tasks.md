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

构建的一些参数

|                    option                                        |   desp                                       | 
| ---------------------------------------------------------------- | -------------------------------------------- | 
| `gradle build -x test`                                           | skip test                                    |
| `gradle build -Pfoo=bar`                                         | 指定自定义参数                                  |
| `gradle --debug build`                                           | 打印执行task时的debug信息                       |
| `gradle --rerun-tasks build`                                     | 强制执行task，gradle会skip掉up-to-date的task    |


# Tasks

```groovy
task hello {
    doLast {
        println 'Hello world!'
    }
}
task intro(dependsOn: hello) {
    doLast {
        println "I'm Gradle"
    }
}
```

Output:

```
> gradle -q intro
Hello world!
I'm Gradle
```



### 向已存在的task添加行为
```groovy
task hello {
    doLast {
        println 'Hello Earth'
    }
}
hello.doFirst {
    println 'Hello Venus'
}
hello.doLast {
    println 'Hello Mars'
}
hello {
    doLast {
        println 'Hello Jupiter'
    }
}
```

Output:

```groovy
> gradle -q hello
Hello Venus
Hello Earth
Hello Mars
Hello Jupiter
```



