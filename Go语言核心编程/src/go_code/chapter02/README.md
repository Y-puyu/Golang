# GO语言核心编程 chapter2

`chapter1` 主要是 `golang` 语言的发展背景、学习方向、应用领域的简单介绍，强调了学习这门课程的**学习方法**。

`chapter2` 主要是 `go` 语言的简单概述、搭建开发环境，简单讲解了转义字符和如何查看官方文档，最后简单讲解了 `Dos` 命令(感觉在 `vscode` 下直接运行也很香)

`chapter1` 相关内容：

![20201117144410](https://raw.githubusercontent.com/Y-puyu/picture/main/images/20201117144410.png)

`chapter2` 相关内容：

![20201117144503](https://raw.githubusercontent.com/Y-puyu/picture/main/images/20201117144503.png)
![20201117144525](https://raw.githubusercontent.com/Y-puyu/picture/main/images/20201117144525.png)

---

## **两个重要的在线学习网站：**

[**Go语言中文网**](https://studygolang.com/)

[**Go语言中文网：Golang标准库文档**](https://studygolang.com/pkgdoc)

---

## **重点知识点回顾与总结：**

**Go 语言的 SDK 是什么？**

- SDK(Software Development Kit 软件开发工具包)，做 Go 开发首先需要安装并配置好 SDK
- Go 官方网站可能被和谐掉，推荐 Golang 中国，或者其它门户网站进行下载。

**Golang 环境变量的配置及作用：**

- GOROOT：指定 go sdk 安装目录
- path：包含三个重要指令 sdk\bin 目录，其中包括 go.exe godoc.exe gofmt.exe
- GOPATH：go 项目的工作目录，所有项目的所有源码都放在该目录下

**Golang 程序的编写、编译、运行步骤？一步执行的方法：**

- 类比 gcc 与 g++
- go build test.go 生成一个 .exe 文件，将所依赖库添加到 .exe 文件中，对方可以直接使用该 .exe 文件。
- go build -o a.exe test.go
- go run test.go 直接运行，不生成 .exe 文件。对方电脑需要安装有 go 环境才可以使用该源码

**老师这边手把手配置了 Windows、Linux、macOS 下的 go 环境，十分详细。**

**学习项目文件的目录结构是很重要的。以往的代码量小，同时也没有注意将项目程序按源码、章节进行存放，导致整个项目结构比较混杂，是一个偷懒很不好的习惯。**

---

## 疑问

2020/11/17 15:12

通过配置 `GO111MODULE` 环境变量，**解决了我一直 package main 有红色波浪线的问题**。简单查阅了下，没怎么看懂。先囫囵吞枣了，往后学习看看能否解决。目前不影响使用。
