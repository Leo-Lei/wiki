---
layout: post
title: Debian / Ubuntu
date: 2017-03-12 11:10:00
tags:
- docker
categories: Java
description: docker
---

# apt-get

|                           command                                       |          usage                               | 
| ----------------------------------------------------------------------- | -------------------------------------------- | 
| `apt-get install vim`                                                   | 安装vim                                       | 
| `apt-get update`                                                        | 更新apt-get源                                 |


# init.d服务

```bash
#!/bin/bash
### BEGIN INIT INFO
#
# Provides:     location_server
# Required-Start:    $local_fs  $remote_fs
# Required-Stop:    $local_fs  $remote_fs
# Default-Start:     2 3 4 5
# Default-Stop:     0 1 6
# Short-Description:    initscript
# Description:     This file should be used to construct scripts to be placed in /etc/init.d.
#
### END INIT INFO

## Fill in name of program here.
SERVICE_NAME=app
EXEC_START="java -jar /opt/app.jar"
PID_PATH="/var/run/"

start() {
    if [ -e "$PID_PATH/$SERVICE_NAME.pid" ]; then
        ## Program is running, exit with error.
        echo "Error! $PROG is currently running!" 1>&2
        exit 1
    else
        ## Change from /dev/null to something like /var/log/$PROG if you want to save output.
        $EXEC_START 2>&1 >/var/log/$SERVICE_NAME &
        pid=`ps -ef | grep '/opt/app.jar' |grep -v "grep"|awk '{print $2}'`
        echo "$SERVICE_NAME started"
        echo $pid > "$PID_PATH/$SERVICE_NAME.pid"
    fi
}

stop() {
    echo "begin stop"
    if [ -e "$PID_PATH/$SERVICE_NAME.pid" ]; then
        ## Program is running, so stop it
        pid=`ps -ef | grep '/opt/app.jar' |grep -v "grep"|awk '{print $2}'`
        kill -9 $pid
        rm -f  "$PID_PATH/$SERVICE_NAME.pid"
        echo "$SERVICE_NAME stopped"
    else
        ## Program is not running, exit with error.
        echo "Error! $SERVICE_NAME not started!" 1>&2
        exit 1
    fi
}

## Check to see if we are running as root first.
## Found at http://www.cyberciti.biz/tips/shell-root-user-check-script.html
if [ "$(id -u)" != "0" ]; then
    echo "This script must be run as root" 1>&2
    exit 1
fi

case "$1" in
    start)
        start
        exit 0
    ;;
    stop)
        stop
        exit 0
    ;;
    reload|restart|force-reload)
        stop
        start
        exit 0
    ;;
    **)
        echo "Usage: $0 {start|stop|reload}" 1>&2
        exit 1
    ;;
esac
```
