---
layout: post
title: Nginx
date: 2016-10-03 15:50:00
tags:
- Atom
categories: Text Editor
description: Nginx
---

# Yum安装nginx
yum默认源中是没有nginx的，需要我们自己添加nginx的yum源。创建文件`/etc/yum.repos.d/nginx.repo`。文件内容如下：
CentOS:
```bash
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=0
enabled=1
```
RHEL:
```bash
[nginx]
name=nginx repo
baseurl=http://nginx.org/packages/rhel/$releasever/$basearch/
gpgcheck=0
enabled=1
```
然后就可以使用命令来安装nginx了
```bash
yum install nginx
```


# nginx命令

|          command         |                                   |
| ------------------------ | --------------------------------- |
| `nginx`                  | 启动nginx                          |
| `nginx -s stop`          | 停止nginx                          |
| `nginx -s reload`        | reload配置                         |

# nginx配置
如果是用yum安装的nginx，配置文件目录在`/etc/nginx`。主配置文件是`/etc/nginx/nginx.conf`，该文件会include`/etc/nginx/conf.d`目录中以`conf`结尾的文件。
* `/etc/nginx/nginx.conf`:nginx全部配置。    
* `/etc/nginx/conf.d/*.conf`: 每一个server的配置。    

`/etc/nginx/nginx.conf`
```bash
user  nginx;
worker_processes  1;

error_log  /var/log/nginx/error.log warn;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    log_format  app   '"$remote_addr" "$remote_user" "[$time_local]" "$request" "$status" "$body_bytes_sent" "$http_referer" "$http_user_agent" "$request_method $scheme://$host$request_uri" "$host" "$http_x_forwarded_for" "$request_time" "$remote_port" "$upstream_response_time" "$http_x_readtime" "$uri" "$upstream_status" "$upstream_addr"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;
}
```

使用upstream，forward到后台服务
```bash
upstream app-api {
        least_conn;
        server 172.31.10.10:8080;
}

server {
    listen       80;
    server_name  app-api.mycompany.com;

    access_log  /var/log/nginx/app-api.mycompany.com.access.log  app;

    location / {
        proxy_pass http://app-api;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
    }
}
```


