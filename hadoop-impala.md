---
layout: post
title: Hadoop Impala
date: 2015-09-08 10:30:00
tags:
- Hadoop
categories: Hadoop
---

# 1. Impala
Impala Shell
Impala Shell Options:

| Option                                                    |         Description                                   |
| --------------------------------------------------------- | ----------------------------------------------------- |
| `-i hostname` or `--impalad=hostname[:portnum]`           | Set the host name.                                    |
| `-f query_file` or  `--query_file=query_file`             | Specify the query file.                               |
| `-d default_db` or `--database=default_db`                | Set the default database.                             |
| `-o filename` or `--output_file filename`                 | Redirect all query results in the specified file.     |

`impala-shell -i <impala_host> -f <file_path> -d my_db`

# 2. Impala shell command

| Command                          |         Description                                            |
| -------------------------------- | -------------------------------------------------------------- |
| `connect`                        | Connect to a impala server.                                    |
| `use <db_name>`                  | Use the specified database.                                    |
| `show tables`                    | Specify the query file.                                        |
| `DESCRIBE [FORMATTED] table`     | Display metadata about a table or view.                        |


