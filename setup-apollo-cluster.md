---
title: Setup Apollo Cluster
date: 2017-06-10 10:22:23
categories:
- Music
tags:
- Music
---

# JAR部署情况：

|    env    |  private ip:port      |  public ip:port     |          jar          |
| --------- | --------------------- | ------------------- | --------------------- |
| dev       | 192.168.4.100:8080    |                     | config-service.jar    |
| dev       | 192.168.4.100:8090    |                     | admin-service.jar     |
| pro       | 172.16.128.100:8080   | 11.22.33.44:1111    | config-service.jar    |
| pro       | 172.16.128.100:8090   | 11.22.33.44:2222    | admin-service.jar     |
| dev       | 192.168.4.200:8070    |                     | portal.jar            |

我们选择将Portal部署在dev环境    

# 数据库部署情况：

|     env    |                          jdbc                                            |  user  |         |
| ---------- | ------------------------------------------------------------------------ | ------ | ------- |
| dev        | jdbc:mysql://192.168.4.300:3306/ApolloConfigDB?characterEncoding=utf8    | dev    |  dev    |
| pro        | jdbc:mysql://172.16.128.200:3306/ApolloConfigDB?characterEncoding=utf8   | pro    |  pro    |
| dev        | jdbc:mysql://192.168.4.300:3306/ApolloPortalDB?characterEncoding=utf8    | dev    |  dev    |


# 执行数据库
1. 在dev环境执行${Apollo}/scripts/sql/apolloconfigdb.sql
1. 在pro环境执行${Apollo}/scripts/sql/apolloconfigdb.sql
3. 在dev环境执行${Apollo}/scripts/sql/apolloportaldb.sql.sql
4. 配置ApolloPortalDB
```sql
use ApolloPortalDB
update ServerConfig set Value='dev,pro' where `Key`='apollo.portal.envs'
```

# 构建dev环境的Config-Service和Admin-Service的JAR包
1. 下载Apollo源码
2. 编辑${Apollo}/scripts/build.sh
```bash
# apollo config db info
apollo_config_db_url=jdbc:mysql://192.168.4.300:3306/ApolloConfigDB?characterEncoding=utf8
apollo_config_db_username=dev
apollo_config_db_password=dev

# apollo portal db info
apollo_portal_db_url=jdbc:mysql://192.168.4.300:3306/ApolloPortalDB?characterEncoding=utf8
apollo_portal_db_username=dev
apollo_portal_db_password=dev

# meta server url, different environments should have different meta server addresses
dev_meta=http://192.168.4.100:8080
fat_meta=http://someIp:8080
uat_meta=http://anotherIp:8080
pro_meta=http://11.22.33.44:1111
```
3. 执行${Apollo}/scripts/build.sh文件    
构建好的文件在如下位置：    
* apollo-configservice/target/apollo-configservice-0.9.0-SNAPSHOT-github.zip
* apollo-adminservice/target/apollo-adminservice-0.9.0-SNAPSHOT-github.zip

# 构建pro环境的Config-Service和Admin-Service的JAR包
1. 下载Apollo源码
2. 编辑${Apollo}/scripts/build.sh
```bash
# apollo config db info
apollo_config_db_url=jdbc:mysql://172.16.128.200:3306/ApolloConfigDB?characterEncoding=utf8
apollo_config_db_username=dev
apollo_config_db_password=dev

# apollo portal db info
apollo_portal_db_url=jdbc:mysql://192.168.4.300:3306/ApolloPortalDB?characterEncoding=utf8
apollo_portal_db_username=dev
apollo_portal_db_password=dev

# meta server url, different environments should have different meta server addresses
dev_meta=http://192.168.4.100:8080
fat_meta=http://someIp:8080
uat_meta=http://anotherIp:8080
pro_meta=http://11.22.33.44:1111
```
3. 执行${Apollo}/scripts/build.sh文件    
构建好的文件在如下位置：    
* apollo-configservice/target/apollo-configservice-0.9.0-SNAPSHOT-github.zip
* apollo-adminservice/target/apollo-adminservice-0.9.0-SNAPSHOT-github.zip

# 构建Portal的JAR包
在上面执行build.sh文件的时候，也会同时生成Portal的JAR文件，在如下路径：
* apollo-portal/target/apollo-portal-0.9.0-SNAPSHOT-github.zip

# 构建Client的JAR包    
Client的JAR包是需要上传到Maven私服的，比如Nexus。
在执行上面的build.sh文件时，会执行`mvn clean install`命令，将client的jar包安装到本地的maven仓库，可以自己编辑build.sh文件，将命令改成`mvn clean deploy`，同时在maven的setting.xml文件中配置相应的nexus私服信息。

1. 编辑${Apollo}/build.sh文件，将`mvn clean install`改成`mvn clean deploy`    
```bash
echo "==== starting to build client ===="

mvn clean deploy -DskipTests -pl apollo-client -am $META_SERVERS_OPTS

echo "==== building client finished ===="
```
2. 编辑`.m2/settings.xml`文件,添加Nexus私服的相关信息
```xml
<servers>
   <server>
       <id>releases</id>
       <username>admin</username>
       <password>admin123</password>
   </server>
   <server>
       <id>snapshots</id>
       <username>admin</username>
       <password>admin123</password>
   </server>
</servers>
```

