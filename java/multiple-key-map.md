---
layout: post
title: Multiple Keys Map
date: 2017-07-05 12:20:00
tags:
- Java
categories: Java
---



```java

import java.util.LinkedHashMap;
import java.util.Map;

public class MultiKeyMap3<T1,T2,T3,T4> {

    private Map<Key3<T1,T2,T3>,T4> map = new LinkedHashMap<Key3<T1, T2,T3>, T4>();

    public void put(T1 t1, T2 t2, T3 t3,T4 t4){
        map.put(new Key3<>(t1,t2,t3),t4);
    }

    public T4 get(T1 t1,T2 t2,T3 t3){
        return map.get(new Key3<>(t1,t2,t3));
    }

    public boolean containsKey(T1 t1,T2 t2,T3 t3){
        Key3<T1,T2,T3> key = new Key3<T1,T2,T3>(t1,t2,t3);
        return map.containsKey(key);
    }

    private class Key3<T1,T2,T3>{
        private T1 t1;
        private T2 t2;
        private T3 t3;

        public Key3(T1 t1,T2 t2,T3 t3){
            this.t1 = t1;
            this.t2 = t2;
            this.t3 = t3;
        }

        @Override
        public boolean equals(Object o){
            if (this == o){return true;}
            if (!(o instanceof Key3)){return false;}
            Key3 key = (Key3) o;
            return t1.equals(key.t1) && t2.equals(key.t2) && t3.equals(key.t3);
        }

        public int hashCode(){
            int result = t1.hashCode() + t2.hashCode() + t3.hashCode();
            return result;
        }
    }
}

```
