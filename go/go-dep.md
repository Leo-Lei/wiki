---
layout: post
title: Go dep
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


Dep是Go语言官方推出的一个依赖管理工具。

Dep通过两个文件来管理依赖:
1. `Gopkg.toml`
2. `Gopkg.lock`
其中`Gopkg`通过命令生成，也可以被用户根据需要手动修改。`Gopkg.lock`是自动生成的，不可修改。