```xml
<profiles>
      <profile>
          <id>pnt</id>
          <activation>
              <jdk>1.8</jdk>
          </activation>
          <properties>
              <releases.repo>http://${nexus_host}:${nexus_port}/nexus/content/repositories/releases</releases.repo>
              <snapshots.repo>http://${nexus_host}:${nexus_port}/nexus/content/repositories/snapshots</snapshots.repo>
          </properties>
          <repositories>
              <repository>
                  <id>nexus</id>
                  <name>local private nexus</name>
                  <url>http://${nexus_host}:${nexus_port}/nexus/content/groups/public/</url>
                  <releases>
                      <enabled>true</enabled>
                  </releases>
                  <snapshots>
                      <enabled>true</enabled>
                  </snapshots>
              </repository>
          </repositories>
          <pluginRepositories>
              <pluginRepository>
                  <id>nexus</id>
                  <name>local private nexus</name>
                  <url>http://${nexus_host}:${nexus_port}/nexus/content/groups/public/</url>
                  <releases>
                      <enabled>true</enabled>
                  </releases>
                  <snapshots>
                      <enabled>true</enabled>
                  </snapshots>
              </pluginRepository>
          </pluginRepositories>
      </profile>
  </profiles>
```

3. 执行build.sh文件          
build.sh脚本会构建client的 JAR包并上传到maven私服

# 运行dev环境的Config Service
1. 将上面生成的zip文件拷贝到服务器，解压
2. 修改/scripts/startup.sh文件
修改日志的目录，不然有可能出现没有权限，或目录不存在的错误
```bash
## Adjust log dir if necessary
LOG_DIR=/opt/logs
```
3. 运行/scripts/startup.sh文件

# 运行dev环境的Admin Service
1. 将上面生成的zip文件拷贝到服务器，解压
2. 修改/scripts/startup.sh文件
修改日志的目录，不然有可能出现没有权限，或目录不存在的错误
```bash
## Adjust log dir if necessary
LOG_DIR=/opt/logs
```
3. 运行/scripts/startup.sh文件

# 运行pro环境的Config Service
1. 将上面生成的zip文件拷贝到服务器，解压
2. 修改/scripts/startup.sh文件
修改日志的目录，不然有可能出现没有权限，或目录不存在的错误
```bash
## Adjust log dir if necessary
LOG_DIR=/opt/logs
```
Config Service在向Eureka注册的时候，需要将公网的ip和端口注册进去，而不能是内网的地址，需要设置`eureka.instance.homePageUrl`和`eureka.instance.preferIpAddress`参数    
```bash
export JAVA_OPTS="$JAVA_OPTS -Dserver.port=$SERVER_PORT -Dlogging.file=$LOG_DIR/$SERVICE_NAME.log -Xloggc:$LOG_DIR/heap_trace.txt -XX:HeapDumpPath=$LOG_DIR/HeapDumpOnOutOfMemoryError/ -Deureka.instance.homePageUrl=http://11.22.33.44:1111 -Deureka.instance.preferIpAddress=true"
```
3. 运行/scripts/startup.sh文件

# 运行pro环境的Admin Service
1. 将上面生成的zip文件拷贝到服务器，解压
2. 修改/scripts/startup.sh文件
修改日志的目录，不然有可能出现没有权限，或目录不存在的错误
```bash
## Adjust log dir if necessary
LOG_DIR=/opt/logs
```
Config Service在向Eureka注册的时候，需要将公网的ip和端口注册进去，而不能是内网的地址，需要设置`eureka.instance.homePageUrl`和`eureka.instance.preferIpAddress`参数    
```bash
export JAVA_OPTS="$JAVA_OPTS -Dserver.port=$SERVER_PORT -Dlogging.file=$LOG_DIR/$SERVICE_NAME.log -Xloggc:$LOG_DIR/heap_trace.txt -XX:HeapDumpPath=$LOG_DIR/HeapDumpOnOutOfMemoryError/ -Deureka.instance.homePageUrl=http://11.22.33.44:2222 -Deureka.instance.preferIpAddress=true"
```
3. 运行/scripts/startup.sh文件


# 运行Portal
1. 将上面生成的zip文件拷贝到服务器，解压
2. 修改/scripts/startup.sh文件
修改日志的目录，不然有可能出现没有权限，或目录不存在的错误
```bash
## Adjust log dir if necessary
LOG_DIR=/opt/logs
```
3. 运行/scripts/startup.sh文件

# 应用程序使用Client
1. 添加Maven依赖
```groovy
dependencies {
    compile("com.ctrip.framework.apollo:apollo-client:0.9.0-SNAPSHOT")
}
```
2. 运行应用程序
在dev环境运行
```bash
java -jar -Denv=dev apollo-sample.jar
```

在pro环境运行
```bash
java -jar -Denv=pro apollo-sample.jar
```
