---
layout: post
title: SQL Basis
date: 2015-08-12 14:00:00
tags:
- SQL
categories: SQL
---

# 1. Standard SQL 
## 1. CASE WHEN Clause
```sql
select
CASE productLine
	WHEN 'R' THEN 'Road'
	WHEN 'M' THEN 'Mountain'
	WHEN 'T' THEN 'Touring'
	ELSE 'Others'
END as productLine
from table
```
```sql
select 
CASE 
    WHEN ListPrice = 0 THEN 'zero'
    WHEN ListPrice < 50 THEN 'Under 50'
    WHEN ListPrice >= 50 and ListPrice < 250 THEN 'Between 50 and 250'
    WHEN ListPrice >= 250 and ListPrice < 1000 THEN 'Between 250 and 1000'
    ELSE 'Over 1000'
END as ListPrice
from table
```

# 2. SQL Dialect
## 2.1 NVL & isnull
Oracle/PLSQL: NVL Function:

## 2.2 DECODE
Oracle/PLSQL: DECODE Function:           
**SYNTAX**               
`DECODE( expression , search , result [, search , result]... [, default] )`
**EXAMPLE**           
This example decodes the value supplier_id. If supplier_id is 100, then the function returns 'IBM'; if supplier_id is 101, then it returns 'Microsoft';if supplier_id is 102, then it returns 'Hewlett Packard'. If supplier_id is not 100, 101 or 102, then the function returns 'Others'.
```sql
SELECT supplier_name,
DECODE(supplier_id, 100, 'IBM',
                    101, 'Microsoft',
                    102, 'Hewlett Packard',
                    'Others') as result
FROM suppliers;
```


# 3. PIVOT

## 3.1 Examples

| Course    |    Year   |    Earning    |
| --------- | --------- | ------------- |
| .NET      | 2012      | 10000         |
| Java      | 2012      | 20000         |
| .NET      | 2012      | 5000          |
| .NET      | 2013      | 48000         |
| Java      | 2013      | 30000         |

**Pivot the Course column:**
```sql
select *
from CourseSales
PIVOT(
SUM(Earning)
FOR Course IN ([.NET],Java)
) AS PVTTABLE
```

| Year      |    .NET   |       Java    |
| --------- | --------- | ------------- |
| 2012      | 15000     | 20000         |
| 2013      | 48000     | 30000         |

**Pivot the Year column:**
```sql
select * 
from CourseSales
PIVOT(
SUM(Earning)
FOR Year IN (2012,2013)
) AS PVTTABLE
```
| Course    |    2012   |       2013    |
| --------- | --------- | ------------- |
| .NET      | 15000     | 48000         |
| Java      | 20000     | 30000         |

## 3.2 Tips

To execute the PIVOT clause, except for the pivot column and the aggregation column, order by all other columns. Then we will get a result set group by some columns. Then the column values in the FOR clause will becomes the column headings.
