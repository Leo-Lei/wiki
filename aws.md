---
layout: post
title: Docker
date: 2016-07-15 11:10:00
tags:
- docker
categories: Java
description: docker
---


# AWS的服务

|       name      |                 Desc                       |    阿里云等价物   |
| --------------- | ------------------------------------------ | --------------- |
| `EC2`           | EC2=Elastic Compute Cloud                  | `ECS`           |
| `S3`            |                                            | `OSS`           |
| `CloudFront`    |                                            | `CDN`           |
| `RDS`           | 关系型数据库                                 | `RDS`           |
| `DynamoDB`      | 结构化数据库                                 | `OTS`           |
| `ElasticCache`  | 缓存                                        | `OCS`           |
| `VPC`           | Virtual private cloud                      | `VPC`           |
| `Route53`       |                                            | `云解析`         |
| `Cloud Watch`   |                                            | `云监控`         |
| `SQS`           |                                            | `MQS`           |         

# SSH连接EC2
创建了一个新的EC2实例后，会要求生成一个密钥对。AWS会提示你下载私钥，公钥是直接保存在EC2机器上的。比如下载的文件是`aws-key.pem`
```bash
cp aws-key.pem ~/.ssh
cd ~/.ssh
ssh-add aws-key.pem
```


登陆EC2
```bash
ssh ec2-user@12.34.56.78
```






# EC2使用root登陆
AWS为了安全，不能直接使用root登陆。会出现***Please login as the user "ec2-user" rather than the user "root"***的错误提示。    
1. 创建EC2实例时，创建key,将该key添加到ssh agent中
```bash
cp aws-key.pem ~/.ssh
cd ~/.ssh
ssh-add aws-key.pem
```
2. 使用ec2-user用户SSH到EC2
```
ssh ec2-user@12.34.56.78
```

3. 修改root用户密码
```bash
sudo passwd root
```
shell会提示输入两次新的root密码。密码有一定的要求的，要包含大小写字母，数字，特殊字符，还不能是简单的英文单词。下面是一个有效的密码
```bash
Passw0rd@123
```

4. 切换到root账号
```bash
su - root
```
输入刚才设定的root密码

5. 修改`/etc/ssh/sshd_config`文件
```bash
vi /etc/ssh/sshd_config
```
将下面这行的注释去掉
```bash
#PermitRootLogin yes
```
改成
```bash
PermitRootLogin yes
```

6. 重启ssh服务
```bash
service sshd restart
```

7. 修改`/root/.ssh/authorized_keys`文件
把每一行`ssh-rsa`前面的命令注释掉。一般是以`no-port-forwarding`开始的。并且将ssh-rsa另起一行。可以添加其他的public key。
```bash
no-port-forwarding,no-agent-forwarding,no-X11-forwarding,command="echo 'Please login as the user \"ec2-user\" rather than the user \"root\".';echo;sleep 10"
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ......CqX53nAU2UONIrBC2d8QmZmMIL aws-key

```
改成
```bash
#no-port-forwarding,no-agent-forwarding,no-X11-forwarding,command="echo 'Please login as the user \"ec2-user\" rather than the user \"root\".';echo;sleep 10"
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ......CqX53nAU2UONIrBC2d8QmZmMIL aws-key

ssh-rsa AAAAB3NzaC1yc2EAAAADAQABJJDJBl......H3NvhEtWBa3 leiwei@leiweideMacBook-Pro.local
```

8. 以root登陆
```bash
ssh root@12.34.56.78
```
# 查看AWS各个区域的连接ping
[http://www.cloudping.info](http://www.cloudping.info)

# AWS RDS




# AWS IAM







另外mysql有RDS，redis和memcached有ElastiCache，mongodb有cloudformation搭建三个节点的集群，或者可以用托管的dynamodb数据库。消息中间件有kinesis和sqs，消息队列。
