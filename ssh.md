---
layout: post
title: SSH
date: 2016-11-16 14:34:00
tags:
- Java
categories: Java
description: SSH command.
---


# 生成SSH的key

```bash
ssh-keygen -t rsa -C "your_email@example.com"
```
生成id_rsa和id_rsa.pub两个文件，如下面代码所示

# SSH口令登陆    
1. 远程主机收到用户的请求，将主机自己的公钥发给用户。          
2. 用户使用这个公钥，将登陆密码加密后，发送回来。      
3. 远程主机用自己的私钥揭秘登陆密码，如果密码正确就同意用户登陆。    

# SSH key登陆
1. 用户将自己的公钥储存在远程主机上。是公钥，不是私钥。    
2. 登陆时远程主机向用户发送一段随机的字符串。    
3. 用户用自己的私钥揭秘后，再发回来。    
4. 远程主机用事先储存的用户的公钥来解密，如果成功就证明用户是可信的。直接登陆shell，不再要求密码。    

# 添加key到ssh-agent
如果一台server的authorized_keys中有公钥public_key,我们有该公钥对应的private_key。可以通过ssh-add命令，将该私钥添加到本机的ssh-agent中。就可以登陆远程server了。    
```
ssh-add ~/.ssh/id_rsa
```
注意:添加到ssh-agent的是私钥。

如果使用 ssh-add ~/.ssh/id_rsa的时候报如下错误，则需要先运行一下 ssh-agent bash 命令后再执行 ssh-add ...等命令    
```bash
Could not open a connection to your authentication agent.
```
# authorized_keys文件
如果需要使用key登陆，远程主机需要将用户的公钥，保存在登陆后的用户主目录的`$HOME/.ssh/authorized_keys`文件中。公钥就是一段字符串。authorized_keys文件也是一个文本文件。只需要将公钥这段字符串加入到authorized_keys文件中就可以了。一个key一行。    
注意：    
1.`.ssh/authorized_keys`文件权限必须是600(待验证，网上都是这么说的,但我的权限是644，`rw-r--r--`好像也可以生效。。。)

# 重启sshd服务
修改了authorized_keys文件后，需要重启sshd服务来让更改生效。    
```
service sshd restart
```
# 禁止密码登陆
只允许使用key登陆，禁止密码登陆。        
修改`/etc/ssh/sshd_config`,找到下面这行，应该是在文件的最后几行：
```
PasswordAuthentication yes
```
修改为
```
PasswordAuthentication no
```
保存后重启sshd服务
```
service sshd restart
```

# 生产环境SSH安全
生产环境为了安全，只能让有限的人，通过某种可控的途径来登陆线上服务器。

1. 所有机器禁止密码登陆，只能使用key登陆。
2. 所有机器通过跳板机来登陆。所有机器上只放置一个key，那就是跳板机的key。
3. 所有机器的sshd端口只开放给跳板机。
4. 跳板机要保证高可用性。不然跳板机一旦挂掉，所有的机器都登陆不了啦。所以需要把跳板机的key备份一下，备份一下私钥和公钥。这个私钥只能是管理员自己知道。防止跳板机挂机了。这时候可以在本机上`ssh-add ~/.ssh/id_rsa`来导入跳板机的私钥。这样就可以登陆线上机器了。





# 常见SSH登陆错误
```
# 服务器只能使用key登陆，但是本机的private key不对。
Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
```
