---
layout: post
title: Nginx
date: 2016-10-03 15:50:00
tags:
- Atom
categories: Text Editor
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
### 将主配置和各个server的配置分开
如果是用yum安装的nginx，配置文件目录在`/etc/nginx`。主配置文件是`/etc/nginx/nginx.conf`，该文件会include`/etc/nginx/conf.d`目录中以`conf`结尾的文件。
* `/etc/nginx/nginx.conf`:nginx全部配置。    
* `/etc/nginx/conf.d/*.conf`: 每一个server的配置。    

`/etc/nginx/nginx.conf`

### Nginx主配置
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

### 转发到后台服务器，比如tomcat
`/etc/nginx/config.d/app-api.conf`:
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


```bash
#user  nobody;
worker_processes  2;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;

events {
    use epoll;
    worker_connections  3096;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    log_format	app   '"$remote_addr" "$remote_user" "[$time_local]" "$request" "$status" "$body_bytes_sent" "$http_referer" "$http_user_agent" "$request_method $scheme://$host$request_uri" "$host" "$http_x_forwarded_for" "$request_time" "$remote_port" "$upstream_response_time" "$http_x_readtime" "$uri" "$upstream_status" "$upstream_addr"';

    sendfile       on;
    tcp_nopush     on;
    tcp_nodelay    on;

    keepalive_timeout  60;
    server_tokens     off;

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
        server 172.16.128.10:20202;
    }

    upstream Console_bff {
        #sticky;
        least_conn;
        server 172.16.128.17:8080;
    }



##    server {
##        listen       80;
##        server_name  localhost;
##
##	location / {
##            root   html;
##            index  index.html index.htm;
##        }
##
##	error_page   500 502 503 504  /50x.html;
##        location = /50x.html {
##            root   html;
##        }
##     }
##
###### APP_bff server
    server {
      listen       80;
      server_name  app.api.hellotech.com app-api.hellotech.com;
      access_log   logs/app.api.hellotech.com.access.log  app;

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
      server_name  console.api.hellotech.com console-api.hellotech.com;
      access_log   logs/console.api.hellotech.com.access.log  app;

      location / {
        proxy_pass http://Console_bff;
        proxy_set_header   Host             $host;
        proxy_set_header   X-Real-IP        $remote_addr;
        proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
      }
    }

####------------------------Static-------WEB------------------#########

#######OSS;download
    server {
        listen       80;
        server_name  download.hellotech.com;
        access_log   logs/download.hellotech.com.access.log  app;

        location / {
            root   download;
            index  index.html;
    }
        location ~ .*\.(gif|jpg|png|flv|ico|swf)(.*) {
            root download;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 3d;
        }
        location ~ .*\.(htm|html|css|js)(.*) {
            root download;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 1d;
        }
   }

####Static
    server {
        listen       80;
        server_name  static-invite.hellotech.com;
        access_log   logs/static-invite.hellotech.com.access.log  app;

        location / {
            root   Static/invite;
            index  index.html;
    }
        location ~ .*\.(gif|jpg|png|flv|ico|swf)(.*) {
            root Static/invite;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 3d;
        }
        location ~ .*\.(htm|html|css|js)(.*) {
            root Static/invite;
            proxy_cache cache_one;
            proxy_cache_valid 200 302 1h;
            proxy_cache_valid 301 1d;
            proxy_cache_valid any 1m;
            expires 1d;
        }
   }
}

```

# 配置多个Server name
```bash
server {
        listen       80;
        server_name  mycompany.com *.mycompany.com;
        ......
}
```

# 配置静态web站点
```bash
server {
	#监听8089端口下的www.aabbccdd.com服务请求进行处理
        listen       8890;
        server_name  localhost;

        location / {
          # root /Users/leiwei/tmp/blog/public;
           root /Users/leiwei/workspace/github/blog/public;
            #默认请求转到root路径下的index.html页面。
	    index index.html;
        }
}
```

# Resources     
[http://www.cnblogs.com/Gukw/archive/2012/05/13/2498328.html](http://www.cnblogs.com/Gukw/archive/2012/05/13/2498328.html)     
[http://www.codepool.biz/how-to-configure-and-install-nginx-on-mac-os-x.html](http://www.codepool.biz/how-to-configure-and-install-nginx-on-mac-os-x.html)
