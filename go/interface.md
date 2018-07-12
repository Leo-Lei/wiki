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

# 使用接口(接口方法集)
假设有如下的方法，方法接收一个repository接口
```go
func saveData(repo repository){
	repo.save("hello")
}
```
下面的代码中实例化了一些对象，然后把它传递给saveData方法:
```go
mongo := mongoRepo{}
saveData(mongo)
saveData(&mongo)

mysql := mysqlRepo{}
//saveData(mysql)     // 编译错误
saveData(&mysql)
```
**注意**: saveData(mysql)出现了编译错误，因为我们在定义mysql实现repository接口时，使用的是指针接受者。在调用saveData(repo repository)方法时，必须要传递一个mysqlRepository的指针。不能穿mysqlRepository的值。

可以这么来理解:
1. 使用指针接受者实现接口时，只有指针类型实现了该接口，值类型没有实现该接口。
2. 使用值类型接受者实现接口时，值类型和指针类型都实现了该接口。
以上的规则也被称为是Go的方法集，如下面的表格:

|  实现接口时的方法接收者  |    值和指针哪些实现了该接口    |          
| ---------------------- | ---------------------------- |
| `(t T)`                | `T` 和 `*T`                  |
| `(t *T)`               | `*T`                         |

