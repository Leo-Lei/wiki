---
layout: post
title: Git rebase
date: 2016-06-17 14:50:00
tags:
- git
categories: Windows
---





```bash

touch hello.txt
git add hello.txt
git commit -m "add hello.txt"

echo 11111 >> hello.txt
git add hello.txt
git commit -m "add 11111 to hello.txt"

echo 22222 >> hello.txt
git add hello.txt
git commit -m "add 22222 to hello.txt"

echo 33333 >> hello.txt
git add hello.txt
git commit -m "add 33333 to hello.txt"
```

查看git提交记录
```bash
git log

commit c75fa6f2cff8e60da488e3cb1c25daafb5d22ab2
add 33333 to hello.txt

commit fe6e6904bac6c1b3a3c8fe58c71ac4efd6ac26cd
add 22222 to hello.txt

commit 0a456ea26f243fab5f808396c30d46c8df4732a9
add 11111 to hello.txt

commit b47ddd8c039e1180c3c81ccf4c3263dcce134233
add hello.txt
```


```bash
git rebase -i b47ddd8c039e1180c3c81ccf4c3263dcce134233
```

```bash
pick 0a456e add 11111 to hello.txt
pick fe6e69 add 22222 to hello.txt
pick c75fa6 add 33333 to hello.txt

 # Rebase fe6e690..c75fa6f onto fe6e690 (1 command)
  4 #
  5 # Commands:
  6 # p, pick = use commit
  7 # r, reword = use commit, but edit the commit message
  8 # e, edit = use commit, but stop for amending
  9 # s, squash = use commit, but meld into previous commit
 10 # f, fixup = like "squash", but discard this commit's log message
 11 # x, exec = run command (the rest of the line) using shell
 12 # d, drop = remove commit
 13 #
 14 # These lines can be re-ordered; they are executed from top to bottom.
 15 #
```


```bash
git rebase -i b47ddd8c039e1180c3c81ccf4c3263dcce134233
```

```bash
pick 0a456e add 11111 to hello.txt
squash fe6e69 add 22222 to hello.txt
squash c75fa6 add 33333 to hello.txt
```
`wq`保存并退出。

```bash
# This is a combination of 3 commits.
  2 # This is the 1st commit message:
  3 
  4 add 11111 to hello.txt
  5 
  6 # This is the commit message #2:
  7 
  8 add 22222 to hello.txt
  9 
 10 # This is the commit message #3:
 11 
 12 add 33333 to hello.txt
 13 
 14 # Please enter the commit message for your changes. Lines starting
 15 # with '#' will be ignored, and an empty message aborts the commit.
```

在弹出来的框中进行编辑commit注释。


```bash
# This is a combination of 3 commits.

Add 11111, 22222, 33333 to hello.txt
  
# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.
```


```bash
git log

commit dabca8193b6950a72bf7eece6ced31e8949246ee
Add 11111, 22222, 33333 to hello.txt
commit b47ddd8c039e1180c3c81ccf4c3263dcce134233
add hello.txt
```


