---
layout: post
title: Authentication in Web application.
date: 2016-07-14 15:30:00
tags:
- Java
categories: Java
---



# Token认证
App端登陆时，提供用户名和密码，服务端校验用户名和密码，如果正确，生成一个token返回给app端。app端会将token保存在本地，接下来的请求会带上这个token。服务端会将token保存在redis中，比如用户id是userid001，那么redis中的key就是user_token_userid001。
Token中包括的信息可以分为2部分。XXXXXXXX:53453。前面一段是clientid + 当前时间，然后加密一下。加密的算法支持反向解密。后面的53453也是一个时间戳。每次app端发请求过来，都会更新redis里面的这个时间戳。
比如：用户输入用户名tom和密码123456，服务端校验通过，生成一个token。同时会给用户分配一个userId，比如是userId001。
```java
String token = encrypt("userId001" + ":" + new Date())    //这里加的Date是为了保证token不一样，用随机数代替date也可以的。
```
比如客户在多个客户端ios，android登陆，如果token里只包含userId，那么多个客户端是一样的token。做不到同时只能有一个客户端登陆，登陆了ios的，再登陆android会将ios的登陆踢掉。
这个token是会传给客户端的。`5dhfgbfg0trgdgfgfdfgd`。 考虑到token过期，服务端会在这个token后加上一个时间，保存在redis里面。`5dhfgbfg0trgdgfgfdfgd:34753459`
```java
token = token + new Date();
```
客户端拿到token后，保存到本地。接下来的请求，会把token在http头里面带上,`5dhfgbfg0trgdgfgfdfgd`.服务端拿到token后，进行解密，得到userId + Date.根据userId去redis中找key:比如是userId001。找到`5dhfgbfg0trgdgfgfdfgd:34753459`.比较5dhfgbfg0trgdgfgfdfgd是不是和客户端传来的token一致。然后比较redis里面最后面的34753459，把这个时间戳和当前时间比较。如果相差的比较大，比如超过了10天，那么认为token过期了。将token失效，需要用户重新登陆。如果在token失效期内，认为可以登陆，更新redis中的最后面的时间戳为当前时间。这样之后用户在连续10天都没有登陆的情况下，才会token过期，需要重新登陆。

为什么不直接设置redis的过期时间？过期时间是一个int，单位是毫秒。我们的过期时间是一个月，是超过int的最大值的，所以，这个方案就不能用了。




