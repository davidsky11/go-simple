# go-simple

## 介绍
**go 语言 练手工程**       
https://github.com/davidsky11/go-simple

### 用例来源
+ 《Go语言101》
+ 《GO语言开发实战 慕课版》
+ [A Tour of Go](https://go.dev/tour/list)
+ [ Go By Example ](https://github.com/everyx/gobyexample)

## go 说明
### go version
+ ~~go1.18.4~~ [❌]
+ go1.19.1 [✔️]

### 组织方式
+ 使用 go module

### 单元测试
+ go testing
+ 针对单个package，只构建一个单元测试文件，命名：pkg_[pkgName]_test.go

----------------------------

## 如何学习 Go 
👉 系统的学习基础（官方文档、优秀的书籍、博文） + 边学边练                     
入门就看《The Go Programming Language》，最好的学习方法就是把书上的例子都抄一遍。可以翻翻优秀的 Go 项目来学习他们是如何使用这门语言的，前面也提到了，awesome-go 里有很多这样的项目。直到你觉得 Go 的所有语法都了熟于心，就算是入门了吧。

👉 线上项目 + 并发编程 + 底层知识（源码） + 定位问题           
去学习怎么定位 Go 在线上系统的问题，成为一个 Go 的高级工程师。这部分需要了解一些 Go 的底层知识，学习基于 goroutine 和 channel 的各种并发编程模式，以及常用的工具链：比如 pprof 怎么用，怎么用 --base 去找内存泄露，出了性能问题怎么做优化等等。要达到的目标是：线上的 Go 系统出了问题，能够通过自己的知识储备快速定位。

👉 参与社区互动，接触优秀的人，尝试回答别人的问题            
通过多接触各种场景（比如大流量，高并发的，业务的，基础设施的等等），同时与其它语言横向做对比，了解 Go 语言在各种场景下的优缺点，不要成为一个语言原教旨主义者，比如我在工作的过程中就看到过不少 Go 其实就没法应付的场景，大家硬着头皮用，硬着头皮 hack，项目搞到最后优化起来也很痛苦，可能还不如直接去用 Rust。尽量多思考，也不要忌讳与其它语言的熟手交流。

## 学习资料
**入门：**     
+ 《GO语言开发实战 慕课版》         
+ [《Go by Example 中文》](https://books.studygolang.com/gobyexample/)

**进阶：**     
+ 《GO语言高级编程》     
+ 《Go语言101》    
+ 《Mastering Go》第二版  [看云](https://www.kancloud.cn/cloud001/golang/1601804)    [gitbook](https://hantmac.gitbook.io/mastering-go-second/)   

**底层：**      
+ 《Go 设计与实现》

**电子书：**       
+ [《GO语言高性能编程 》](https://geektutu.com/post/high-performance-go.html)

**微信公众号：**
+ 貘艺工作室 Go 101
+ Go语言中文网
+ Go Official Blog

**知识星球：**       
+ Go 语言研习社




Go 库列表：
+ testify                    
  `testify`可以说是最流行的（从 GitHub star 数来看）Go 语言测试库了。testify提供了很多方便的函数帮助我们做assert和错误信息输出。使用标准库testing，我们需要自己编写各种条件判断，根据判断结果决定输出对应的信息。           
  testify核心有三部分内容：
  - assert：断言；
  - mock：测试替身；
  - suite：测试套件。
+ testing
  `testing` 是 Go 语言标准库自带的测试库。在 Go 语言中编写测试很简单，只需要遵循 Go 测试的几个约定，与编写正常的 Go 代码没有什么区别。Go 语言中有 3 种类型的测试：单元测试，性能测试，示例测试。





```bash
# go get -u github.com/stretchr/testify
```