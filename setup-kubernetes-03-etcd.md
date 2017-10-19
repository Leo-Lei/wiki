---
layout: post
title: Setup Kubernetes Cluster - Etcd
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
description: docker
---

# 部署高可用etcd集群

Kubernetes使用etcd来存储数据。      

|          IP        |     Node              |
| ------------------ | --------------------- |
| 192.168.5.100      | etcd-host0            |
| 192.168.5.101      | etcd-host1            |
| 192.168.5.102      | etcd-host2            |

# 192.168.5.100配置环境变量
使用的环境变量如下
```
export NODE_NAME=etcd-host0 # 当前部署的机器名称(随便定义，只要能区分不同机器即可)
export NODE_IP=192.168.5.100 # 当前部署的机器 IP
export NODE_IPS="192.168.5.100 192.168.5.101 192.168.5.102" # etcd 集群所有机器 IP
# etcd 集群间通信的IP和端口
export ETCD_NODES=etcd-host0=https://192.168.5.100:2380,etcd-host1=https://192.168.5.101:2380,etcd-host2=https://192.168.5.102:2380
# 导入用到的其它全局变量：ETCD_ENDPOINTS、FLANNEL_ETCD_PREFIX、CLUSTER_CIDR
source /root/local/bin/environment.sh
```
<!-- more -->

# 192.168.5.100下载etcd二进制文件
到 https://github.com/coreos/etcd/releases 页面下载最新版本的二进制文件:        
```bash
$ wget https://github.com/coreos/etcd/releases/download/v3.1.6/etcd-v3.1.6-linux-amd64.tar.gz
$ tar -xvf etcd-v3.1.6-linux-amd64.tar.gz
$ sudo mv etcd-v3.1.6-linux-amd64/etcd* /root/local/bin
```

# 192.168.5.100创建TLS密钥和证书
创建etcd证书签名请求
```bash
cd /root/local/ssl
touch etcd-csr.json
vim etcd-csr.json
```
etcd-csr.json
```json
{
  "CN": "etcd",
  "hosts": [
    "127.0.0.1",
    "192.168.5.100"        
  ],
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

> 在192.168.5.101和192.168.5.102机器上，需要把hosts中的192.168.5.100替换成对应机器的IP 
生成etcd证书和私钥
```bash
$ cfssl gencert -ca=/etc/kubernetes/ssl/ca.pem \
  -ca-key=/etc/kubernetes/ssl/ca-key.pem \
  -config=/etc/kubernetes/ssl/ca-config.json \
  -profile=kubernetes etcd-csr.json | cfssljson -bare etcd
$ ls etcd*
etcd.csr  etcd-csr.json  etcd-key.pem etcd.pem
$ sudo mkdir -p /etc/etcd/ssl
$ sudo mv etcd*.pem /etc/etcd/ssl
$ rm etcd.csr  etcd-csr.json
```

# 192.168.5.100创建etcd的systemd unit文件
```bash
sudo mkdir -p /var/lib/etcd  # 必须先创建工作目录
$ cat > etcd.service <<EOF
[Unit]
Description=Etcd Server
After=network.target
After=network-online.target
Wants=network-online.target
Documentation=https://github.com/coreos

[Service]
Type=notify
WorkingDirectory=/var/lib/etcd/
ExecStart=/root/local/bin/etcd \\
  --name=${NODE_NAME} \\
  --cert-file=/etc/etcd/ssl/etcd.pem \\
  --key-file=/etc/etcd/ssl/etcd-key.pem \\
  --peer-cert-file=/etc/etcd/ssl/etcd.pem \\
  --peer-key-file=/etc/etcd/ssl/etcd-key.pem \\
  --trusted-ca-file=/etc/kubernetes/ssl/ca.pem \\
  --peer-trusted-ca-file=/etc/kubernetes/ssl/ca.pem \\
  --initial-advertise-peer-urls=https://${NODE_IP}:2380 \\
  --listen-peer-urls=https://${NODE_IP}:2380 \\
  --listen-client-urls=https://${NODE_IP}:2379,http://127.0.0.1:2379 \\
  --advertise-client-urls=https://${NODE_IP}:2379 \\
  --initial-cluster-token=etcd-cluster-0 \\
  --initial-cluster=${ETCD_NODES} \\
  --initial-cluster-state=new \\
  --data-dir=/var/lib/etcd
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF
```

# 192.168.5.100启动etcd服务
```bash
$ sudo mv etcd.service /etc/systemd/system/
$ sudo systemctl daemon-reload
$ sudo systemctl enable etcd
$ sudo systemctl start etcd
$ systemctl status etcd
```

* 最先启动的etcd进程会卡住一段时间，等待其他节点上的etcd进程加入集群，这是正常现象。    


# 验证etcd服务
在任意一个etcd节点上执行命令
```bash
for ip in ${NODE_IPS}; do
  ETCDCTL_API=3 /root/local/bin/etcdctl \
  --endpoints=https://${ip}:2379  \
  --cacert=/etc/kubernetes/ssl/ca.pem \
  --cert=/etc/etcd/ssl/etcd.pem \
  --key=/etc/etcd/ssl/etcd-key.pem \
  endpoint health; done
```
预期输入结果是
```text
2017-09-27 17:52:33.917986 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
https://192.168.5.100:2379 is healthy: successfully committed proposal: took = 2.021248ms
2017-09-27 17:52:33.955574 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
https://192.168.5.101:2379 is healthy: successfully committed proposal: took = 1.777773ms
2017-09-27 17:52:33.996138 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
https://192.168.5.102:2379 is healthy: successfully committed proposal: took = 2.413539ms
```
