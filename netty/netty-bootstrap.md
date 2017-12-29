---
layout: post
title: Netty BootStrap
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---

|                                 |    BootStrap             |  ServerBootStrap               |
| ------------------------------- | ------------------------ |--------------------------------|
| 作用                             | 连接到远程主机             | 绑定到一个本地端口                |
| EventLoopGroup数量               |    1                     | 2                              |


# ServerBootStrap
为了帮助用户快速构建基于Netty的服务，Netty提供了两个启动器ServerBootstrap和Bootstrap，分别用于启动服务端和客户端程序。

|          method                 |                      Desc                    |
| ------------------------------- | -------------------------------------------- |
| group(EventLoopGroup...)        | 指定一个或多个Reactor                          |
| channel(Channel)                | 指定一个Channel工厂                            |
| option(Key,Value)               | 指定TCP相关的参数以及Netty自定义的参数            |
| childHandler()                  | 指定subReactor中的处理器                       |
| handler()                       | 指定mainReactor的处理器                        |

### handler
默认情况下，mainReactor中已经添加了acceptor处理器，所以无需再指定。


# Handler

### ChannelInitializer
这是一个特殊的Handler，功能是初始化多个Handler。handler()和childHandler()方法并不能多次调用以达到增加多个Handler的目的，所以引入了ChannelInitializer。


```java
b.childHandler(new ChannelInitializer<SocketChannel>() {
         @Override
         public void initChannel(SocketChannel ch) throws Exception {
             ChannelPipeline p = ch.pipeline();
             p.addLast(new DecoderHandler());   // 解码处理器
             p.addLast(new EncoderHandler());   // 编码处理器
             p.addLast(threadPool, new ComputeWithSqlHandler());   // 附带SQL查询的计算
         }
    });
```




# Netty的demo

```java
package com.leibangzhu.netty.sample;

import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;


public class DiscardServer {

    private int port;

    public DiscardServer(int port){
        this.port = port;
    }

    public void run() throws Exception{
        EventLoopGroup bossGroup = new NioEventLoopGroup();         // main Reactor，接收客户端请求
        EventLoopGroup workerGroup = new NioEventLoopGroup();       // sub Reactor，处理客户端请求 
        try {
            ServerBootstrap b = new ServerBootstrap();
            b.group(bossGroup,workerGroup)
             .channel(NioServerSocketChannel.class)
             .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        protected void initChannel(SocketChannel ch) throws Exception {
                            ch.pipeline().addLast(new DiscardServerHandler());
                        }
             })
             .option(ChannelOption.SO_BACKLOG,128)             // 设置TCP参数
             .childOption(ChannelOption.SO_KEEPALIVE,true);
            // 绑定到本地端口等待客户端连接
            ChannelFuture f = b.bind(port).sync();
            // 等待接收客户端连接的Channel被关闭
            f.channel().closeFuture().sync();
        }finally {
            workerGroup.shutdownGracefully();
            bossGroup.shutdownGracefully();
        }
    }

    public static void main(String[] args) throws Exception {
        int port = 8080;
        new DiscardServer(port).run();
    }
}

```
