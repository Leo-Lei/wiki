---
layout: post
title: Setup Python Environment
date: 2016-06-22 09:58:17
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---

# 1. Install Python
# 2. Code Style of Python
# 2.1 Naming Conventions

| Type                    |            Public           |               Internal                                         |
| ----------------------- | --------------------------  | -------------------------------------------------------------- |
| Package                 | `lower_with_under`          |                                                                |
| Module                  | `lower_with_under`          | `_lower_with_underscore`                                       |
| Class                   | `CapWords`                  | `_CapWords`                                                    |
| Exception               | `Capwords`                  | `_Capworks`                                                    |
| Function                | `firstLowerCapWords()`      | `firstLowerCapWords()`                                         |
| Global/Class Variable   | `lower_with_under`          | `_lower_with_underscore`                                       |
| Global/Class Constant   | `CAPS_WITH_UNDER`           | `_CAPS_WITH_UNDERSCORE`                                        |
| Instance Variables      | `lower_with_under`          | `_lower_with_under`(protected) or `__lower_with_under`(private)|
| Function Parameters     | `lower_with_under`          |                                                                |
| Local Variable          | `lower_with_under`          |                                                                |

# 3. Install 3rd-party Modules/Packages
# 3.1 Online Install
Use `easy_install` or `pip` to install the package online.
# 3.2 Offline Install
You can find the details from the following post.
[How to pip install python packages offline](http://www.nyayapati.com/srao/2014/06/how-to-pip-install-python-packages-offline/)
