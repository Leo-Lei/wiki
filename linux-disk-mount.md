---
layout: post
title: Linux Disk Mount
date: 2017-06-14 11:30:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the useage of Linux.
---

# 流程
1. 磁盘分区
2. 分区格式化
3. 将分区挂载到某个目录

# 查看是否已识别磁盘
通过`fdisk -l`命令来查看磁盘是否已识别。
```bash
fdisk -l

磁盘 /dev/sda：64.4 GB, 64424509440 字节，125829120 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：dos
磁盘标识符：0x000a47ad

   设备 Boot      Start         End      Blocks   Id  System
/dev/sda1   *        2048     1026047      512000   83  Linux
/dev/sda2         1026048   125829119    62401536   8e  Linux LVM

磁盘 /dev/sdb：1649.3 GB, 1649267441664 字节，3221225472 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
```
可以看到，有2块磁盘。/dev/sda和/dev/sdb。/dev/sda已经分好区了，但是/dev/sdb还没有分好区，需要对其进行分区。

# 使用fdisk命令进行建立分区
```bash
fdisk /dev/sdb
```
fdisk命令如下：
```bash
命令(输入 m 获取帮助)：m
命令操作
   a   toggle a bootable flag
   b   edit bsd disklabel
   c   toggle the dos compatibility flag
   d   delete a partition
   g   create a new empty GPT partition table
   G   create an IRIX (SGI) partition table
   l   list known partition types
   m   print this menu
   n   add a new partition
   o   create a new empty DOS partition table
   p   print the partition table
   q   quit without saving changes
   s   create a new empty Sun disklabel
   t   change a partition's system id
   u   change display/entry units
   v   verify the partition table
   w   write table to disk and exit
   x   extra functionality (experts only)
```
# 输入`n`新建分区(只有一个主分区)
```bash
命令(输入 m 获取帮助)：n
Partition type:
   p   primary (0 primary, 0 extended, 4 free)
   e   extended
Select (default p): 
Using default response p
分区号 (1-4，默认 1)：
起始 扇区 (2048-3221225471，默认为 2048)：
将使用默认值 2048
Last 扇区, +扇区 or +size{K,M,G} (2048-3221225471，默认为 3221225471)：
将使用默认值 3221225471
分区 1 已设置为 Linux 类型，大小设为 1.5 TiB
```
# 保存分区
```bash
命令(输入 m 获取帮助)：w
The partition table has been altered!

Calling ioctl() to re-read partition table.
正在同步磁盘。
```
# 使用fdisk查看，分区已经有了
```bash
磁盘 /dev/sdb：1649.3 GB, 1649267441664 字节，3221225472 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：dos
磁盘标识符：0xe0bc0098

   设备 Boot      Start         End      Blocks   Id  System
/dev/sdb1            2048  3221225471  1610611712   83  Linux
```
# 格式化分区
```bash
[root@localhost home]# mkfs.xfs -f /dev/sd
sda   sda1  sda2  sdb   sdb1  
[root@localhost home]# mkfs.xfs -f /dev/sdb1
meta-data=/dev/sdb1              isize=256    agcount=4, agsize=100663232 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=0        finobt=0
data     =                       bsize=4096   blocks=402652928, imaxpct=5
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=0
log      =internal log           bsize=4096   blocks=196607, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
```
# 选择一个挂载点，将磁盘分区挂载到该目录
```bash
mount /dev/sdb1 /opt/data/
```
# 查看挂载是否成功
```bash
df -TH /opt/data/
文件系统       类型  容量  已用  可用 已用% 挂载点
/dev/sdb1      xfs   1.7T   34M  1.7T    1% /opt/data
```
# 在/etc/fstab中加入配置，让系统启动后自动挂载
```bash
vim /etc/fstab
```
加入一行记录:
```bash
#
# /etc/fstab
# Created by anaconda on Mon May  8 03:33:12 2017
#
# Accessible filesystems, by reference, are maintained under '/dev/disk'
# See man pages fstab(5), findfs(8), mount(8) and/or blkid(8) for more info
#
UUID=41b13286-90d7-4c0f-9afb-ba92382eda59 /                       ext4    defaults        1 1
/dev/vdb1  /opt/data  xfs   defaults   0  0
```









