---
layout: post
title: How to Use PIP
date: 2016-01-15 14:00:00
tags:
- Python
categories: Python
description: The tutoria will show you how to set up the Python environment.
---

# Install package online
```bash
pip install Flask
```
How to find the location of Python 3rd-party packages?
```bash
$ python
>>> import site;
>>> site.getsitepackages()
['C:\\Python27', 'C:\\Python27\\lib\\site-packages']
>>> exit()
```

Install package behind a proxy
```bash    
$ pip install Flask --proxy="http://127.0.0.1:8080"
```

# PIP install package offline
In some scenario, you may are working with no internet connection, or it is unavailable to connect to Python repository and download the packages, due to some network problems. In this case, you can use the pip under offline mode.

The following steps illustrate how to use the PIP offline:    
1. Create a directory as the PIP repository.    
2. Use PIP to download the packages to the PIP repository you created, while you can access the internet.    
3. Copy the PIP repository to the environment which has no internet access.    
4. Use the PIP to install package by specifying to use the local PIP repository, rather than the remote repository.    

**Sample:**    
Create a local PIP repository:
```bash    
$ mkdir C:\foo\my-pip-repository
```

Download the Flask(include all its dependencies) package to the local PIP repository:
```bash   
$ pip install --download C:\foo\my-pip-repository Flask --proxy="http://127.0.0.1:8080"
```

Install Flask from local PIP repository:
```bash    
$ pip install --no-index --find-links=C:\foo\my-pip-repository Flask
```

