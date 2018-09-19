---
layout: post
title: Linux Commands
date: 2016-07-18 15:10:00
tags:
- Linux
categories: Linux
---

### 文件操作
|                           command                        |                                                        |
| -------------------------------------------------------- | ------------------------------------------------------ |
| `pwd`                                                    | print working directory                                | 
| <code>cat app.log &#124; less</code>                     | 分页查看                                                |            
| `cd <path/to/directory>`                                 | chande directory                                       | 

### CentOS／RedHat命令
|                           command                        |                                                        |
| -------------------------------------------------------- | ------------------------------------------------------ |
| `pwd`                                                    | print working directory                                | 
| <code>cat app.log &#124; less</code>                     | 分页查看                                                | 
| `cd <path/to/directory>`                                 | chande directory                                       | 

### Debina/Ubuntu命令
|                           command                        |                                                        |
| -------------------------------------------------------- | ------------------------------------------------------ |
| `apt-get install mysql-server`                           | software/package management                            | 
| `apt-get update`                                         | software/package management                            | 


### 系统命令
|                           command                        |                                                        |
| -------------------------------------------------------- | ------------------------------------------------------ |
| `apt-get install mysql-server`                           | software/package management                            | 
| `chmod <permission> <file_name>`                         | change permission of a file or directory               | 
| `chown tomcat:tomcat hello.log`                          | 修改hello.log文件的用户和用户组                           | 
| `chown -R tomcat:tomcat /home/doc`                       | 修改/home/doc及子目录的用户和用户组                        | 
| `cat /etc/passwd`                                        | 查看所有用户列表                                         |
| `curl http://www.google.com`                             | a http command tool                                   | 
| `curl -fSL https://github.com/hello.zip -o hello.zip`    | 下载文件                                               |
| `cp 1.txt 2.txt`                                         | 复制文件／文件夹                                         |
| `df -hl`                                                 | 查看磁盘空间                                             |
| `echo hello world`                                       | output some string to stdout                          |
| `find -name "*.php"`                                     | find file                                             |
| `ls <path/to/directory>`                                 | list the files and directories                         | 
| `vi <file_name>`                                         | open the vi(a text editor)                             | 
| `mkdir <dir_name>`                                       | create a new directory                                 |
| `rm <file_name>` `rm <directory_name>`                   | remove a file or directory                             | 
| `rm -rf <dir_name>`                                      | 强制删除文件或文件夹                                      | 
| `unzip file.zip -d destination_folder`                   | unzip a zip file to some folder                       | 
| `scp file user@dest_host:dest_folder`                    | secure copy file                                      | 
| `sudo -u my_name ls /home/comphope/hope`                 | super user do                                         | 
| `lsof -i:20880`                                          | 检查端口使用情况                                         |
| `kill <process_id>`                                      | 杀死进程                                               |
| `ssh admin@192.168.20.147`                               | SSH登陆到服务                                          |
| <code>ps aux &#124; grep rsyslog</code>                  | 查看系统进程                                           | 
| `tar -cvf hello.tar /foo/bar`                            | 将目录/foo/bar压缩为hello.tar文件                      |
| `tar -czvf hello.tar.gz /foo/bar`                        | 将目录/foo/bar压缩为hello.tar.gz文件                   |
| `tar -xvf hello.tar -C /foo/bar`                         | 将hello.tar文件解压到/foo/bar目录                      |
| `tar -xzvf hello.tar.gz -C /foo/bar`                     | 将hello.tar.gz文件解压到/foo/bar目录                   |
| <code>tail -n400 hello.log &#124; less</code>            | 分页查看hello.log文件的最后400行                        |
| `find / -name hello`                                     | 在整个目录下搜索匹配hello的文件或文件夹                   |
| `free -m`                                                | 查看系统内存使用情况，用M为单位显示                       |
| `top`                                                    | 查看系统cpu，内存使用情况                               |
| `rpm -lq mysql`                                          | 查看RPM安装的mysql的信息，如果不是RPM安装，提示没安装      |
| `rpm -qa`                                                | 查看RPM安装的包                                       |
| `su username`                                            | 切换用户，不切换环境变量                                |
| `su - username`                                          | 切换用户，还切换环境变量                                |
| `ssh-add ~/.ssh/id_rsa`                                  | 添加密钥到本地key agent                                |
| `uname -a` 或`uname --all`                               | 查看 本机所有信息                                      |
| `uname -s`或`uname --sysname`                            | 查看 操作系统名称                                      |
| `uname -m`或`uname --machine`                            | 查看 电脑类型                                         |
| `uname --help`                                           | 显示 帮助                                            |
| `wget -c http://www.mycompany.com/hello.zip -O /home/hello.zip` | 下载文件                                      |
| `wget --no-check-certificate https://www.hello.com/world.zip`   | 下载文件                                      |
| `who`                                                    | 查看当前系统的登陆用户                                 |
| `w`                                                      | 查看当前系统的登陆用户                                 |
| <code>yum list &#124; grep mysql</code>                  | 查看yum上可用的包                                     |
| `du -ah /opt`                                            | 查看文件夹及其子文件夹的大小                             |
| `du -sh /var/opt/*`                                      | 查看文件夹的大小，不包括子文件夹，避免输出信息太多           |
| `date -s "2008-08-08 12:00:00"`                          | 修改时间                                             |
| `ntpdate 1.cn.pool.ntp.org`                              | 从服务器同步系统时间                                   |
| `clock -w`                                               | 把系统时间写入CMOS                                    |
| `iptables -L -nv`                                        | 查看iptables配置                                     |
| `/usr/sbin/sestatus -v`                                  | 查看SELinux状态                                                         |
| `cmd > file`                                             | 重定向命令的stdout到一个文件                  |
| `cmd 1> file`                                            | 效果和 `cmd > file`一样                                               |
| `cmd 2> file`                                            | 重定向命令的stderr到一个文件                   |
| `cmd >> file`                                            | 追加命令的输出到一个文件                                         |
| `cmd 2>> file`                                           | 追加命令的stderr到一个文件                                         |
| `cmd &> file`                                            | 重定向命令的stdout和stderr到一个文件                            |
| `cmd \| tee file`                                         | 重定向命令的stdout到一个文件，并打印到控制台                 |
| `cmd  2>&1 \| tee -a  file`                               | 追加命令的stdout和stderr到一个文件，并打印到控制台       |
| `cmd1 \| cmd2`                                            | 重定向cmd1的stdout到cmd2的stdin. 效果等于cmd1 >> (cmd2)    |
| `cmd1 \|& cmd2`                                           | 重定向cmd1的stdout和stderr到cmd2的stdin。(bash 4.0+ only)    |
| `cmd1 2>&1 \| cmd2`                                       | 效果等于`cmd1 |& cmd2`. 在老版本的shell上使用                 |
| `ifconfig`                                                | 查看ip地址                                                  |
| `/sbin/ifconfig`                                          | `ifconfig`                                                 |
| `ip addr`                                                 | 查看ip地址                                                  |



