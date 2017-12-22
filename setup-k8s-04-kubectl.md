---
layout: post
title: Setup Kubernetes Cluster - Install Kubectl
date: 2017-07-08 11:10:00
tags:
- docker
categories: Java
---

# 安装kubectl命令行工具

* k8s的Master和Node节点都需要安装kubectl
* k8s的Master和Node节点共用一份`.kube/config`文件
* 现在192.168.5.100上安装kubectl，生成`.kube/config`文件，并将`.kube/config`文件拷贝到192.168.5.101和192.168.5.102上。在192.168.5.101和192.168.5.102上
只用安装kubectl，不用再生成`.kube/config`文件了。

# 使用的变量
```bash
export MASTER_IP=192.168.5.100 # 替换为 kubernetes master 集群任一机器 IP,本文档中只有一个Master 192.168.5.100
export KUBE_APISERVER="https://${MASTER_IP}:6443"
```

# 192.168.5.100下载kubectl
```bash
$ wget https://dl.k8s.io/v1.6.2/kubernetes-client-linux-amd64.tar.gz
$ tar -xzvf kubernetes-client-linux-amd64.tar.gz
$ sudo cp kubernetes/client/bin/kube* /root/local/bin/
$ chmod a+x /root/local/bin/kube*
$ export PATH=/root/local/bin:$PATH
```

* 下载`kubernetes-client-linux-amd64.tar.gz`时可能比较慢

# 192.168.5.100创建admin证书
```bash
$ cat admin-csr.json
{
  "CN": "admin",
  "hosts": [],
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "C": "CN",
      "ST": "BeiJing",
      "L": "BeiJing",
      "O": "system:masters",
      "OU": "System"
    }
  ]
}
```
* 后续 kube-apiserver 使用 RBAC 对客户端(如 kubelet、kube-proxy、Pod)请求进行授权；
* kube-apiserver 预定义了一些 RBAC 使用的 RoleBindings，如 cluster-admin 将 Group system:masters 与 Role cluster-admin 绑定，该 Role 授予了调用kube-apiserver 所有 API的权限；
* O 指定该证书的 Group 为 system:masters，kubelet 使用该证书访问 kube-apiserver 时 ，由于证书被 CA 签名，所以认证通过，同时由于证书用户组为经过预授权的 system:masters，所以被授予访问所有 API 的权限；
* hosts 属性值为空列表；

生成admin证书和私钥
```bash
$ cfssl gencert -ca=/etc/kubernetes/ssl/ca.pem \
  -ca-key=/etc/kubernetes/ssl/ca-key.pem \
  -config=/etc/kubernetes/ssl/ca-config.json \
  -profile=kubernetes admin-csr.json | cfssljson -bare admin
$ ls admin*
admin.csr  admin-csr.json  admin-key.pem  admin.pem
$ sudo mv admin*.pem /etc/kubernetes/ssl/
$ rm admin.csr admin-csr.json
```

# 192.168.5.100创建kubectl kubeconfig文件
```bash
$ # 设置集群参数
$ kubectl config set-cluster kubernetes \
  --certificate-authority=/etc/kubernetes/ssl/ca.pem \
  --embed-certs=true \
  --server=${KUBE_APISERVER}
$ # 设置客户端认证参数
$ kubectl config set-credentials admin \
  --client-certificate=/etc/kubernetes/ssl/admin.pem \
  --embed-certs=true \
  --client-key=/etc/kubernetes/ssl/admin-key.pem
$ # 设置上下文参数
$ kubectl config set-context kubernetes \
  --cluster=kubernetes \
  --user=admin
$ # 设置默认上下文
$ kubectl config use-context kubernetes
```
* 生成的 kubeconfig 被保存到 ~/.kube/config 文件；
* admin.pem 证书 O 字段值为 system:masters，kube-apiserver 预定义的 RoleBinding cluster-admin 将 Group system:masters 与 Role cluster-admin 绑定，该 Role 授予了调用kube-apiserver 相关 API 的权限；

# 分发kubeconfig文件
将`/root/.kube/config`文件拷贝到其他运行kubectl命令的机器的`/root/.kube`目录下。

# 192.168.5.101安装kubectl
# 192.168.5.102安装kubectl


