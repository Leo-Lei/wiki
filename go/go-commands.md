---
layout: post
title: Go Command
date: 2017-03-08 15:30:00
tags:
- Linux
categories: Linux
---


|                   Command                     |                             Desc                                        |
| --------------------------------------------- | ----------------------------------------------------------------------- |
| `go build`                                    |                                                                         |
| `go install`                                  |                                                                         |
| `go test -c`                                  | 编译test二进制文件,但不运行单元测试                                        |
| `go test -c -o hello.test`                    | 编译test二进制文件,但不运行单元测试。二进制文件是hello.test                 |





```go
func Test_Json1(t *testing.T) {

	f, err := os.Create("cpu-profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	json1()
	pprof.StopCPUProfile()
}

func Test_Json2(t *testing.T) {

	f, err := os.Create("cpu-profile2.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	json2()
	pprof.StopCPUProfile()
}

func json1() {
	s := `{"name":"tom","age":10,"courses":[{"name":"english","score":80},{"name":"math","score":90},{"name":"math","score":70}]}`

	for i := 0; i < 10000; i++ {
		student := &Student{}
		json.Unmarshal([]byte(s), student)
		fmt.Println(*student)
	}
}

func json2() {

	s := `{"name":"tom","age":10}`

	for i := 0; i < 10000; i++ {
		student := &Student{}
		json.Unmarshal([]byte(s), student)
		fmt.Println(*student)
	}
}

type Student struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Courses []*Course `json:"courses"`
}

type Course struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

```

* 使用`go test -c`编译test为二进制文件profile.test。但不运行测试。
* 使用`go test -test.run=Test_Json1`运行测试Test_Json1。会生成cpu-profile.prof文件。
* 使用`go tool pprof profile.test cpu-profile.prof`。进入命令行交互模式。
* 使用`top`显示消耗cpu时间片最多的方法。
* 使用`top -cum`显示累计消耗cpu时间片最多的方法。
* 使用`list Test_Json1`显示Test_Json1方法哪一行消耗cpu时间片最多。




如果想知道，应用或者一个模块，或者一个方法中哪个环节的性能，可以使用上面的方式，runtime/pprof。定位热点。
如果一个功能有好几种实现，想比较一下哪个实现方式的性能更好，可以使用benchmark测试。




* 如果某个功能有多个实现方案，想知道哪个性能比较好，那这时候就不太适合用cpu profile了。可以独立运行这几个测试，比较测试消耗的时间。
* 如果想知道某一个应用，模块，类，方法运行时，哪部分比较耗时，可以使用cpu profile。查看在运行的这段时间里，哪个代码块运行的比较慢。可以先定位到方法，再使用list查看方法的哪一行比较耗时。




