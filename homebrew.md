---
layout: post
title: Homebrew
date: 2017-07-11 13:05:00
tags:
- Homebrew
categories: Java
---


|              command          |                    desc                    |
| ----------------------------- | ------------------------------------------ |
| brew install mongo            | 安装mongo                                  |
| brew uninstall mongo          | 卸载mongo                                  |
| brew search git               | 搜索git                                    |
| brew list                     | 显示已经安装的所有软件包                     |
| brew info git                 | brew info git #查看软件包信息               |






brew --help #简洁命令帮助
```bash
$ man brew #完整命令帮助
$ brew install git #安装软件包(这里是示例安装的Git版本控制)
$ brew uninstall git #卸载软件包
$ brew update #同步远程最新更新情况，对本机已经安装并有更新的软件用*标明
$ brew outdated #查看已安装的哪些软件包需要更新
$ brew upgrade git #更新单个软件包
$ brew info git #查看软件包信息
$ brew home git #访问软件包官方站
$ brew cleanup #清理所有已安装软件包的历史老版本
$ brew cleanup git #清理单个已安装软件包的历史版本
```
