---
layout: post
title: rsyslog
date: 2016-11-11 15:20:00
tags:
- Java
categories: Java
description: rsyslog
---

# 安装rsyslog
推荐使用yum来安装rsyslog

1. 从官网的[rpm仓库](http://rpms.adiscon.com)下载`rsyslog.repo`文件。放到`/etc/yum.repos.d/`目录中。      
2. 运行命令`yum install rsyslog`

```bash
cd /etc/yum.repos.d/
wget http://rpms.adiscon.com/v8-stable/rsyslog.repo
yum install rsyslog
```

> 注意，rsyslog的各个版本之间兼容性不好，所以安装时要注意版本。

# 运行／停止 rsyslog
* 运行： `systemctl restart rsyslog`    
* 停止： `systemctl stop rsyslog`    
可以用`ps aux | grep rsyslog`来查看rsyslog进程是否已启动。        

# rsyslog配置文件        
rsyslog最主要的配置文件是:`/etc/rsyslog.conf`。在`/etc/rsyslog.conf`中会引用`/etc/rsyslog.d/*.conf`，即`/etc/rsyslog.d`目录下的所有以conf结尾的文件。







# sample
```ruby
input( type="imfile" PersistStateInterval="1000" Tag="tag_app"         File="/some/path/logs/app.log" )
input( type="imfile" PersistStateInterval="1000" Tag="tag_app"         File="/some/path/logs/app-error.log" )
input( type="imfile" PersistStateInterval="1000" Tag="tag_invite" readMode="1" File="/some/path/test.log" )

template( type="string" name="tn_app" string="%msg%")

# output
if ($syslogtag == "tag_app") then {
	 action(
		      type="omfwd"
		      Target="172.16.128.1"
	        Port="25839"
		      Protocol="udp"
		      template="tn_app"
      )
   action(
          type="omfile"
          dirCreateMode="0700"
          FileCreateMode="0644"
          File="/some/path/debug_rsyslog.txt"
          template="tn_app"
        )
	stop
}

```

# Resources
http://www.rsyslog.com
