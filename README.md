# golang_sdudy
## 选择框架：martini
Martini框架是使用Go语言作为开发语言的一个强力的快速构建模块化web应用与服务的开发框架。Martini是一个专门用来处理Web相关内容的框架，其并没有自带有关ORM或详细的分层内容。所以当我们使用Martini作为我们的开发框架时，我们还需要选取适合的ORM等其他包。
### 说明：
w := martini.Classic()创建一个典型的martini实例。
w.Get("/", func() string { ... })接收对\的GET方法请求，第二个参数是对一请求的处理方法。
w.Run()运行服务器。
### martini框架优势
使用极其简单.
无侵入式的设计.
很好的与其他的Go语言包协同使用.
超赞的路径匹配和路由.
模块化的设计 - 容易插入功能件，也容易将其拔出来.
已有很多的中间件可以直接使用.
框架内已拥有很好的开箱即用的功能支持.
 
## curl测试

 ![image](https://github.com/lqAsuna/golang_sdudy/blob/master/image/res_1.png)
 
## ab测试

 ![image](https://github.com/lqAsuna/golang_sdudy/blob/master/image/res_2.png)
 ![image](https://github.com/lqAsuna/golang_sdudy/blob/master/image/res_3.png)
 
### ab测试结果说明

ab测试结果说明：

Server Software：服务器使用的软件

Server Hostname：服务器主机名（localhost）

Server Port：服务器端口(9090)

Document Path：文档路径(hello/your)

Document Length：文档长度(8 bytes)

Concurrency Level：并发(100)

Time taken for tests：测试所用时间(0.216)

Complete requests：完成了的请求数(1000)

Failed requests：失败了的请求数(0)

Total transferred：传输的字节数(127000 bytes)

HTML transferred：传输的 HTML 报文体的字节数(10000 bytes)

Requests per second：每秒请求数(8443.45r/s)

Time per request：每个请求花费的平均时间(12.345s/r)

Time per request：平法情况下每个请求花费的平均时间(0.547s/r)

Transfer rate：每秒传输的平均千字节数(1258.42)

Connection Times (ms) 传输时间统计，百分之五十的请求可以在12ms内得到回应，百分之九十五的请求可以在29ms内得到回应，最长的一个请求回应时间达到了45ms
