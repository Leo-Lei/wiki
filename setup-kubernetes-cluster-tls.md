---
layout: post
title: Setup Kubernetes Cluster - CA Certification and private secret
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
banner: http://ohaq3i4w3.bkt.clouddn.com/docker-01.png
---

# 创建CA证书和秘钥

3台机器上都需要证书，但是只会在其中一台机器上生成证书，然后所有机器都用这个证书。

# 192.168.5.100安装CFSSL
登录192.168.5.100机器

```bash
cd /root/local
$ wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64
$ chmod +x cfssl_linux-amd64
$ sudo mv cfssl_linux-amd64 /root/local/bin/cfssl

$ wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64
$ chmod +x cfssljson_linux-amd64
$ sudo mv cfssljson_linux-amd64 /root/local/bin/cfssljson

$ wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64
$ chmod +x cfssl-certinfo_linux-amd64
$ sudo mv cfssl-certinfo_linux-amd64 /root/local/bin/cfssl-certinfo

$ export PATH=/root/local/bin:$PATH
$ mkdir -p /root/local/ssl
$ cd /root/local/ssl
$ cfssl print-defaults config > config.json   # 可以不执行，只是生成一个默认的模板 
$ cfssl print-defaults csr > csr.json         # 可以不执行，只是生成一个默认的模板
```
# 192.168.5.100创建CA（Certification Authority）
### 创建CA配置文件：
```bash
cd /root/local/ssl
touch ca-config.json
vim ca-config.json
```

ca-config.json
```json
{
  "signing": {
    "default": {
      "expiry": "8760h"
    },
    "profiles": {
      "kubernetes": {
        "usages": [
            "signing",
            "key encipherment",
            "server auth",
            "client auth"
        ],
        "expiry": "8760h"
      }
    }
  }
}
```
### 创建CA证书签名请求
```bash
cd /root/local/ssl
touch ca-csr.json
vim ca-csr.json
```
ca-csr.json
```json
{
  "CN": "kubernetes",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "BeiJing",
      "L": "BeiJing",
      "O": "k8s",
      "OU": "System"
    }
  ]
}
```
### 生成CA证书和私钥
```bash
$ cfssl gencert -initca ca-csr.json | cfssljson -bare ca
$ ls ca*
ca-config.json ca.csr ca-csr.json ca-key.pem ca.pem
```

# 分发证书
将生成的CA证书，秘钥文件，配置文件拷贝到所有机器的`/etc/kubernetes/ssl`目录下        
1. 将证书拷贝到192.168.5.100的`/etc/kubernetes/ssl`目录
```bash
mkdir -p /etc/kubernetes/ssl
cd /root/local/ssl
cp ca* /etc/kubernetes/ssl
```
2. 将证书拷贝到192.168.5.101的`/etc/kubernetes/ssl`目录
```bash
cd /root/local/ssl
cp ca* root@192.168.5.101:/etc/kubernetes/ssl
```
3. 将证书拷贝到192.168.5.102的`/etc/kubernetes/ssl`目录
```bash
cd /root/local/ssl
cp ca* root@192.168.5.102:/etc/kubernetes/ssl
```

> 注意： 在其中的某一台机器上生成证书和私钥，然后将这些文件拷贝到所有机器上去，千万记得所有的机器都使用该证书




