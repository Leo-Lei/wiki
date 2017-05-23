---
layout: post
title: Python
date: 2017-05-23 16:30:17
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---




# 删除历史文件
```python
#! /usr/bin/python
# -*- coding: utf-8 -*-

import datetime
import os


def run():
    root_path = '/opt/logs'
    history_logs = {
        'app-bff': {'filename': 'app-bff.log.{%Y%m%d}.log', 'keep_days': 5},
        'catalina': {'filename': 'catalina.{%Y-%m-%d}.log', 'keep_days': 3},
        'heartbeat': {'filename': 'heartbeat.log.{%Y%m%d}.log', 'keep_days': 3},
        'host-manager': {'filename': 'host-manager.{%Y-%m-%d}.log', 'keep_days': 3},
        'localhost': {'filename': 'localhost.{%Y-%m-%d}.log', 'keep_days': 3},
        'localhost_access': {'filename': 'localhost_access_log.{%Y-%m-%d}.txt', 'keep_days': 3},
        'manager': {'filename': 'manager.{%Y-%m-%d}.log', 'keep_days': 3}
    }

    now = datetime.datetime.now()
    for (k, v) in history_logs.items():
        keep_days = v['keep_days']
        filename = v['filename']
        for i in range(keep_days, 100):
            early_date = now + datetime.timedelta(days=-i)
            date_format_str = extract_date_from_filename(filename)
            early_date_str = early_date.strftime(date_format_str)  # the format like 20170317
            real_filename = os.path.join(root_path,filename.replace('{' + date_format_str + '}', early_date_str))
            cmd = "rm -rf {0}".format(real_filename)

            if os.path.exists(real_filename):
                print cmd
                os.system(cmd)


def extract_date_from_filename(filename):
    a = filename.index('{')
    b = filename.index('}')
    return filename[a + 1:b]


if __name__ == '__main__':
    run()

```
