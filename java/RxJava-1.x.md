---
title: RxJava
date: 2017-06-09 10:22:23
categories:
- Music
tags:
- Music
---


# RxJava基本概念
* Observable: 可观察者，即被观察者
* Observer: 观察者，即订阅者
* subscribe: 订阅。
* 事件:
其中:
* Observable和Observer通过subscribe()方法实现订阅关系。从而Observable可以在需要的时候发出事件来通知Observer。






# Observer
```java
Observer<String> observer = new Observer<String>() {
    @Override
    public void onCompleted() {
        System.out.println("Completed");
    }

    @Override
    public void onError(Throwable e) {
        e.printStackTrace();
    }

    @Override
    public void onNext(String s) {
        System.out.println(s);
    }
};
```

# subscriber
```java
Subscriber<String> subscriber = new Subscriber<String>() {
    @Override
    public void onCompleted() {
        // ...
    }

    @Override
    public void onError(Throwable e) {
        // ...
    }

    @Override
    public void onNext(String s) {
        // ...
    } 
    
    @Override
    public void onStart() {
        // ...
    }
};
subscriber.onStart();
subscriber.unsubscribe();
```
RxJava的Observer和Subscriber大部分用法是一样的。它们的区别有两点:
* onstart(): Subscriber添加的方法。可用于在事件还未触发前，做一些准备工作。
* unsubscribe(): Subscriber添加的新方法，用于取消订阅。

# subscribe()方法
通过subscribe()方法将observable和observer关联起来。
```java
observable.subscribe(observer);
observable.subscribe(subscriber);
```

subscribe()方法内部实现是这样子的(仅核心代码):
```java
// 注意：这不是 subscribe() 的源码，而是将源码中与性能、兼容性、扩展性有关的代码剔除后的核心代码。
// 如果需要看源码，可以去 RxJava 的 GitHub 仓库下载。
public Subscription subscribe(Subscriber subscriber) {
    subscriber.onStart();
    onSubscribe.call(subscriber);
    return subscriber;
}
```

```java
// 自动创建 Subscriber ，并使用 onNextAction 来定义 onNext()
observable.subscribe(onNextAction);
// 自动创建 Subscriber ，并使用 onNextAction 和 onErrorAction 来定义 onNext() 和 onError()
observable.subscribe(onNextAction, onErrorAction);
// 自动创建 Subscriber ，并使用 onNextAction、 onErrorAction 和 onCompletedAction 来定义 onNext()、 onError() 和 onCompleted()
observable.subscribe(onNextAction, onErrorAction, onCompletedAction);
```


