---
layout: post
title: Java Design Pattern - Factory Method
date: 2015-09-02 15:10:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the Abstract Factory design pattern in Java.
---

# 1. An Example

**Weapon**: the interface
```java
public interface Weapon {
}
```
**ElfWeapon**,**OrcWeapon**: the concrete weapons.
```java
public class ElfWeapon implements Weapon{

    private WeaponType weaponType;

    public ElfWeapon(WeaponType weaponType){
        this.weaponType = weaponType;
    }

    @Override
    public String toString(){
        return "Elf " + weaponType;
    }
}

package com.leo.designpattern.factorymethod;

public class OrcWeapon implements Weapon{

    private WeaponType weaponType;

    public OrcWeapon(WeaponType weaponType){
        this.weaponType = weaponType;
    }

    @Override
    public String toString(){
        return "Orc " + weaponType;
    }
}
```
**Blacksmith**: the creater interface.
```java
public interface BlackSmith {
    Weapon create(WeaponType type);
}
```
**ElfBlacksmith**,**OrcBlacksmith**: the concrete implementation of Blacksmith(*Creator*)
```java
public class ElfBlackSmith implements BlackSmith{
    @Override
    public Weapon create(WeaponType type) {
        return new ElfWeapon(type);
    }
}

public class OrcBlackSmith implements BlackSmith{
    @Override
    public Weapon create(WeaponType type) {
        return new OrcWeapon(type);
    }
}
```
**App**: the client.
```java
public class App {

    public static void main(String[] args) {
        BlackSmith blacksmith;
        Weapon weapon;

        blacksmith = new OrcBlackSmith();
        weapon = blacksmith.create(WeaponType.SPEAR);
        System.out.println(weapon);
        weapon = blacksmith.create(WeaponType.AXE);
        System.out.println(weapon);

        blacksmith = new ElfBlackSmith();
        weapon = blacksmith.create(WeaponType.SHORT_SWORD);
        System.out.println(weapon);
        weapon = blacksmith.create(WeaponType.SPEAR);
        System.out.println(weapon);
    }
}
```










