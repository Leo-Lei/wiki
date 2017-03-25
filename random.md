---
layout: post
title: Random in Java
date: 2016-06-24 14:15:00
tags:
- Java
categories: Java
description: The post will cover the Random in Java.
---


```text
╔══════════════╦═════════════════════╦═══════════════════╦══════════════════════╗
║   Property   ║       HashMap       ║      TreeMap      ║     LinkedHashMap    ║
╠══════════════╬═════════════════════╬═══════════════════╬══════════════════════╣
║              ║  no guarantee order ║ sorted according  ║                      ║
║   Order      ║ will remain constant║ to the natural    ║    insertion-order   ║
║              ║      over time      ║    ordering       ║                      ║
╠══════════════╬═════════════════════╬═══════════════════╬══════════════════════╣
║  Get/put     ║                     ║                   ║                      ║
║   remove     ║         O(1)        ║      O(log(n))    ║         O(1)         ║
║ containsKey  ║                     ║                   ║                      ║
╠══════════════╬═════════════════════╬═══════════════════╬══════════════════════╣
║              ║                     ║   NavigableMap    ║                      ║
║  Interfaces  ║         Map         ║       Map         ║         Map          ║
║              ║                     ║    SortedMap      ║                      ║
╠══════════════╬═════════════════════╬═══════════════════╬══════════════════════╣
║              ║                     ║                   ║                      ║
║     Null     ║       allowed       ║    only values    ║       allowed        ║
║ values/keys  ║                     ║                   ║                      ║
╠══════════════╬═════════════════════╩═══════════════════╩══════════════════════╣
║              ║   Fail-fast behavior of an iterator cannot be guaranteed       ║
║   Fail-fast  ║ impossible to make any hard guarantees in the presence of      ║
║   behavior   ║           unsynchronized concurrent modification               ║
╠══════════════╬═════════════════════╦═══════════════════╦══════════════════════╣
║              ║                     ║                   ║                      ║
║Implementation║      buckets        ║   Red-Black Tree  ║    double-linked     ║
║              ║                     ║                   ║       buckets        ║
╠══════════════╬═════════════════════╩═══════════════════╩══════════════════════╣
║      Is      ║                                                                ║
║ synchronized ║              implementation is not synchronized                ║
╚══════════════╩════════════════════════════════════════════════════════════════╝
```
