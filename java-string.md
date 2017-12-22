---
layout: post
title: Dive a Little Deep to the String in Java
date: 2015-06-26 15:30:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the useage of Linux.
---

# 1. String Memory Model
Where is a string stored in memory? We are well aware that String is not a primitive type, so it should be stored in heap. But String is a little special than other normal reference objects.
In a word, String will be stored in **Literal String Pool** or heap.

# 2. Constant String Pool
There is a Constant String Pool in Java, the pool maintain all the literal String. For example, we have below java code:
`String str = "abc"`. Java compiler will first try to check whether "abc" already exists in the Constant String Pool. If it exists, assign the address of "abc" in Constant String Pool to variable str;
If "abc" not exists in Constant String Pool, it will be created in Constant String Pool, and the address of new string will be assigned to variable str.

Creating a string object and adding it to the Constant String Pool is executed in Java Compile phase.
When a string is not declared as a literal(constant) string, it will be created on heap, like other non-primitive type objects.
For example, bellow clause will create string object on heap, rather than Constant String Pool.

`String s = SomeFunRetuenStrType()`
`String s = new String("abc")`

Below sample illustrate when string stored in Constant String Pool or heap.

```java
String _s0 = "foo";          //_s0 in Constant String Pool
String _s1 = "bar";         //_s1 in Constant String Pool

String s0 = "foobar";        //s0 in Constant String Pool
String s1 = "foobar";        //s1 in Constant String Pool
String s2 = "foo" + "bar";   //s2 is made up of 2 constant string, so s2 is also a constant string
String s3 = _s0 + _s1;      //s3 is not a constant string, in heap

System.out.println(s0 == s1);    //true. s0 and s1 both point to the same object in Constant String Pool.
System.out.println(s0 == s2);    //true. s0 and s2 both point to the same object in Constant String Pool.
System.out.println(s0 == s3);    //false. s0 point to object in Constant String Pool, while s3 point to object in heap.
```

```java
String s0 = "foobar";                   //Constant string. Located at Constant String Pool.
String s1 = new String("foobar");       //Not a constant string, located in heap.
String s2 = "foo" + new String("bar");  //Not a constant string, located in heap. And the address is different with s1.

System.out.println(s0 == s1);    //false
System.out.println(s0 == s2);    //false
System.out.println(s1 == s2);    //false
```

```java
String s1 = "foo";
String s2 = "bar";
String s3 = s1 + s2;
String s4 = "foo" + "bar";

System.out.println(s3 == "foobar");   //false
System.out.println(s4 == "foobar");   //true
```

## 2.1 String.intern() method
The intern method is designed to add a string to Constant String Pool at runtime.
When an instance of String call this method, JVM will find the string in Constant String Pool, if it exists, return the address, if not exists, create the string in Constant String Pool, and return the address of new-created String object.

Pay attention that, calling intern method will not affect existing string object on heap.

Bellow sample illustrate the usage of the intern method.

```java
String s0= "foobar";
String s1 = new String("foobar");
String s2 = new String("foobar");
System.out.println( s0 == s1 );             //false

s1.intern();
s2 = s2.intern(); //Assign the address of "foobar" located in Constant String Pool to variable s2
System.out.println( s0 == s1);              //false
System.out.println( s0 == s1.intern() );    //true
System.out.println( s0 == s2 );             //true
```

```java
String s1 = new String("foobar");
String s2 = s1.intern();
System.out.println(s1 == s1.intern());   //false
System.out.println(s2 == s1.intern());   //true
```


