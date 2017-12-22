---
layout: post
title: Git Guidance
date: 2015-06-25 19:30:00
tags:
- Git
categories: Git
---


|      command       |                                  |                                                              |
| ------------------ | -------------------------------- | ------------------------------------------------------------ |
| `reset`            | `reset <commit_id>`              | 将一个分支的HEAD重置到某个版本                                   |

# yum安装git
```bash
yum install git
```

# 创建新仓库    
* 新建文件夹`mk dir my_git_repo`    
* 进入新建的文件夹`cd my_git_repo`    
* 执行命令`git init`    

# 克隆远程仓库        
|                        command                   |                     description                  |
| ------------------------------------------------ | ------------------------------------------------ |
| git clone <repo_url> path/to/folder              | 将远程仓库克隆到本地的<当前目录>/path/to/folder       |
| git clone <url> -b branch_name path/to/folder    | 将远程仓库克隆到本地的<当前目录>/path/to/folder.<br>并checkout出`branch_name`分支.|

> git仓库地址有多种协议，如http或ssh。http／https需要你在克隆或提交的时候提供用户名和密码。ssh需要你在本地生成一个SSH key，并把这个public key添加到git server中。比如，github中。使用`ssh-keygen`来生成SSH key。    

# Git工作流    
本地仓库由git维护的3棵树组成。
1. `工作目录`，持有实际文件。    
2. `暂存区`,像一个缓存区域，临时保存你的改动。    
3. `HEAD`,指向你最后一次提交的结果。    


working diretory    ------------->    stage    ------------->  HEAD
                         add                      commit

# Git添加和提交
把更改添加到暂存区：`git add <file_name>` 或 `git add *`    
实际提交改动：`git commit -m "这是注释"`    
现在你的改动已经提交到流HEAD，但还没到你的远程仓库。执行命令`git push origin <branch_name>`把改动提交到远端分支。
如果你还没有为你到本地仓库添加远程仓库，可以使用命令`git remote add <origin_name> <remote_repo_url>`         

# Git分支操作
在创建仓库的时候，master是默认分支。可以在其他分支上开发某一个feature，然后将该分支合并到master上。
## 创建分支
使用命令`git checkout -b <branch_name>`来创建一个分支，并切换过去。
## 删除分支
使用命令`git branch -d <branch_name>`来删除一个分支。
除非你将分支push到远端仓库，不然该分支就是不为他人所知的，只是存在与你的本地仓库中。
## 切换分支
使用`git checkout <branch_name>`来切换分支。
## 更新分支
使用`git pull`来更新本地仓库到最新改动。
## 合并
`git merge <branch_name>`将分支上的改动合并到当前分支。
## 重置本地改动
`git checkout -- <file_name>`来使用`HEAD`中的版本替换掉工作目录中的文件。已添加到暂存区的改动以及新文件都不会收到影响。
如果你想丢弃掉你在本地的所有改动与提交，可以到服务器上获取最新的版本，并将你的本地分支指向它。
`git fetch origin` `git reset --hard origin/master`








# Git命令详解
## reset
将一个分支的HEAD重置到某个版本。有如下几个选项：

1. `--soft <commit_id>`        将一个分支的HEAD重置到某个版本。但是缓存区和工作目录不改变。           
2. `reset --mixed <commit_id>` 将一个分支的HEAD重置到某个版本。默认选项。影响缓存区。工作目录不影响。
3. `reset --hard <commit_id>`  将一个分支的HEAD重置到某个版本。影响缓存区和工作目录。               


# Resources
[http://rogerdudler.github.io/git-guide/index.zh.html](http://rogerdudler.github.io/git-guide/index.zh.html)
[http://marklodato.github.io/visual-git-guide/index-zh-cn.html](http://marklodato.github.io/visual-git-guide/index-zh-cn.html)        
[http://www.cnblogs.com/itech/p/5188933.html](http://www.cnblogs.com/itech/p/5188933.html)      
