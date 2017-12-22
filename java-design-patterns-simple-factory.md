---
layout: post
title: Java Design Pattern - Simple Factory
date: 2015-09-02 15:15:00
tags:
- Linux
categories: Linux
---

# 1. An Example
**Weapon**: The interface.
```java
public interface Weapon {
}
```
**ShortSword**,**Spear** and **AXE**: Some concrete classes of **Weapon** interface.
```java
public class ShortSword implements Weapon{
}
public class Spear implements Weapon {
}
public class AXE implements Weapon{
}
```
**WeaponFactory**: the simple factory.
```java
public class WeaponFactory {
    public static Weapon create(WeaponType type){
        Weapon weapon = null;
        switch (type){
            case SHORT_SWORD:
                weapon = new ShortSword();
                break;
            case SPEAR:
                weapon = new Spear();
                break;
            case AXE:
                weapon = new AXE();
        }
        return weapon;
    }
}
```
**App**: the client class.
```java
public class App {

    public static void  main(String[] args){

        Weapon weapon;

        weapon = WeaponFactory.create(WeaponType.SHORT_SWORD);
        System.out.println(weapon);

        weapon = WeaponFactory.create(WeaponType.SPEAR);
        System.out.println(weapon);

        weapon = WeaponFactory.create(WeaponType.AXE);
        System.out.println(weapon);
    }
}
```




