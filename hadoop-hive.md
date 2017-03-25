---
layout: post
title: Hadoop Hive
date: 2015-09-08 10:00:00
tags:
- Hadoop
categories: Hadoop
description: The tutoria will describe the useage of Hadoop Hive.
---

# 1. Hive Command-Line Interface
run `hive` to enter Hive command-line interface.

# 2. Run a Script File by Hive
You can run below command to run a file by Hive.      
`hive -f <file_path>`

# 3. Create Tables
**Create external table**
```sql
CREATE EXTERNAL TABLE IF NOT EXISTS MY_TABLE(
`NAME` STRING,
`AGE` INT,
`PHONE_NUMBER` STRING,
`ID_NUMBER` STRING,
`UPDATE_TIME` TIMESTAMP)
ROW FORMAT DELIMITED
FIELDS TERMINATED BY ','
ESCAPED BY '\\'
STORED AS TEXTFILE;
```

```sql
use my_db;
ALTER TABLE table_name [PARTITION (partition_spec)] SET LOCATION 'hdfs_path_of_directory';
```

```sql
ALTER TABLE logs SET LOCATION 'hdfs://user/darcy/logs/2012/12/18';
ALTER TABLE logs PARTITION(year = 2012, month = 12, day = 18) SET LOCATION 'hdfs://user/darcy/logs/2012/12/18';
```

```sql    
TBLPROPERTIES ('serialization.null.format' = '');
```


# 4. Resources

[Hive internal and external table](http://litten.github.io/)
