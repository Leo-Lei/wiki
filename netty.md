---
layout: post
title: Netty
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
description: The post will introduce a text editor Atom.
---



# EventLoopGroup
是Netty实现的线程池接口。一般的应用中都会使用两个线程池：bossGroup和workderGroup，分别对应Reactor模式中的mainReactor和subReactor，其中boss专门用于接收客户端连接，workder用于处理IO事件。

# ServerBootStrap
为了帮助用户快速构建基于Netty的服务，Netty提供了两个启动器ServerBootstrap和Bootstrap，分别用于启动服务端和客户端程序。


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
             .option(ChannelOption.SO_BACKLOG,128)             // 设置TCP参数
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








