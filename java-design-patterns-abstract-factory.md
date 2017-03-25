---
layout: post
title: Java Design Pattern - Abstract Factory
date: 2015-06-26 15:30:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the Abstract Factory design pattern in Java.
---

# 1. An Example
**KingdomFactory**: the abstract factory.
```java
public interface KingdomFactory {
    Castle createCastle();
    King createKing();
    Army createArmy();
}
```
**Castle**,**King** and **Army**: the *Product* interface.
```java
public interface Castle {
}
public interface King {
}
public interface Army {
}
```
**ElfKingdomFactory** and **OrcKingdomFactory**:
```java
public class ElfKingdomFactory implements KingdomFactory{
    @Override
    public Castle createCastle() {
        return new ElfCastle();
    }

    @Override
    public King createKing() {
        return new ElfKing();
    }

    @Override
    public Army createArmy() {
        return new ElfArmy();
    }
}

public class OrcKingdomFactory implements KingdomFactory{
    @Override
    public Castle createCastle() {
        return new OrcCastle();
    }

    @Override
    public King createKing() {
        return new OrcKing();
    }

    @Override
    public Army createArmy() {
        return new OrcArmy();
    }
}
```
**ElfCastle**,**ElfKing**,**ElfArmy**,**OrcCastle**,**OrcKing** and **OrcArmy**: the concrete products.
```java
public class ElfArmy implements Army{

    @Override
    public String toString(){
        return "This is the Eleven army!";
    }
}

public class ElfCastle implements Castle{

    @Override
    public String toString(){
        return "This is the Elven castle!";
    }
}

public class ElfKing implements King{

    @Override
    public String toString(){
        return "This is the Elven king!";
    }
}

// the OrcCastle, OrcKing and OrcArmy is similar with Elf.
```
**App**: the client
```java
public class App {

    public static void main(String[] args){
        createKingdom(new ElfKingdomFactory());
        createKingdom(new OrcKingdomFactory());
    }

    private static void createKingdom(KingdomFactory factory){
        King king = factory.createKing();
        Castle castle = factory.createCastle();
        Army army = factory.createArmy();

        System.out.println("The kingdom was created.");
        System.out.println(king);
        System.out.println(castle);
        System.out.println(army);
    }
}
```



