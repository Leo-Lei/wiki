---
layout: post
title: Base useage of Linux
date: 2015-06-26 15:30:00
tags:
- Linux
categories: Linux
---

# 2. VI
VI is a text editor in Linux.
VI has 3 modes:    
1. Command mode.    
2. Insert mode.    
3. Last line mode.    
In **Command** mode, you can't edit the file. press `i` to enter **Insert** mode. In `Insert` mode, you can edit the file. While finish editing the file, press the `Esc` to switch to **Command** mode. Then press the `:` to enter **Last Line** mode. In the **Last Line** mode, type `wq` to save and quit vi, or `q!` to exit **without** saving changes.

# 3. From Windows to Linux
## 3.1 Login to Linux via SSH
Login to Linux from Windows is using SSH protocal. So you need to isntall a SSH client. The most popular SSH client is `putty`. Putty is green and free.
Putty has only one window, so you can use below plugin to extend Putty to support multiple-Tabbed:
[MTPutty](http://www.ttyplus.com/multi-tabbed-putty).
Other SSH clients for Windows: XShell.

## 3.2 FTP From Windows to Linux
Windows contains the `ftp` command out of box. So no additional client is required.

## 3.3 SFTP From Windows to Linux
Windows doesn't support `sftp` command. In order to use SFTP functionality in Windows, some SFTP client is required. The simplest tool is `psftp`. You can download the tool from [Putty Download Page](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html). The psftp is just an standalone executable exe file.
You can double-click the psftp.exe to run the program, but in this case copy and paste with clipboard is not supported. So I suggest you to add the path of psftp.exe to the Windows System Environment `Path`, then you can type `psftp` in CMD prompt window to run the program. In this case, copy and paste is supported. 

The following is common *psftp* command:

| Command   |             Description                    |      Syntax                                          |
| --------- | ------------------------------------------ | ---------------------------------------------------- |
| pwd       | print remote directory                     | `pwd`                                                |
| lpwd      | print local directory                      | `lpwd`                                               |
| ls        | list content of remote directory           | `ls`                                                 |
| get       | get a file from remote to local directory  | `get <remote_file>` `get <remote_file> <local_path>` |
| put       | create a new directory                     | `mkdir <folder_name>`                                |
| cd        | change remote directory                    | `cd <remote_directory>`                              |
| lcd       | change local directory                     | `lcd <local_directory>`                              |

# 4. Linux File Link
## 4.1 Concept of Linux Link
There are two kinds of file link in Linux, one is **Hard Link**, the other is **Symbolic Link**.
## 4.2 Hard Link
In Linux, every physical file located on the driver has a uniq id in File System, which is known as **Inode Index**. Pay attention to it that I said physical file. You can use the `ls -li` command to view the Inode Index. Sometimes, you may find that two files in the list have the same Inode Index. In this condition, there is a Hard Link among these files. You can image the Hard Link as a point. Say that we have a file f1, then we create a Hard Link f2 to f1. Then we  have two files, f1 and f2. These 2 files has the same Inode Index.
## 4.3 Symbolic Link
When create a Symbolic link to a file, we indeed create a new file on the driver. The file has a different Inode Index, and actually, the Symbolic file is a text file, contains the link information.
## 4.4 How to Create/Delete File Link
The command `ln` is used in creating Link. The Syntax is `ln <original_file> <link_file>`. By default, it will create a Hard Link, you can use the `-s` option to specify creating a Symbolic Link. The following are some examples:
1. Create a Hard Link f2, linked to file f1
`ln f1 f2`
2. Create a Symbolic Link f3, linked to file f1
`ln -s f1 f3`
To delete the link files is the same with deleting a normal file, i.e. use the `rm` command. As the link file is basically a file.

## 4.5 A Simple Sample of Link
```bash
[oracle@Linux]$ touch f1          #创建一个测试文件f1
[oracle@Linux]$ ln f1 f2          #创建f1的一个硬连接文件f2
[oracle@Linux]$ ln -s f1 f3       #创建f1的一个符号连接文件f3
[oracle@Linux]$ ls -li            # -i参数显示文件的inode节点信息
total 0
9797648 -rw-r--r--  2 oracle oinstall 0 Apr 21 08:11 f1
9797648 -rw-r--r--  2 oracle oinstall 0 Apr 21 08:11 f2
9797649 lrwxrwxrwx  1 oracle oinstall 2 Apr 21 08:11 f3 -> f1
```
You can see, the Hard Link f2 has the same Inode Index with the original file f1. While the Symbolic file has a different Inode Index.
## 4.6 What Happened while Operating the Link Files
Say that we have some files:
1. f1: a normal text file with the content "Hello World"
2. f2: a Hard Link to f1.
3. f3: a Symbolic link to f1.

Bellow are some test scenarios:

| No  | Action                      |             f1               |      f2                        |   f3                         |
| --- | --------------------------- | ---------------------------- | ------------------------------ | ---------------------------- |
| 1   | Delete f1.(`rm -f f1`)      | run `cat f1`: No such file   | run `cat f2`: Hello World      | run `cat f3`: No such file.  |
| 2   | Delete f2.(`rm -f f2`)      | No such file                 | No such file                   | No such file                 |

When we run the command `rm -f f1`, the actual file located on Driver is not been deleted, until we delete the f2 in the following. Just when the original file and all Hard Links to it are deleted, the actual physical file on Driver will be deleted from Driver.

Here are some conclusion:
1. Delete Symbolic Link(f3) will not delete file on Driver. It has no effect on f1 and f2.
2. Just delete Hard Link f2, and remain f1 will not delete file on Driver.


# 4. FTP/SFTP
Login to a FTP/SFTP server:
`sftp <server>`

# 5. Bash Redirection

|             Command              |            Descrip                                                      |
| -------------------------------- | ----------------------------------------------------------------------- |
| `cmd > file`                     | Redirect the standard output(stdout) of cmd to a file.                  |
| `cmd 1> file`                    | The same as `cmd > file`.                                               |
| `cmd 2> file`                    | Redirect the standard error(stderr) of cmd to a file.                   |
| `cmd >> file`                    | Append output of cmd to a file.                                         |
| `cmd 2>> file`                   | Append stderr of cmd to a file.                                         |
| `cmd &> file`                    | Redirect stdout and stderr of cmd to a file.                            |
| `cmd | tee file`                 | Redirect stdout of cmd to a file and print it to screen                 |
| `cmd  2>&1 | tee -a  file`       | Append stdout and stderr of cmd to a file and print it to screen.       |
| `cmd1 | cmd2`                    | Redirect stdout of cmd1 to stdin of cmd2. The same as cmd1 >> (cmd2)    |
| `cmd1 |& cmd2`                   | Redirect stdout and stderr of cmd1 to stdin of cmd2(bash 4.0+ only).    |
| `cmd1 2>&1 | cmd2`               | The same as  `cmd1 |& cmd2`. Use this for older bashes.                 |