### curl

|                   Commands                              |      Usage                         |
| ------------------------------------------------------- | ---------------------------------- |
| `curl -i www.sina.com`                                  | 显示头信息                           |
| `curl -X POST --data "data=xxx" example.com/form.cgi`   | POST                               |







# 时间函数

|                           command                        |                                                        |
| -------------------------------------------------------- | ------------------------------------------------------ |
| `date -d '2 days' +%Y-%m-%d`                             | 显示2天后的日期，格式为2017-02-13                          | 
| `date -d '2 days ago' +%Y-%m-%d`                         | 显示2天前的日期，格式为2017-02-13                          | 


## rm
`rm` command has some options:
* -r: Delete the directory recursively.
* -f: Delete the file or directory forcely, no matter the directory is empty or not.

So if you want to delete all the files and subdirectories in a folder, you can run the command `rm -rf <directory_path>`

## ls
`ls` command has following options:
* -l: Show the result as a list view.
* -a: Show all files, include hidden files.
* -t: Sort the files by date. The latest modified at the top.

## mkdir
mkdir dirname
mkdir dirname1 dirname2
mkdir -p dir1/dir2/dir3

## curl
1. `curl http://www.example.com`:    
Get the html document of *http://www.example.com*. It will display in the standard output.
2. `curl http://www.example.com > page.html`:    
Get the html document and save to a file page.html.

## find
`find <location> <filter-criteria> <search-term>`     
Note: Search sub directory recursively.               

