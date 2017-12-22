---
layout: post
title: Map in Java
date: 2016-06-22 14:20:00
tags:
- Java
categories: Java
---

# 1. Map Overview     
If we use one sentence to describe each Map implementations, it would be as the following:
* **HashMap** is implemented as a hash table, and there is **no** ordering on keys or values.
* **TreeMap** is implemented based on red-black tree structure, and it is ordered by the key.
* **LinkedHashMap** preserves the insertion order.
* **HashTable** is synchronized, in contrast to HashMap.     

| Class            | Ordering?           |  Synchronized?  |
| ---------------- | --------------------| --------------- |
| HashMap          | ---                 |  ---            |
| LinkedHashMap    | insertion-order     |                 |
| TreeMap          | ordered by key      |                 |
| HashTable        | ---                 |  Y              |

> This gives us the reason that HashMap should be used if it is thread-safe, since HashTable has overhead for synchronization.           

# 2. HashMap      
Some essential points of HashMap:          
1. If key of the HashMap is self-defined objects, then [equals() and hashCode() contract](http://www.programcreek.com/2011/07/java-equals-and-hashcode-contract/) need to be followed.    
2. HashMap is **not** ordering.

Below example illustrate the usage of HashMap.           
```java
public class Book {
    String isbn;

    Book(String isbn) {
        this.isbn = isbn;
    }

    public boolean equals(Object o){
        return ((Book)o).isbn.equals(this.isbn);
    }

    public int hashCode(){
        return isbn.length();
    }

    public String toString(){
        return isbn + " book";
    }
}

class TestMap {
    public static void main(String[] args) {
        Map<Book,String> hashMap = new HashMap<Book,String>();
        Book b1 = new Book("001");
        Book b2 = new Book("123");
        Book b3 = new Book("010");
        Book b4 = new Book("010");

        hashMap.put(b1, "b1");
        hashMap.put(b2, "b2");
        hashMap.put(b3, "b3");
        hashMap.put(b4, "b4");

//loop HashMap
        for (java.util.Map.Entry<Book,String> entry : hashMap.entrySet()) {
            System.out.println(entry.getKey().toString() + " - " + entry.getValue());
        }
    }
}
```
The output:       
```bash
010 book - b4
123 book - b2
001 book - b1
```

> Note: If we didn't overwrite the equals and hashCode function, there will be 2 "010" books in the output. But we expect only the book "b4".

# 3. TreeMap       
1. A TreeMap is soreted by keys.        
2. The key has to implements the `Comparable` interface.
3. The key do not need to overwrite equals() and hashCode() methods.

Let's take a look at the following example to understand the "sorted by keys" idea.    
```java
class Dog {
    String color;

    Dog(String c) {
        color = c;
    }

    public boolean equals(Object o) {
        return ((Dog) o).color.equals(this.color);
    }

    public int hashCode() {
        return color.length();
    }

    public String toString(){
        return color + " dog";
    }
}

class TestHashMap {
    public static void main(String[] args) {
        Map<Dog,Integer> hashMap = new TreeMap<Dog,Integer>();
        Dog d1 = new Dog("red");
        Dog d2 = new Dog("black");
        Dog d3 = new Dog("white");
        Dog d4 = new Dog("white");

        hashMap.put(d1, 10);
        hashMap.put(d2, 15);
        hashMap.put(d3, 5);
        hashMap.put(d4, 20);

//loop HashMap
        for (java.util.Map.Entry<Dog,Integer> entry : hashMap.entrySet()) {
            System.out.println(entry.getKey().toString() + " - " + entry.getValue());
        }
    }
}

```
While running the example, a `ClassCastException` error occurred.    
```bash
Exception in thread "main" java.lang.ClassCastException: Dog cannot be cast to java.lang.Comparable
```
Since TreeMaps are sorted by keys, the object for key has to be able to compare with each other, that's why it has to implement `Comparable` interdace.

Change the Book class as below:
```java
public class Book implements Comparable<Book> {
    String isbn;
    int price;

    Book(String isbn,int price) {
        this.isbn = isbn;
        this.price = price;
    }

    public String toString(){
        return isbn + " book";
    }

    @Override
    public int compareTo(Book o) {
        return o.price - this.price;
    }
}


class TestMap {
    public static void main(String[] args) {
        Map<Book,String> hashMap = new TreeMap<Book,String>();
        Book b1 = new Book("001",30);
        Book b2 = new Book("123",20);
        Book b3 = new Book("010",10);
        Book b4 = new Book("010",10);

        hashMap.put(b1, "b1");
        hashMap.put(b2, "b2");
        hashMap.put(b3, "b3");
        hashMap.put(b4, "b4");

//loop HashMap
        for (java.util.Map.Entry<Book,String> entry : hashMap.entrySet()) {
            System.out.println(entry.getKey().toString() + " - " + entry.getValue());
        }
    }
}

// output:
// 001 book - b1
// 123 book - b2
// 010 book - b4
```

If we change the book4 from "Book b4 = new Book("010",10)" to "Book b4 = new Book("010",40)", it will produce the following output:   
```java
        Book b1 = new Book("001",30);
        Book b2 = new Book("123",20);
        Book b3 = new Book("010",10);
        Book b4 = new Book("010",40);
        
// output:
// 010 book - b4
// 001 book - b1
// 123 book - b2
// 010 book - b3
```

# 4. HashTable

From Java Doc: *The HashMap class is roughly equivalent to Hashtable, except that it is unsynchronized and permits nulls*. So, HashTable is:    
1. Not ordering.
2. Synchronized, i.e. thread-safe.

# 5. LinkedHashMap

LinkedHashMap maintains the insertion-order.
Let's replace the HashMap with LinkedHashMap using the same code used for HashMap.    
```java
class Book{
......
}
Map<Book,String> hashMap = new LinkedHashMap<Book,String>();
......
```
Output is:
```bash
001 book - b1
123 book - b2
010 book - b4
```



