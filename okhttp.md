---
layout: post
title: okhttp
date: 2017-03-12 11:10:00
tags:
- docker
categories: Java
---



```java
// 默认连接数是5，改成50
ConnectionPool connectionPool = new ConnectionPool(50, 5L, TimeUnit.MINUTES);

private final OkHttpClient client = new OkHttpClient.Builder()
        .connectionPool(connectionPool)
        .build();
```


```java
    String url = "http://127.0.0.1:20000/http";

    RequestBody formBody = new FormBody.Builder()
            .add("hello","world")
            .add("foo","bar")
            .build();

    Request request = new Request.Builder()
            .url(url)
            .post(formBody)
            .build();

    try (Response response = client.newCall(request).execute()) {
        if (!response.isSuccessful()) throw new IOException("Unexpected code " + response);

        String result = response.body().string();
        System.out.println(result);
    }
```
