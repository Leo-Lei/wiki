---
layout: post
title: Multiple projects in Gradle
date: 2015-06-28 18:58:20
tags:
- Gradle
categories: 
- Java
- Gradle
---



# 显示所有的project
```groovy
> gradle -q projects

------------------------------------------------------------
Root project
------------------------------------------------------------

Root project 'multiproject'
+--- Project ':api'
+--- Project ':services'
|    +--- Project ':services:shared'
|    \--- Project ':services:webservice'
\--- Project ':shared'

To see a list of the tasks of a project, run gradle <project-path>:tasks
For example, try running gradle :api:tasks

```


