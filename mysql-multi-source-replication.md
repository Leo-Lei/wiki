---
layout: post
title: MySql多源复制
date: 2017-02-14 14:20:00
tags:
- Mysql
categories: 
- Mysql
description: MySql
---

# 概述
有2台阿里云的RDS实例，现在想做一些数据分析，希望将这2个RDS实例的数据都同步到一个自建的数据库实例中。
方案：
```
         RDS-1(阿里云RDS，5.6)                  RDS-2(阿里云RDS,5.6)
                  |                                     |
                  |                                     |
                  V                                     V
         Mysql-1(自建mysql，5.7)                Mysql-2(自建mysql，5.7)
                   \                                    /
                    \                                  /
                     \                                /
                      \                              /           
                       \                            /
                           Mysql-all(自建mysql，5.7)

```
下面的文档中，两台阿里云的机器命名为RDS-1和RDS-2，自建的两台用来同步RDS的机器命名为mysql-1和mysql-2。将数据汇总的机器命名为mysql-all。    
首先将2个RDS的数据同步到自建的数据库，通过mysql的master-slave机制来同步，但rds到mysql和常规的mysql到mysql之间的同步还是有些不同之处，这个在后面会说明。mysql5.6及之前都只支持一主多从，mysql5.7开始支持了多主一从，即多源复制。        
![mysql-多源复制](http://images2015.cnblogs.com/blog/645933/201601/645933-20160122152721406-1865138030.png)

说明一下，我在这个方案中，使用了docker来在一台机器上运行多个mysql，如果不使用docker，那么每个mysql需要一台机器。当然一台机器上运行多个mysql实例也是可以的，但是配置比较复杂，而且应该是tar包安装的。如果用yum或apt-get等包管理器是很难做到的，就不标准了，一台机器运行多个mysql弊大于利吧。

> 注意：在进行操作前，建议先清理下本机的环境，不然会出现一些问题，比如：
> * 确保端口3306没有被占用
> * 将备份文件恢复到一个空的文件夹里

# 从RDS备份文件恢复到自建数据库
首先我们集于某个时间点的rds备份文件，恢复到自建数据库。然后再基于这个备份的快照进行主从备份。主从备份的时候是基于gtid的，而master上的gtid过段时间会被purge掉的，所以如果master已经运行一段时间了，我们必须先从RDS备份文件进行恢复，然后再同步。       
参考了如下的阿里云官方文档，实际做的过程中还是遇到了一些坑，我还是完整的记录下整个过程吧。        
[RDS for MySQL 备份文件恢复到自建数据库](https://help.aliyun.com/knowledge_detail/41817.html)        
### 从阿里云控制台下载备份文件
[https://help.aliyun.com/knowledge_detail/5990796.html?spm=5176.7741817.2.14.DFVwwT](https://help.aliyun.com/knowledge_detail/5990796.html?spm=5176.7741817.2.14.DFVwwT)    
将文件下载到`/root/hins1891915_data_20170213022343.tar.gz`
### 解压
**1.下载专门的解压工具`rds_backup_extract`**         
`wget 'http://oss.aliyuncs.com/aliyunecs/rds_backup_extract.sh?spm=5176.7741817.2.15.DFVwwT&file=rds_backup_extract.sh' -O /root/rds_backup_extract.sh`    
**2.解压备份文件**    
`bash rds_backup_extract.sh -f /root/hins1891915_data_20170213022343.tar.gz -C /opt/mysql/mysql-1`       
### 恢复数据
**1.安装专门的恢复工具`Percona-XtraBackup`**      
先安装专门的恢复工具`Percona-XtraBackup`。因为RDS是用这个工具进行备份的。那么我们也需要用这个工具进行恢复。文档上提示RDS备份使用的版本是2.2.9，建议下载2.2.9或更新的版本。我的宿主机器是centos7，安装的是2.2.12版本，实测可用。参考官方文档[https://www.percona.com/doc/percona-xtrabackup/2.2/installation/yum_repo.html](https://www.percona.com/doc/percona-xtrabackup/2.2/installation/yum_repo.html)。       
使用以下命令来安装`percona-xtrabackup`:    
```
wget https://www.percona.com/downloads/XtraBackup/Percona-XtraBackup-2.2.12/binary/redhat/7/x86_64/percona-xtrabackup-2.2.12-1.el7.x86_64.rpm

yum localinstall percona-xtrabackup-2.2.12-1.el7.x86_64.rpm
```
> 官方文档上说使用下载rpm文件，使用`yum localinstall`来安装，需要自己去解决rpm包之间的依赖关系，这个rpm包的确会依赖一些其他的包，就是说必须要先安装它的依赖包，才能正确安装xtrabackup。但是我用`yum localinstall`的时候，yum检测出了它的依赖包，并一起帮我安装好了。这里和官方文档好像有些出入。

**2.恢复数据**    
安装好extrabackup后，就可以使用命令`innobackupex`来恢复数据了。    
```
innobackupex --defaults-file=/opt/mysql/mysql-1/backup-my.cnf --apply-log /opt/mysql/mysql-1
```

### 简单测试下备份恢复是否正确          
该步骤只是做一个简单的检测，不属于主流程。可以跳过。                    
修改文件所属为mysql用户。        
`chown -R mysql:mysql /opt/mysql/mysql-1`
修改backup-my.cnf文件。将以下几行注释掉：
```
innodb_checksum_algorithm=innodb                   --从RDS的备份的backup.my中拷贝过来
#innodb_log_checksum_algorithm=innodb              --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
innodb_data_file_path=ibdata1:200M:autoextend      --从RDS的备份的backup.my中拷贝过来
innodb_log_files_in_group=2                        --从RDS的备份的backup.my中拷贝过来
innodb_log_file_size=1572864000                    --从RDS的备份的backup.my中拷贝过来
#innodb_fast_checksum=false                        --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_page_size=16384                            --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_log_block_size=512                         --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
innodb_undo_directory=.                            --从RDS的备份的backup.my中拷贝过来
innodb_undo_tablespaces=0                          --从RDS的备份的backup.my中拷贝过来

#rds_encrypt_data=false                            --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_encrypt_algorithm=aes_128_ecb              --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
```
启动mysqld进程，并验证启动成功                
`mysqld_safe --defaults-file=/opt/mysql/mysql-1/backup-my.cnf --user=mysql --datadir=/opt/mysql/mysql-1 &`        
`mysql -uroot`查看是否可以登录，并且数据也恢复了。          

# 启动自建的mysql，进行RDS主从同步
### 启动docker容器
```
docker run --name mysql-1 -it -d -v /opt/mysql/mysql-1:/var/lib/mysql mysql:5.7.16 /bin/bash
```
一些参数的说明：        
* --name:给容器取名，方便容器间的通讯
* -it:打开一个伪终端，这样就可以进入容器的命令行
* -d:daemon,后台运行容器
* -v:将本地的文件mouont到容器里面。/var/lib/mysql是mysql默认的数据存储目录。镜像是只读的，容器是在镜像的只读层上附加了很薄的可写层。但是容器被删除的时候，容器中的数据是会丢失的，需要将容器内需要持久化的目录挂载到宿主的某个目录，这样容器内的应用在写入数据的时候，实际上就写到了宿主机器。    
### 修改my.cnf文件
需要在my.cnf中添加一些关键的参数。
attach到mysql-1容器中:        
`docker attach mysql-1` 
官方的mysql容器是没有vim的，需要我们自己安装:        
```bash
# 更新apt-get的软件源
apt-get update
# 安装vim
apt-get install vim
```
编辑/etc/mysql/my.cnf文件。docker的mysql镜像的默认配置文件是`/etc/mysql/my.cnf`,而不是`/etc/my.cnf`。
```
[mysqld]

#datadir=/opt/mysql/mysql-1

server-id=5                                        --server-id需保证唯一
master-info-repository=file
relay-log-info_repository=file
binlog-format=ROW
gtid-mode=on                                       --开启gtid模式
log_slave_updates=1   
enforce-gtid-consistency=true
log-bin=mysql-bin                                  --binlog文件存储的路径

replicate-wild-ignore-table=mysql.%
character-set-server=utf8

innodb_checksum_algorithm=innodb                   --从RDS的备份的backup.my中拷贝过来
#innodb_log_checksum_algorithm=innodb              --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
innodb_data_file_path=ibdata1:200M:autoextend      --从RDS的备份的backup.my中拷贝过来
innodb_log_files_in_group=2                        --从RDS的备份的backup.my中拷贝过来
innodb_log_file_size=1572864000                    --从RDS的备份的backup.my中拷贝过来
#innodb_fast_checksum=false                        --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_page_size=16384                            --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_log_block_size=512                         --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
innodb_undo_directory=.                            --从RDS的备份的backup.my中拷贝过来
innodb_undo_tablespaces=0                          --从RDS的备份的backup.my中拷贝过来

#rds_encrypt_data=false                            --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉
#innodb_encrypt_algorithm=aes_128_ecb              --从RDS的备份的backup.my中拷贝过来,但是RDS特有的配置，需注释掉

[mysql]
default-character-set=utf8

[client]
default-character-set=utf8
```
修改/var/lib/mysql目录的owner    
```
chown -R mysql:mysql /var/lib/mysql
```
### 启动数据库
`service mysql start`
### 设置主从同步    
**1.reset slave**    
```sql
truncate table  mysql.slave_relay_log_info;
truncate table  mysql.slave_master_info;
truncate table  mysql.slave_worker_info;

reset slave;
```
执行mysql_upgrade命令来升级。mysql是5.7的，使用的备份文件是5.6的，需要执行这个命令来升级mysql里面的一些表结构。        
```
mysql_upgrade -uroot --force
```
重启mysql    
**2.SET @@GLOBAL.GTID_PURGED**    
通过命令`head -n 50 /root/mysql-1-dump.sql`查看RDS备份文件中的xtrabackup_slave_info文件，文件内容应该如下：
```
SET GLOBAL gtid_purged='58f13aed-94d7-11e6-9d6d-1051721c39f4:1-2962994';
CHANGE MASTER TO MASTER_AUTO_POSITION=1
```
这里有一个gtid_purged.表示在备份的那个时间点，当前mysql实例当前的GTID。
```
reset master
SET @@GLOBAL.GTID_PURGED='58f13aed-94d7-11e6-9d6d-1051721c39f4:1-2962994';
```
设置了这个GTID之后，slave在同步的时候就不会从GTID=1开始同步了，而是从GTID=2962994开始同步。因为如果master运行了一段时间，master上不会保留全部的GTID，只会保留最近的一些GTID，这样如果从1开始同步，就会出错，因为master上根本就没有GTID=1，也就是1236ERROR。所以这个GTID_PURGED一定要设置正确。如果是用xtraback工具进行备份的，可以在`xtrabackup_slave_info`文件中找到这个GTID，如果是用mysqldump工具备份的，可以在生成的`.sql`的dump文件中找到。这个在后面会有说明。    
**3.change master to**    
设置本地mysql与RDS的复制关系。注意这里的用户。这个账号可以是任意的有读写权限的账号。并不需要对账号进行`grant slave replication`操作。
```sql
CHANGE MASTER TO MASTER_HOST='something.rds.aliyuncs.com',MASTER_USER='someuser', MASTER_PASSWORD='somepassword',master_auto_position=1;  
```    
**4.查看同步状态**
```sql
start slave;
show slave status \G
```    

# 多源复制
## 在自建数据库上创建复制账号
```
create user 'replicator'@'%' identified by 'password';
grant replication slave on *.* to 'replicator'@'%' identified by 'password';
flush privileges;
```
### 在自建mysql上导出数据库
`mysqldump -uroot --master-data=2 --single-transaction --add-drop-database --all-databases > /root/mysql-1-dump.sql`
### 启动一个mysql来merge所有的数据
`docker run --name mysql-all -it -d -v /opt/mysql/rds-all:/var/lib/mysql --link=mysql-1:mysql-1 --link=mysql-2:mysql-2 -p 3307:3306 mysql:5.7.16 /bin/bash`
### 将备份拷贝到merge msql
```
docker cp mysql-1:/root/mysql-1-dump.sql /root
docker cp /root/mysql-1-dump.sql mysql-all:/root

docker cp mysql-2:/root/mysql-2-dump.sql /root
docker cp /root/mysql-2-dump.sql rds-all:/root
```
### 修改merge mysql的my.cnf文件
`[root@rds-all] vim /etc/mysql/my.cnf`
```
[mysqld]
server-id=18
binlog-format=ROW
gtid-mode=on
log_slave_updates=1
enforce-gtid-consistency=true
log-bin=mysql-bin
master_info_repository=TABLE
relay_log_info_repository=TABLE

replicate-wild-ignore-table=mysql.%
character-set-server=utf8

[mysql]
default-character-set=utf8

[client]
default-character-set=utf8
```
### 修改导出的sql文件
在导出的sql文件的最开头部分，有以下的设置GTID的语句：
mysql-1:    
```
--
-- GTID state at the beginning of the backup
--

-- SET @@GLOBAL.GTID_PURGED='6c974343-94d6-11e6-9d67-1051721c39f4:1-32772706,
-- 72c54af0-f1bf-11e6-88fa-0242ac110002:1-4';
```
mysql-2:    
```
--
-- GTID state at the beginning of the backup
--

-- SET @@GLOBAL.GTID_PURGED='416e3511-f1d0-11e6-bb84-0242ac110003:1-2,
-- 58f13aed-94d7-11e6-9d6d-1051721c39f4:1-3141347';
```
需要将这两个GTID记录下来，后面会用到。然后再将这条语句***注释掉***，不要去执行，我们后面会自己手动去执行该语句，自动执行会有问题。这里的GTID可能有一个，可能有多个，视具体情况而定。
### 数据导入
由于mysql-1和mysql-2都是5.6的，而all是5.7的，直接将5.6的dump文件导入到5.7是有错误的，会提示一些字段没有或者schema不匹配等，需先执行`mysql_upgrade`命令来升级mysql。在mysql-all的机器上执行。    
```
service mysql start
mysql_upgrade -uroot
```    
然后需要重启mysql-all机器。重启后就可以导入数据了。
```
mysql -uroot < /root/mysql-1-dump.sql
mysql -uroot < /root/mysql-1-dump.sql
```
### set global.gtid_urged
将上一步中在dump的sql文件中出现的所有的GTID都添加到global.gtid_purged中。
```
reset master;
SET @@GLOBAL.GTID_PURGED='6c974343-94d6-11e6-9d67-1051721c39f4:1-32863117,72c54af0-f1bf-11e6-88fa-0242ac110002:1-4,416e3511-f1d0-11e6-bb84-0242ac110003:1-2,58f13aed-94d7-11e6-9d6d-1051721c39f4:1-3175513';
```
### change master
```
CHANGE MASTER TO MASTER_HOST='mysql-1',MASTER_USER='replicator', MASTER_PASSWORD='replicator',master_auto_position=1 FOR CHANNEL 'mysql-1';
CHANGE MASTER TO MASTER_HOST='mysql-2',MASTER_USER='replicator', MASTER_PASSWORD='replicator',master_auto_position=1 FOR CHANNEL 'mysql-2';

start slave for channel 'mysql-1';
start slave for channel 'mysql-2';

show slave status for channel 'mysql-1';
show slave status for channel 'mysql-2';
```

# 重启mysql后需要重新设置global.gtid_urged
经测试，在重启了mysql后，mysql的主从同步会失败，因为之前的gtid_purged的设置丢失了。需要重新设置global.gtid_purged。重启后设置gtid_purged会有些麻烦。我的方法是在重启之前记录下当前mysql的gtid状态。
```
SHOW GLOBAL VARIABLES LIKE '%gtid%';
```
输出的结果如下：
```
| gtid_executed                    | 08bf633d-f297-11e6-bcff-0242ac110004:1-3,
416e3511-f1d0-11e6-bb84-0242ac110003:1-2,
58f13aed-94d7-11e6-9d6d-1051721c39f4:1-3241262,
6c974343-94d6-11e6-9d67-1051721c39f4:1-33028994,
72c54af0-f1bf-11e6-88fa-0242ac110002:1-4 |
| gtid_executed_compression_period | 1000                                                                                                                                                                                                                          |
| gtid_mode                        | ON                                                                                                                                                                                                                                                                                                                                                  |
| gtid_purged                      | 416e3511-f1d0-11e6-bb84-0242ac110003:1-2,
58f13aed-94d7-11e6-9d6d-1051721c39f4:1-3175513,
6c974343-94d6-11e6-9d67-1051721c39f4:1-32863117,
72c54af0-f1bf-11e6-88fa-0242ac110002:1-4                                           |
```
里面有gtid_executed和gtid_purged.我们要用到的是gtid_executed参数。重启后将global.gtid_purged设置为重启前的global.gtid_executed就可以啦。      
先reset master      
```sql
reset master
stop slave
```
set global.gtid_executed        
```
SET @@GLOBAL.GTID_PURGED='08bf633d-f297-11e6-bcff-0242ac110004:1-3,416e3511-f1d0-11e6-bb84-0242ac110003:1-2,58f13aed-94d7-11e6-9d6d-1051721c39f4:1-3241262,6c974343-94d6-11e6-9d67-1051721c39f4:1-33028994,72c54af0-f1bf-11e6-88fa-0242ac110002:1-4'
```
start slave    
```
start slave for channel 'mysql-1';
start slave for channel 'mysql-2';
```
如果重启前没有记录gtid_executed.也可以在重启后，执行
```
show slave status for channel 'mysql-1' \G
show slave status for channel 'mysql-2' \G
将输出结果里面的所有executed_gtid合并起来，再set global.gtid_purged就可以啦
```
# 其他
我用docker是因为机器不够，而且在实验这套方案的时候，我可能需要尝试不同版本的mysql，需要频繁的删除，安装mysql，在一台物理机上多次删除，安装mysql试过的人应该知道是什么感觉。。。就像在windows上多次安装删除sqlserver一样，谁用谁知道。。。。另外，安装好mysql后，我们肯定需要对mysql做一些配置的，比如在my.cnf中配置字符集，添加mysql主从同步的配置。通过docker，我们可以将my.cnf文件从宿主机器mount到mysql容器中，宿主中的my.cnf可以托管在git，svn等仓库中，防止丢失。最后，如果不使用docker，有人将mysql机器搞挂了，比如将文件误删掉，或配置改错了，怎么办？排查和恢复很困难。如果用docker，因为docker镜像的只读性和快速启动性，很容器恢复。而且，如果搭配了docker compose可以做到一条命令就可以搭建起整个的环境，启动3台mysql容器，配置并运行。而镜像的创建是可以基于DockerFile创建的。DockerFile和docker compose都是文本文件，mysql的配置文件也是文本文件，所以，搭建这一整套复杂的环境所需要的都是一些文本文件。这些文本文件可以托管在git里，防止丢失和进行版本控制。最极端的情况，运行mysql的机器彻底毁掉了，只需要再给我一台新机器，我在这个机器上安装好docker engine，然后将docker compose文件拷贝到这台机器上，运行一个`dockercompose up`命令就可以了。     