|    command                                |      description                                  |  example                        |
| ----------------------------------------- | ------------------------------------------------- | ------------------------------- |
| `find`                                    | List all files in current and sub directories.    |  $ find<br>.<br>./abc.txt<br>./foo<br>./foo/hello.java| 
| `find test`                               | List all files in the test directory.             |  $ find test<br>test<br>test/abc.txt<br>test/foo<br>test/foo/hello.java| 
| `find test -name "abc.txt"`               | Search file abc.txt.                              |  $ find test -name "abc.txt"<br>test/abc.txt|
| `find test -name "*.java"`                | Search file by wildcards.                         |  $ find test -name "*.java"<br>test/hello.java<br>test/foo/world.java|
| `find test -name "*.java"`                | Search file by wildcards.                         |  $ find test -name "*.java"<br>test/hello.java<br>test/foo/world.java|
| `find test -iname "*.jAvA"`               | Ignore the case, i.e. case-insensitive            |  $ find test -iname "*.jAvA"<br>test/hello.java<br>test/foo/world.java|
| `find test -maxdepth 2 -name "*.java"`    | Limit depth of directory traversal                |  $ find test -maxdepth 2 -name "\*.java"<br>test/hello.java<br>test/foo/world.java<br><br>$ find test -maxdepth 1 -name "\*.java"<br>test/hello.java| 
| `find test -not -name "*.java"` <br> `find test ! -name "*.java"` | Search file not match pattern                     |  $ find test-not -iname "*.java"<br>test<br>test/foo<br>test/abc.txt|
| `find test -name 'abc*' -name '*.java'`   | Search start with 'abc' **and** end with 'java'                 |  $ find test -name 'abc\*' -name '\*.java'<br>test/abc.java<br>test/foo/abcdef.java|
| `find -name '*.php' -o -name '*.java'`    | Search end with 'php' **or** end with 'java'                 |  $ find -name '\*.php' -o -name '\*.java'<br>hello.php<br>foo/world.java|
| `find -type f -name 'abc*'` <br> `find -type d -name 'abc*'`   | Search only files or directories                 |  $ find -name 'abc\*'<br>abc<br>abc.txt<br><br>-- only files<br>$ find -type f -name 'abc\*'<br>abc.txt<br><br>-- only directories<br>$ find -type d -name 'abc*'<br>abc|
| `find dir1 dir2 -name '*java'`            | Search multiple directories                       |  $ find dir1 dir2 -name '*java' <br> dir1/hello.java<br>dir2/world.java |
| `find -user bob -name '*java'`            | Find files belonging to particular user           |  $ find -user bob -name '*java' <br>abc.java   |
| `find -group developer -name '*java'`     | Find files belonging to particular user           |  $ find -group developer -name '*java' <br>abc.java   |                                     
| `find -mtime -5`                          | Find files modified less than 5 days ago.         |                    |
| `find -mmin -5`                           | Find files modified less than 5 minutes ago.      |                    |
| `find -mmin +5`                           | Find files modified more than 5 minutes ago.      |                    |
| `find -mmin 5`                            | Find files modified exactly 5 minutes ago.        |                    |
| `find -cmin -5`                           | Find files **created** less than 5 minutes ago.   |                    |
| `find -size +50M -size -100M`             | Find files between 50M and 100M. |                     |

# cp

|               command               |                                                      |
| ----------------------------------- | ---------------------------------------------------- |
| `cp -r /dir/dir1/* /dir/dir2`         | 将文件夹dir1下的文件和子文件夹复制到dir2                 |
| `cp -r /dir/dir1/. /dir/dir2`         | 效果和`cp -r /dir/dir1/* /dir/dir2`一样              |
| `cp -r /dir/dir1/ /dir/dir2`         | 将文件夹dir1复制到dir2                                 |


## less

|  command |                          |
| -------- | ------------------------ |
| `q`      | 退出less模式              |
| `空格`   | 查看下一页                 |
| `b`     | 查看上一页                 |


## free
`free -m`查看系统内存使用情况，输出如下：
```
             total       used       free     shared    buffers     cached
Mem:          7824       7661        163         80        177       1584
-/+ buffers/cache:       5899       1924
Swap:            0          0          0
```
其中Mem这一行是从OS的角度来查看内存占用情况的。一共有7824M内存。用了7661M，还剩余163M。
在使用了的7661M内存中，有177M是OS用于buffers的，有1584M是OS用于cached的。这些用于buffers和cached的内存是操作系统管理的。而且这部分内存是很容易回收的。像Linux和Windows这些比较成熟的操作系统为了提高IO的读性能，都会缓存大量的数据。所以，cached通常会比较大。第三行的内存占用是从应用程序角度来看的。虽然实际物理内存占用了7824M，但这其中有177M是buffers的，1584M是cached的，这些是操作系统管理的，为了提高系统的IO性能。而且这部分内存可以很容易的被回收。所以，从应用程序的角度来看，系统被使用的内存没有7661M，只有5899M。还剩余了1924M。
