---
layout: post
title: Equals and HashCode Contract in Java
date: 2016-06-22 18:20:00
tags:
- Java
categories: Java
---

# 1. Overview
The Contract between equals() and hashCode() is:           
1. If two objects are equal, then they must have the same hash code.              
2. If two objects have the same hash code, they may or may not be equal.               


