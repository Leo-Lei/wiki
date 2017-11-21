---
title: Java Generic
date: 2017-06-09 10:22:23
categories:
- Music
tags:
- Music
---

# Java Generic
```java
public class Foo<K,V> {
    
    public <K,V> Bar<K,V> build(){
        return new Bar();
    }
}
```

```java
public class Bar<K,V> {
    public void hello(K k,V v){
        // ......
    }
}
// 通过new Foo().build()获取泛型的Bar
Bar<String,String> bar1 = new Foo().build();
Bar<String,Integer> bar2 = new Foo().build();
```

