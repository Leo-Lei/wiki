---
layout: post
title: iTerm
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
description: The tutoria will describe the useage of Linux.
---

# iTerm常用命令    

|             Command              |                          Description                              |
| -------------------------------- | ----------------------------------------------------------------- |
| `Cmd` + `T`                      | 打开新tab                                                          |
| `Cmd` + `N`                      | 打开新窗口                                                          |
| `Cmd` + `W`                      | 关闭当前命令行                                                       |
| `Cmd` + `左箭头`                  | 切换到左边的tab                                                      |
| `Cmd` + `右箭头`                  | 切换到右边的tab                                                      |


# iTerm profile
登陆机器A
```bash
#!/usr/bin/expect -f

set port 22
set user root
set host 192.196.100.100
set timeout -1

spawn ssh $user@$host
interact
expect eof
```

登陆机器A，然后在A上登陆到机器B
```bash
#!/usr/bin/expect -f

set port 22
set user root
set host 100.200.200.100
set timeout -1

spawn ssh $user@$host
expect "*aliyun*"
send "ssh root@172.16.128.100\r"
interact
#expect "*@*"
```

