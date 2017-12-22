---
layout: post
title: Amazon DynamoDB
date: 2017-05-26 11:10:00
tags:
- docker
categories: Java
---



# DynamoDB基本概念
* 表：类似于关系型数据库中的表
* 项目： 类似于关系型数据库的行
* 属性： 类似于关系型数据库的字段
* 主键：每个表都必须有主键。

# 主键
每个表都必须有主键，在创建表的时候就要指定主键。    
DynamoDB支持2种主键：
1. 分区键。即主键只有一个字段。
比如下面`Person`表的`PersonID`:
```json
{
    "PersonID":100,
    "Name":"Smith",
    "Phone":"12345"
}
```
2. 分区键和排序键，称为复合主键
比如下面`Music`表的`Artist`和`SongTitle`:
```json
{
    "Artist":"No One You Know",
    "SongTitle":"Somewhere Down the Road",
    "AlbumTitle":"Somewhat Famous",
    "Genre":"Country",
    "Year":1984
}
```

#  二级索引
对于下面的Music表，主键是Artist分区键和SongTitle排序键。我们可以按Artist或者Artist和SongTitle查询数据。如果还想按Genre和AlbumTitle查询，可以在这些属性上创建一个索引，然后来进行查询。
```json
{
    "Artist":"No One You Know",
    "SongTitle":"Somewhere Down the Road",
    "AlbumTitle":"Somewhat Famous",
    "Genre":"Country",
    "Year":1984
}
```
创建的二级索引如下所示:

```text

                      Music                                                  GenreAlbumTitle

{                                                                 {                       
    "Artist":"No One You Know",                                       "AlbumTitle":"Hey Now",
    "SongTitle":"My Dog Spot",                                        "Genre":"Country",
    "AlbumTitle":"Hey Now",                                           "Artist":"No One You Know",
    "Genre":"Country",                                                "SongTitle":"My Dog Spot"
    "Price":1.98,                                                 }
    "Year":1984,
    "CriticRating":8.4
}


{                                                                 {
    "Artist":"No One You Know",                                       "Genre":"Country",
    "SongTitle":"Somewhere Down the Road",                            "AlbumTitle":"Somewhat Famous",
    "AlbumTitle":"Somewhat Famous",                                   "Artist":"",
    "Genre":"Country",                                                "SongTitle":"Somewhere Down the Road"
    "Year":1984,                                                  }
    "CriticRating":8.4
}


{                                                                 {
    "Artist":"The Acme Band",                                         "Genre":"",
    "SongTitle":"Still in Love",                                      "AlbumTitle":"",
    "AlbumTitle":"The Buck Starts Here",                              "Artist":"",
    "Genre":"Rock",                                                   "SongTitle":""
    "Year":1984,                                                  }
    "Price":2.47,
    "PromotionInfo":{
        "RadioStationsPlaying":[
            "KHCR",
            "KQBX",
            "WTNR",
            "WJJH"
        ],
        "TourDates":{
            "Seattle":"20150625",
            "Cleveland":"20150630"
        },
        "Rotation":"Heavy"
    }
}

{                                                               {
    "Artist":"The Acme Band",                                       "Genre":"",
    "SongTitle":"Look Out, World",                                  "AlbumTitle":"",
    "AlbumTitle":"The Buck Starts Here",                            "Artist":"",
    "Genre":"Country",                                              "SongTitle":""
    "Price":0.99                                                }
}

```


# 链接
* [Introduction.html](http://docs.aws.amazon.com/zh_cn/amazondynamodb/latest/developerguide/Introduction.html)
* [new-geo-library-for-dynamodb/](https://aws.amazon.com/cn/blogs/aws/new-geo-library-for-dynamodb/)
* [geo-library-for-amazon-dynamodb-part-1-table-structure/](https://aws.amazon.com/cn/blogs/mobile/geo-library-for-amazon-dynamodb-part-1-table-structure/)
* [geo-library-for-amazon-dynamodb-part-2-geodatamanagerconfiguration/](https://aws.amazon.com/cn/blogs/mobile/geo-library-for-amazon-dynamodb-part-2-geodatamanagerconfiguration/)
* [geo-library-for-amazon-dynamodb-part-3-creating-amazon-dynamodb-tables/](https://aws.amazon.com/cn/blogs/mobile/geo-library-for-amazon-dynamodb-part-3-creating-amazon-dynamodb-tables/)
* [geo-library-for-amazon-dynamodb-part-4-put-geo-point/](https://aws.amazon.com/cn/blogs/mobile/geo-library-for-amazon-dynamodb-part-4-put-geo-point/)
* [DynamoDBLocal.html](http://docs.aws.amazon.com/zh_cn/amazondynamodb/latest/developerguide/DynamoDBLocal.html)
* [JavaDocumentAPIItemCRUD.html](http://docs.aws.amazon.com/zh_cn/amazondynamodb/latest/developerguide/JavaDocumentAPIItemCRUD.html)
