---
title: ShadowSocks
date: 2017-06-09 10:22:30
categories:
- Music
tags:
- Music
---

# 安装ShadowSocks服务端
以此输入以下命令来安装Shadowsocks
```bash
sudo yum install -y python-setuptools
sudo easy_install pip
sudo pip install shadowsocks
```
仔细看看有没有错误，如无错误即可继续
配置Shadowsocks
在终端输入
```bash
ssserver -h
```
查看Shadowsocks的帮助，适合进阶玩家。
如果提示没有ssserver这个命令，可以通过
which ssserver
查看Shadowsocks的路径，一般是在/usr/local/bin目录下，我们只需要把/usr/local/bin加入到/etc/profile文件中即可。
接下来创建shadowsocks目录，用于存放配置文件
mkdir /etc/shadowsocks
创建其配置文件
sudo vim /etc/shadowsocks/config.json
配置文件的内容如下
```json
{
"server":"0.0.0.0",
"server_port":443,
"local_address":"127.0.0.1",
"local_port":1080,
"password":"celerysoft.github.io",
"timeout":300,
"method":"aes-256-cfb",
"fast_open":false,
"workers": 1
}
```
配置文件说明
 
server
服务端监听地址(IPv4或IPv6)
server_port
服务端端口，一般为443
local_address
本地监听地址，缺省为127.0.0.1
local_port
本地监听端口，一般为1080
password
用以加密的密匙
timeout
超时时间（秒）
method
加密方法，默认为aes-256-cfb，更多请查阅Encryption
fast_open
是否启用TCP-Fast-Open，true或者false
workers
worker数量，如果不理解含义请不要改（这个只在Unix和Linux下有用）
启动Shadowsocks服务器
依次输入以下命令来启动Shadowsocks
sudo ssserver -c /etc/shadowsocks/config.json -d start
如果想停止Shadowsocks服务，可以这样停止
sudo ssserver -c /etc/shadowsocks/config.json -d stop
如果更改了Shadowsocks的配置文件，可以通过restart命令来重启Shadowsocks服务
sudo ssserver -c /etc/shadowsocks/config.json -d restart
设置Shadowsocks开机启动
服务器运行久了，偶尔需要重启一下，重启时每次都要手动启动hadowsocks的话就太麻烦了，可以将其加到开机启动项。
sudo vi /etc/rc.local
将带有ssserver内容的行删除，最后加入
sudo ssserver -c /etc/shadowsocks.json -d start
然后保存退出，这样，服务器上的操作就算完成了，接下来改对本地计算机进行操作了。




# Shadowsocks的Mac客户端下载

https://github.com/shadowsocks/ShadowsocksX-NG



# Chrome浏览器插件SwitchOmega
下载SwitchOmega插件。

# GFW规则
https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt

