---
layout: post
title: Netty Channel Pipeline
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---




# Pipeline

```text
    +------------------------------------------------------------------------------------------------+
    |      +-------------------------+              +-------------------------+                      |
----|----->|  ChannelInboundHandler  |------------->|  ChannelInboundHandler  |----------------------|----->
    |      +-------------------------+              +-------------------------+                      |
    |                                                                                                |         Netty App  
    |                                                                                                |        
    |                 +--------------------------+             +--------------------------+          |
<---|--------<--------|  ChannelOutboundHandler  |<------------|  ChannelOutboundHandler  |<---------|-----
    |                 +--------------------------+             +--------------------------+          |
    +------------------------------------------------------------------------------------------------+

```
注意： 
* 图中的Netty App是Netty应用所在的位置
* 上面的 ----> 是进站，即Inbound
* 下面的 <---- 是出站，即Outbound


# ChannelHandler

|          handler                    |                      Desc                    |
| ----------------------------------- | -------------------------------------------- |
| ChannelInboundHandler               | 拦截和处理入站事件                              |
| ChannelOutboundHandler              | 拦截和处理出站事件                              |


# 事件在ChannelPipeline中流动
```java
ChannelPipeline p = ...;
p.addLast("1", new InboundHandlerA());
p.addLast("2", new InboundHandlerB());
p.addLast("3", new OutboundHandlerA());
p.addLast("4", new OutboundHandlerB());
p.addLast("5", new InboundOutboundHandlerX());
```
对于入站事件，处理顺序为: 1 -> 2 -> 5。对于出站事件，顺序是: 5 -> 4 -> 3。    
事件不会在ChannelPipeline中自动流动，完全由用户控制
```java
public class InboundHandlerA implements ChannelInboundHandler {
        @Override
        public void channelActive(ChannelHandlerContext ctx) {
            System.out.println("Connected!");    // 用户自定义处理逻辑
            ctx.fireChannelActive();             // 将channelActive事件传播到InboundHandlerB
        }
    }

    public class OutboundHandlerB extends ChannelOutboundHandler{
        @Override
        public void close(ChannelHandlerContext ctx, ChannelPromise promise) {
            System.out.println("Closing ..");   // 用户自定义处理逻辑
            ctx.close(promise);                 // 将close事件传播到OutboundHandlerA
        }
    }
```


# 入站事件

入站事件一般由I／O线程触发。

|             event             |                                           |
| ----------------------------- | ----------------------------------------- |
| ChannelRegistered()           | channel注册到EventLoop                     |
| ChannelActive()               | channel激活                                |
| ChannelRead(Object)           | channel读取到数据                           |
| ChannelReadComplete           | channel读取数据完毕                         |
| ExceptionCaught(Throwable)    | 捕获到异常                                  |
| UserEventTriggered(Object)    | 用户自定义事件                               |
| ChannelWritablilityChanged()  | channel可写性改变，由写高低位控制              |
| ChannelInactive()             | Channel不再激活                             |
| ChannelUnregistered()         | channel从EventLoop中注销                    |


# 出站事件

出站事件一般由用户触发

|                          event                          |                                           |
| ------------------------------------------------------- | ----------------------------------------- |
| bind(SocketAddress, ChannelPromise)                     | 绑定到本地地址                              |
| connect(SocketAddress, SocketAddress, ChannelPromise)   | 连接一个远程机器                             |
| write(Object, ChannelPromise)                           | 写数据，写到Netty出站缓冲区                   |
| flush()                                                 | flush数据，实际执行底层写                     |
| read()                                                  | 读数据                                      |
| disconnect(ChannelPromise)                              | 断开连接                                     |
| close(ChannelPromise)                                   | 关闭channel                                 |
| deregister(ChannelPromise)                              | 从EventLoop注销channel                      |
 
# 使用自定义线程池来执行比较耗时的Handler
```java
static final EventExecutorGroup group = new DefaultEventExecutorGroup(16);
...
ChannelPipeline pipeline = ch.pipeline();
// 简单非阻塞业务，可以使用I/O线程执行
pipeline.addLast("decoder", new MyProtocolDecoder());
pipeline.addLast("encoder", new MyProtocolEncoder());
// 复杂耗时业务，使用新的线程池
pipeline.addLast(group, "handler", new MyBusinessLogicHandler());
```
Netty的原则是不阻塞I/O线程。I／O线程即我们在BootStrap中指定的workerGroup，也即Reactor模式中的subReactor。如果handler中车处理比较耗时，应该使用一个自定义的线程池来处理handler。

# Netty发送消息

1. 直接写到Channel中,消息会从channel-pipeline的尾端开始流动          
2. 写入到ChannelHandlerContext,从channel pipeline的下一个handler开始流动




ChannelHandlerContext和Channel，ChannelPipeline有很多类似的方法，比如read，write等。但是有一个重要的区别:    
调用Channel或者ChannelPipeline：沿着整个ChannelPipeline进行传播。即会流进整个ChannelPipeline    
调用ChannelHandlerContext：从当前所关联的ChannelHandler开始。        










```text
                             
                     +---------------------------------------------------------------------------------------------------+
                     |                                                                                                   |
                     |                                                                                                   |
                     |         +-------------------------+   +-----------------------+    +----------------------+       |
                     |         |                         |   |                       |    |                      |       |
                     |         |                         |   |                       |    |                      |       |
+--------------------+----+    |    ChannelHandler       |   |     ChannelHandler    |    |    ChannelHandler    |       |
|                         |    |                         |   |                       |    |                      |       |
|                         |    |                         |   |                       |    |                      |       |
|         Channel         |    +------------+------------+   +------------+----------+    +-----------+----------+       |
|                         |                 |                             |                           |                  |
|                         |                 |                             |                           |                  |
+--------------------+----+                 |                             |                           |                  |
                     |         +------------+------------+   +------------+------------+  +-----------+-------------+    |
                     |         |                         |   |                         |  |                         |    |
                     |         |  ChannelHandlerContext  |   |  ChannelHandlerContext  |  |  ChannelHandlerContext  |    |
                     |         |                         |   |                         |  |                         |    |
                     |         +-------------------------+   +-------------------------+  +-------------------------+    |
                     |                                                                                                   |
                     +---------------------------------------------------------------------------------------------------+


```






