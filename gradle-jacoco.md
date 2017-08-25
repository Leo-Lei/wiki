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

# 环境
* Gradle版本4.1


# build.gradle
```groovy
version 'unspecified'

apply plugin: 'java'
apply plugin: "jacoco"

sourceCompatibility = 1.8

repositories {
    mavenCentral()
}

dependencies {
    testCompile group: 'junit', name: 'junit', version: '4.11'
}

jacoco {
    toolVersion = "0.7.6.201602180812"
    reportsDir = file("$buildDir/customJacocoReportDir")
}

jacocoTestCoverageVerification {
    violationRules {
        rule {
            limit {
                minimum = 0.8
            }
        }

        rule {
            enabled = false
            element = 'CLASS'
            includes = ['org.gradle.*']

            limit {
                counter = 'LINE'
                value = 'TOTALCOUNT'
                maximum = 0.3
            }
        }
    }
}

```



