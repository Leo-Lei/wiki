---
layout: post
title: YAML
date: 2017-04-06 11:10:00
tags:
- docker
categories: Java
description: docker
---


# YAML

```yaml
# Menus
menu:
    Home: /
    # Delete this row if you don't want categories in your header nav bar
    Categories:
    About: /about/index.html

# Customize
customize:
    logo:
        width: 165
        height: 60
        url: images/logo-header.png
    theme_color: '#006bde'
    highlight: androidstudio
    sidebar: left # sidebar position, options: left, right
    thumbnail: true # enable posts thumbnail, options: true, false
    favicon: # path to favicon
    social_links: # for more icons, please see http://fontawesome.io/icons/#brand
        twitter: /
        facebook: /
        google-plus: /
        github: https://github.com/ppoffice/hexo-theme-hueman
        weibo: /
        rss: /

# Widgets
widgets:
    - recent_posts
    - category
    - archive
    - tag
    - tagcloud
    - links

# Search
search:
    insight: true # you need to install `hexo-generator-json-content` before using Insight Search
    swiftype: # enter swiftype install key here
    baidu: false # you need to disable other search engines to use Baidu search, options: true, false

# Comment
comment:
    disqus: hexo-theme-hueman # enter disqus shortname here
    duoshuo: # enter duoshuo shortname here
    youyan: # enter youyan uid here
    facebook: # enter true to enable
    isso: # enter the domain name of your own comment isso server eg. comments.example.com

# Share
share: default # options: jiathis, bdshare, addtoany, default

# Plugins
plugins:
    lightgallery: true # options: true, false
    justifiedgallery: true # options: true, false
    google_analytics: # enter the tracking ID for your Google Analytics
    baidu_analytics: # enter Baidu Analytics hash key
    mathjax: false # options: true, false

# Miscellaneous
miscellaneous:
    open_graph: # see http://ogp.me
        fb_app_id:
        fb_admins:
        twitter_id:
        google_plus:
    links:
        Hexo: http://hexo.io
        
dockerfile:
    zookeeper:
        download_url: http://ftp.jaist.ac.jp/pub/apache/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz
    disconf:
        download_url: https://github.com/Leo-Lei/disconf/archive/2.6.36.1.zip
        archive_top_dir: disconf-2.6.36.1
        host: 172.31.10.10

docker-registry: 192.168.10.10:5000

mysql:
    initial_users:
        - user: qibeibike
          password: qibeibike
        - user: admin
          password: admin

rds-mysql-sync:
    rds:
        - name: rds1
          backup_url: www.aliyun.com/foo.zip
          server_id: 100
          host: aaa.mysql.rds.aliyuncs.com:3306
          user: admin
          password: admin

        - name: rds2
          backup_url: www.aliyun.com/bar.zip
          server_id: 200
          host: bbb.mysql.rds.aliyuncs.com:3306
          user: admin
          password: admin
    merge-mysql:
        server_id: 300

```






# ---
在单一文件中`---`可以区分多个文件


