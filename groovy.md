---
layout: post
title: Groovy
date: 2016-08-05 17:10:00
tags:
- Gradle
categories: 
- Java
- Gradle
---

Command chains          

```groovy
// equivalent to: turn(left).then(right)
turn left then right

// equivalent to: take(2.pills).of(chloroquinine).after(6.hours)
take 2.pills of chloroquinine after 6.hours

// equivalent to: paint(wall).with(red, green).and(yellow)
paint wall with red, green and yellow

// with named parameters too
// equivalent to: check(that: margarita).tastes(good)
check that: margarita tastes good

// with closures as parameters
// equivalent to: given({}).when({}).then({})
given { } when { } then { }

// equivalent to: select(all).unique().from(names)
select all unique() from names
```


