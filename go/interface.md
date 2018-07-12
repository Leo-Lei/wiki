---
layout: post
title: Go interface
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---

# 定义一个接口
```go
type repository interface{
    save(s string)
}
```
# 实现接口
go中实现接口有两种方式:
1. 使用值类型接受者实现接口
2. 使用指针类型接受者实现接口
两者在实现接口接口本身，区别不大，只是接受者一个是值，一个是指针。但是在传参给接口时，会有区别。Go语言初学者经常遇到这个陷阱。

**1. 使用值接受者**
```go
func (mondo mongoRepo) save(s string) {
	fmt.Println("Mongo save",s)
}
```

**2. 使用指针接受者**
```go
func (mysql *mysqlRepo) save(s string) {
	fmt.Println("Mysql save",s)
}
```

# 使用接口
```go
func saveData(repo repository){
	repo.save("hello")
}
```
```go
mongo := mongoRepo{}
saveData(mongo)
saveData(&mongo)

mysql := mysqlRepo{}
//saveData(mysql)     // 编译错误
saveData(&mysql)
```


|  实现接口时的方法接收者  |      可接受的参数类型     |          
| ---------------------- | ------------------------ |
| `(t T)`                | `T` 和 `*T`              |
| `(t *T)`               | `*T`                     |





