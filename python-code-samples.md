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
    now = datetime.datetime.now()
    for i in range(1, 10):
        early_date = now + datetime.timedelta(days=-i)
        early_date_str = early_date.strftime('%Y%m%d')   # the format like 20170317
        cmd = "rm -rf /opt/logs/myapp.log.{0}.log".format(early_date_str)
        print cmd
        os.system(cmd)

if __name__ == '__main__':
    run()

```