```
#user  nobody;
worker_processes  1;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;

events {
    use epoll;
    worker_connections  102400;
}

stream {

    upstream lock {
        #least_conn;
        hash $remote_addr consistent;
        server lock.hellotech.com.cn:29121 weight=5;
        server 192.168.128.7:29121 max_fails=3 fail_timeout=30s;
        server unix:/tmp/lock3;
        #######server lock.hellotech.com.cn:29121 weight=5;
        #######server 192.168.128.8:29121 max_fails=3 fail_timeout=30s;
        #######server unix:/tmp/lock3;
    }

    server {
        listen 29121;
        proxy_connect_timeout 60s;
        proxy_timeout 120s;
        proxy_pass lock;
    }


   ####server {
   ####    #listen [::1]:29121;
   ####    listen 192.168.128.7:29121;
   ####    proxy_pass unix:/tmp/stream.socket;
   ####}
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    log_format  app   '"$remote_addr" "$remote_user" "[$time_local]" "$request" "$status" "$body_bytes_sent" "$http_referer" "$http_user_agent" "$request_method $scheme://$host$request_uri" "$host" "$http_x_forwarded_for" "$request_time" "$remote_port" "$upstream_response_time" "$http_x_readtime" "$uri" "$upstream_status" "$upstream_addr"';


    sendfile       on;
    tcp_nopush     on;
    tcp_nodelay    on;

    keepalive_timeout  60;
    open_file_cache max=102400 inactive=20s;
    open_file_cache_valid 30s;
    open_file_cache_min_uses 1;


    gzip               on;
    gzip_vary          on;
    gzip_comp_level    6;
    gzip_buffers       16 8k;
    gzip_min_length    4k;
    gzip_http_version  1.1;
    gzip_proxied       any;
    gzip_disable       "msie6"
    gzip_types         text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript application/javascript;

    ##cache##
    proxy_connect_timeout 5;
    proxy_read_timeout 60;
    proxy_send_timeout 5;
    proxy_buffer_size 16k;
    proxy_buffers 4 64k;
    proxy_busy_buffers_size 128k;
    proxy_temp_file_write_size 128k;
    proxy_temp_path /var/tmp/temp_dir;
    proxy_cache_path /var/tmp/cache levels=1:2 keys_zone=cache_one:200m inactive=1d max_size=30g;
    proxy_next_upstream error timeout invalid_header http_500 http_503;
    client_max_body_size 5m;
    #proxy_next_upstream http_502 http_504 http_404 error timeout invalid_header;
    ##cache end##

    fastcgi_cache_valid 200 302 1h;
    fastcgi_cache_valid 301 1d;
    fastcgi_cache_valid any 1m;

    upstream APP_bff {
        #sticky;
        least_conn;
        server 192.168.128.1:20202;
        server 192.168.128.2:20202;
    }

    upstream Console_bff {
        #sticky;
        least_conn;
        server 192.168.128.1:20203;
        server 192.168.128.2:20203;
    }

    upstream dubbo-admin {
        #sticky;
        least_conn;
        server 192.168.128.1:20200;
    }

    upstream dubbo-monitor {
        #sticky;
        least_conn;
        server 192.168.128.3:8283;
    }

    upstream disconf {
        server 192.168.128.2:20230;
    }

    server {
        listen       80;
        server_name  localhost;

	location / {
            root   html;
            index  index.html index.htm;
        }

	error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
     }

###### APP_bff server
    server {
      listen       80;
      server_name  app-bff.api.hellotech.com.cn app-api-test.hellotech.com;
      access_log   logs/core.access.log  app;

      location / {
        proxy_pass http://APP_bff;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
      }
    }

###### Console_bff server
    server {
      listen       80;
      server_name  console-bff.api.hellotech.com.cn console-api-test.hellotech.com;
      access_log   logs/auth.api.access.log  app;

      location / {
        proxy_pass http://Console_bff;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
      }
    }

########dubbo-admin
    server {
      listen       80;
      server_name  dubboadmin.hellotech.com.cn;
      access_log   logs/dubbo-admin.access.log  main;

      location / {
        proxy_pass http://dubbo-admin;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
      }
    }

########dubbo-monitor
    server {
      listen       80;
      server_name  dubbomonitor.hellotech.com.cn;
      access_log   logs/dubbo-monitor.access.log  main;

      location / {
        proxy_pass http://dubbo-monitor;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
      }
    }

###------------------------------------------------WEB------------------#########

#########static
###    server {
###      listen       80;
###      server_name  static.hellotech.com.cn;
###      access_log   logs/staticoss.access.log  main;
###
###      location / {
###        proxy_pass http://hello-test.oss-cn-shanghai-internal.aliyuncs.com;
###        proxy_set_header   Host             $host;
###        proxy_set_header   X-Real-IP        $remote_addr;
###        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
###      }
###    }

###console_web server
    server {
        listen       80;
        server_name  static-console.hellotech.com.cn;
        access_log   logs/console.access.log  main;

        location / {
            root   static/console;
            index  index.html;
        }
        location ~ .*\.(gif|jpg|png|flv|ico|swf)(.*) {
            root static/console;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 3d;
        }
        location ~ .*\.(htm|html|css|js)(.*) {
            root static/console;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 1d;
        }

    }

###console_web server
    server {
        listen       80;
        server_name  static-tools.hellotech.com.cn;
        access_log   logs/console.access.log  main;

        location / {
            root   static;
            index  index.html;
        }
   ##     location ~ .*\.(gif|jpg|png|flv|ico|swf)(.*) {
   ##         root static;
   ##         proxy_cache cache_one;
   ##         proxy_cache_valid 200 302 1h;
   ##         proxy_cache_valid 301 1d;
   ##         proxy_cache_valid any 1m;
   ##         expires 3d;
   ##     }
   ##     location ~ .*\.(htm|html|css|js)(.*) {
   ##         root static;
   ##         proxy_cache cache_one;
   ##         proxy_cache_valid 200 302 1h;
   ##         proxy_cache_valid 301 1d;
   ##         proxy_cache_valid any 1m;
   ##         expires 1d;
   ##     }

    }




####invite_web server
    server {
        listen       80;
        server_name  static-invite.hellotech.com.cn;
        access_log   logs/invite.access.log  main;

        location / {
            root   static/invite;
            index  index.html;
        }
    ##    location ~ .*\.(gif|jpg|png|flv|ico|swf)(.*) {
    ##        root static/invite;
    ##        proxy_cache cache_one;
    ##        proxy_cache_valid 200 302 1h;
    ##        proxy_cache_valid 301 1d;
    ##        proxy_cache_valid any 1m;
    ##        expires 3d;
    ##    }
    ##    location ~ .*\.(htm|html|css|js)(.*) {
    ##        root static/invite;
    ##        proxy_cache cache_one;
    ##        proxy_cache_valid 200 302 1h;
    ##        proxy_cache_valid 301 1d;
    ##        proxy_cache_valid any 1m;
    ##        expires 1d;
    ##    }

    }


#############disconf-test
    server {
        listen       80;
        server_name  disconf-test.hellotech.com.cn;
        access_log logs/disconf_access.log;
        error_log logs/disconf_error.log;

        location / {
        root /opt/disconf/war/html;
        if ($query_string) {
            expires max;
        }
    }

    location ~ ^/(api|export) {
        proxy_pass_header Server;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Scheme $scheme;
        proxy_pass http://disconf;
    }
   }
}

```


# Resources     
[http://www.cnblogs.com/Gukw/archive/2012/05/13/2498328.html](http://www.cnblogs.com/Gukw/archive/2012/05/13/2498328.html)     
[http://www.codepool.biz/how-to-configure-and-install-nginx-on-mac-os-x.html](http://www.codepool.biz/how-to-configure-and-install-nginx-on-mac-os-x.html)
