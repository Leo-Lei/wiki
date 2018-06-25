---
layout: post
title: Netty Decoder
date: 2016-11-16 10:20:00
tags:
- Atom
categories: Text Editor
---


# Decoder
Decoder就是解码器。解码器就是讲一个字节序列转换成另一种格式(一个消息或另一个字节序列)。    
Decoder是用于处理入站数据的。Netty中的Decoder都继承于ChannelInboundHandler。





# ByteToMessageDecoder
Netty提供了一个抽象类ByteToMessageDecoder来处理字节序列到消息的解码。由于你不可能知道远程节点是否会一次性的发送一个完整的消息，所以这个类会对入站数据进行缓冲，直到它准备好处理。

ByteToMessageDecoder本身也是一个ChannelInboundHandler。它做了以下的事情：
1. 每当新的数据进来了，它会判断上一次是不是还有数据没有处理完，比如上一次是一个半包的数据，无法处理。如果上一次有数据没处理完，那么ByteToMessageDecoder会将这次的数据和上一次的数据进行合并，然后再处理。
2. 数据合并之后，继续进行处理。有一个循环，如果ByteBuf中有可读的数据，就不断去执行用户定义的decode()方法。具体细节如下:
    - 1. 如果out为空，且ByteBuf的读索引没变，说明ByteBuf中数据不全，退出循环。
    - 2. 如果out不为空，说明已经解析出一些消息了，通过fireChannelRead方法，将这些消息传递给后面的ChannelHandler。
3. 处理完之后，将cumulation中已读的字节discard掉。


以下是入口函数

```java
// 一个全局变量。用于保存上一次处理剩下的数据。用于缓存半包数据。
// 每次有新的数据来，都会首先和上一次剩下的数据进行合并
ByteBuf cumulation;        
```

入口函数channelRead
```java
public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        if (msg instanceof ByteBuf) {
            CodecOutputList out = CodecOutputList.newInstance();    //  创建一个新的out list
            try {
                ByteBuf data = (ByteBuf) msg;
                first = cumulation == null;
                if (first) {
                    cumulation = data;
                } else {
                    cumulation = cumulator.cumulate(ctx.alloc(), cumulation, data);      // 合并本次数据和上次遗留数据
                }
                callDecode(ctx, cumulation, out);                  // 处理数据
            } finally {
                if (cumulation != null && !cumulation.isReadable()) { // cumulation的数据刚好都读完了，比如刚好是一个或若干个message
                    numReads = 0;
                    cumulation.release();
                    cumulation = null;
                } else if (++ numReads >= discardAfterReads) {       // cumulation的数据没有完全读完，有残留。
                    // We did enough reads already try to discard some bytes so we not risk to see a OOME.
                    // See https://github.com/netty/netty/issues/4275
                    numReads = 0;
                    discardSomeReadBytes();        // 将cumulation已读的字节discard掉, 只留未读的字节。
                }

                int size = out.size();
                decodeWasNull = !out.insertSinceRecycled();
                fireChannelRead(ctx, out, size);
                out.recycle();
            }
        } else {
            ctx.fireChannelRead(msg);
        }
    }
```


fireChannelRead方法
```java
static void fireChannelRead(ChannelHandlerContext ctx, CodecOutputList msgs, int numElements) {
        for (int i = 0; i < numElements; i ++) {
            ctx.fireChannelRead(msgs.getUnsafe(i));
        }
    }
```
遍历out list中的所有消息，依次调用ChannelHandlerContext.fireChannelRead方法。


```java
protected void callDecode(ChannelHandlerContext ctx, ByteBuf in, List<Object> out) {

            while (in.isReadable()) {
                int outSize = out.size();

                if (outSize > 0) {            // 如果out list中有数据了，说明已解析出数据了，将msg传给后面的Handler。
                    fireChannelRead(ctx, out, outSize);
                    out.clear();

                    // Check if this handler was removed before continuing with decoding.
                    // If it was removed, it is not safe to continue to operate on the buffer.
                    //
                    // See:
                    // - https://github.com/netty/netty/issues/4635
                    if (ctx.isRemoved()) {
                        break;
                    }
                    outSize = 0;
                }

                int oldInputLength = in.readableBytes();
                decodeRemovalReentryProtection(ctx, in, out);      // 调用用户实现的decode方法

                // Check if this handler was removed before continuing the loop.
                // If it was removed, it is not safe to continue to operate on the buffer.
                //
                // See https://github.com/netty/netty/issues/1664
                if (ctx.isRemoved()) {
                    break;
                }

                if (outSize == out.size()) {    // out list没变，且Bytebuf的可读字节没变，比如读指针没变，说明数据不全，退出循环，不处理
                    if (oldInputLength == in.readableBytes()) {
                        break;
                    } else {
                        continue;
                    }
                }

                // 如果out list变化了，但ByteBuf的可读字节没变，抛出异常。
                // 用户实现decode方法时不对。不符合Netty的方式。解析出了msg，ByteBuf的读索引应该相应移动的。
                if (oldInputLength == in.readableBytes()) {
                    throw new DecoderException(  
                            StringUtil.simpleClassName(getClass()) +".decode() did not read anything but decoded a message.");
                }

                if (isSingleDecode()) {
                    break;
                }
            } 
    }

```


# LengthFieldBasedFrameDecoder
在RPC协议中，大的数据会分成多个包来发送，一般的协议会在消息头中定义长度字段来标识消息的总长度。这时候就可以使用LengthFieldBasedFrameDecoder来
进行合并，直到读到了完整的消息，才交给后面的handler处理。同时，这个decoder也可以对消息进行一些处理，比如丢弃掉其中的一些字节。
LengthFieldBasedFrameDecoder的作用:
1. 合并多个包的数据为一个完整的数据
2. 可以丢弃掉一些字节

|             arg            |                         desc                          | 
| -------------------------- | ----------------------------------------------------- |
| `lengthFieldOffset`        | 长度字段的偏移量                                         |
| `lengthFieldLength`        | 长度字段占的字节数                                       |
| `lengthAdjustment`         |                                                        |
| `initialBytesToStrip`      |                                                       |



