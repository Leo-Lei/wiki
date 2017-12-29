---
layout: post
title: Hadoop HDFS
date: 2015-09-08 10:35:00
tags:
- Hadoop
categories: Hadoop
---

# 1. Base HDFS Command

The hdfs command is in the format of `hadoop fs -<command> [-params]`, for example, the command to create a directory is `hdfs fs -mkdir path/to/folder`.

Below are the common used commands, and the prefix `hdfs fs` has been removed to simplify the command.
Please be aware that, the hdfs command is very similar to the general Linux command.

| Command   |             Description                      |      Syntax                                  |
| --------- | -------------------------------------------- | -------------------------------------------- |
| mkdir     | make a directory                             | `hadoop fs -mkdir <path>`                    |
| ls        | list the contents of a directory             | `hadoop fs -ls <path/to/directory>`          |
| rm        | remove a file or directory                   | `hadoop fs -rm <file_name>`                  |
| chmod     | change permission of a file or directory     | `hadoop fs -chmod <permission> <file_name>`  |
| get       | Download files from hdfs to local file system| `hadoop fs -get <hdfs_src> <local_path>`     |
| put       | Upload files from local file system to hdfs  | `hadoop fs -put <local_src> <hdfs_path>`     |
| help      | Get list of supported commands               | `hadoop fs -help`                            |
| cat       | Copies source paths to stdout                | `hadoop fs -cat <hdfs_file>`                 |

1. **mkdir**
Options:
* -p: the `-p` behavior is much like Unix `mkdir -p`, creating parent directories along the path.
Examples:
* `mkdir -p path/to/some/folder/foo/bar` will create all the parent directories.
