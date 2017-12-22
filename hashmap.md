---
layout: post
title: Dive into HashMap
date: 2016-06-22 17:10:00
tags:
- Java
categories: Java
---

# 1. Overview
From Java Doc: 
> Hash table based implementation of the Map interface. This implementation provides all of the optional map operations, and permits null values and the null key. (The HashMap class is roughly equivalent to Hashtable, except that it is unsynchronized and permits nulls.) This class makes no guarantees as to the order of the map; in particular, it does not guarantee that the order will remain constant over time.

```java
public class HashMap<K,V>
    extends AbstractMap<K,V>
    implements Map<K,V>, Cloneable, Serializable{
    ......
}
```

So, there are some essential information as below:      
1. Implement the `Map` interface.     
2. Based on HashTable implementation.      
3. Permit NULL key/value.      
4. Not synchronized.      
5. Not guarantee the order(e.g. the insertion order).      
6. Not guarantee the order will remain unchanged over time.     


# 2. HashMap Data Structure    
Inner HashMap, it is implemented by ***Array*** and ***Link List***.
```text              
┌─────┐    ┌─────┐ 
│  0  │---→| 496 |
|─────|    └─────┘ 
│  1  │
|─────|    ┌─────┐    ┌─────┐  
│  2  │---→| 1   |---→| 337 |
|─────|    └─────┘    └─────┘ 
│  3  │
|─────|
│  4  │
|─────|    ┌─────┐    ┌────┐    ┌─────┐    
│  5  │---→| 12  |---→| 28 |---→| 108 |
|─────|    └─────┘    └────┘    └─────┘ 
│  6  │
└─────┘

```
As shown in the above diagram, inner HashMap, there is an Array, each item of the Array is an Link List.

# 3.The implementation of Put/Get in HashMap     
## 3.1 Put
```java
public V put(K key, V value) {
    // if table is null，inflate the table
    if (table == EMPTY_TABLE) {
        inflateTable(threshold);
    }
    // if key is null，call putForNullKey method,  save null at table the first position，this is why HashMap permit null key/value.
    if (key == null)
        return putForNullKey(value);
    // calculate hash by hashCode the key.
    int hash = hash(key);
    // Find the index in table by the hash code.
    int i = indexFor(hash, table.length);
    // iterate loop the whole array.
    for (Entry<K,V> e = table[i]; e != null; e = e.next) {
        Object k;
        // Check if there is already a item with the same key.
        // If there is, replace it with new value.
        if (e.hash == hash && ((k = e.key) == key || key.equals(k))) {
            V oldValue = e.value;
            e.value = value;
            e.recordAccess(this);
            return oldValue;
        }
    }
    // If the Entry at index i is null, indicate that there is no such item.
    modCount++;
    // add key/value to index of i.
    addEntry(hash, key, value, i);
    return null;
}
```

## 3.2 Get
```java
public V get(Object key) {  
    //  if key = null，return the value which the null point to.
    if (key == null)  
        return getForNullKey(); 
    // calculate hash code by key.
    int hash = hash(key.hashCode()); 
    // find in the Array by the hash code.
    for (Entry<K,V> e = table[indexFor(hash, table.length)];  
        e != null;  
        e = e.next) {  
        Object k;  
        // find the entry by hash code and equals()
        if (e.hash == hash && ((k = e.key) == key || key.equals(k)))  
            return e.value;  
    }  
    // if not found, return null.
    return null;  
}  
```

# 4. How HashMap resolve the hash conflict
1. Initiate a HashMap.

```text              
┌─────┐    ┌─────┐ 
│  0  │---→| 496 |
|─────|    └─────┘ 
│  1  │
|─────|    ┌─────┐    ┌─────┐  
│  2  │---→| 1   |---→| 337 |
|─────|    └─────┘    └─────┘ 
│  3  │
└─────┘
```
2. 







