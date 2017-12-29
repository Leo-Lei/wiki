---
layout: post
title: Netty Event Loop
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---


# Event Loop

* 一个EventLoop包含一个或多个EventLoop
* 一个EventLoop在它的生命周期只和一个Thread绑定
* 所有由EventLoop处理的io事件都将在它专有的Thread上被处理
* 一个 Channel在它的生命周期内只注册于一个EventLoop

