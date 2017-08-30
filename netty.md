---
layout: post
title: Netty
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---



# NioEventLoopGroup
代表一个线程池来处理客户端发起的IO请求。

# ServerBootStrap
一个工具类，来创建一个Server


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
        EventLoopGroup bossGroup = new NioEventLoopGroup();               // main Reactor，接收客户端请求
        EventLoopGroup workerGroup = new NioEventLoopGroup();             // sub Reactor，处理客户端请求 
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
             .option(ChannelOption.SO_BACKLOG,128)
             .childOption(ChannelOption.SO_KEEPALIVE,true);
            
            ChannelFuture f = b.bind(port).sync();

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








