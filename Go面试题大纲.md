# 面试题大纲

## 1、Redis面试题

```python
'''
Redis面试题
1，Redis的缓存用的什么键值对？
2，哈续命令里有一个H GIT ALL是什么命令？
3，Redis的锁怎么解决并发问题？
4，消息中间键，Mysql解决什么样的场景？
5，kafka了解吗？说一说
6，Mysql里app date是什么样的锁级别？
7，能说一下乐观锁和悲观锁吗？区别是什么？
8，场景题一：假设有A,B,C,D,E四个字段，有三个条件，1，A+B;2,B+C;3,C+D+E,请问加索引的条件（依据）是什么？
9，场景题二：TCP,IP在集成协议的哪一层？
10.定时删除、懒汉式删除、定期删除
11.redis压缩表 原理
		https://juejin.cn/post/6847009772353355783
		问题: 上面已经提到了需要用最大长度的字符串大小作为整个数组所有元素的内存大小，如果只有一个元素的长度超大，但是其他的元素长度都比较小，那么我们所有元素的内存都用超大的数字就会导致内存的浪费
		解决: 多大的元素使用多大的内存，一切从实际出发，拒绝浪费。
		
12.跳表 原理  
		https://redisbook.readthedocs.io/en/latest/internal-datastruct/skiplist.html
		和字典、链表或者字符串这几种在 Redis 中大量使用的数据结构不同， 跳跃表在 Redis 的唯一作用， 就是实现有序集数据类型。
'''
```

## 2、Docker面试题

```dockerfile
1、什么是虚拟化 
2、Docker与虚拟机的区别 
3、Docker中的三大概念 
4、什么是镜像
5、什么是容器 
6、什么是镜像仓库 
7、简要说明一下docker网络 
8、docker网络的原理有哪些 
9、Linux网络隔离的方式 
10、Linux不同网络之间互联的方式 
11、Docker网络模式 
12、Docker网络模式的实现原理 
13、怎么实现使用在宿主主机使用127.0.0.1访问容器
14、怎么共享一个容器的网络 
15、怎么实现一个容器完全的网络隔离 
16、怎么实现多对多网络互联 
17、创建一个容器 
18、随机映射端口 
19、固定映射端口 
20、挂载存储卷
21、为容器设置一个名字 
22、以守护进程方式运行一个容器 
23、构建一个Django容器
24、构建一个NGINX容器 
25、使用NGINX代理Django
```



## 3、Go面试题

```go
// 1.七牛云面试题
/*
	1、hash表的扩容设计
	2、有100W个cpu密集型任务需要最快时间执行完，需要一个设计方案
	3、有100W个任务，有的1秒执行，有的一年后执行，设计方案是什么？
	4、详细聊项目，细聊，回聊很多分布式的内容，如何保证数据一致性和高并发的能力，宕机后的回复能力，分别聊了redis、mysql和rabbitMQ，分布式后宕机了怎么保证服务可用和灾备准备情况，宕机后数据丢失了怎么复原。
	5、Go语言的slice作为函数参数有什么问题？
	6、gmp模型原理
	7、netpoll网络模型原理
	8、channel原理
	9、goroutine原理
	10、context原理
	11、消息队列中间件，牛逼的开源项目了解多少？有没有了解过他的实现原理和源码？怎么做分布式的？
	12、make 和new的区别  都是初始化变量用的给变量分配内存地址  还有就是new初始化变量返回的是这个类型的指针  make返回类型本身
	13、Go的内存管理
	14、context怎么用？
		- 1.可以存一些token信息 键值对存储的
		- 2.日志的全链路追踪 就是从amp获取到teacei塞进context里面 函数的第一个参数一定要传context这个参数就行了 
		- 3.取消上下文 设置超时时间
	15、数组和切片的区别？
				数组是定长的  切片是可以扩容的  然后讲下切片的扩容机制  最后在讲下数据底层怎么实现扩容的  
  16、Map当函数的参数传入，并且函数内修改map的值，外面的map是原来的还是会改变？
   			外边的map数据会发生变化，因为在golang中 所有都是值传递，map会改变是因为map底层有个makemap这个方法 这个方法的返回值是map的指针
	17、切片当函数参数传入会发生什么？
        不会对外部的切片产生影响还是原值，函数内部会改变，因为golang都是值传递
	18、定义一个切片占用多少字节？
			- 占用24个字节 因为切片在底层有个sliceHead的结构体 这个结构体里面有3个变量  
				- 1.data 指向的是底层数组的内存地址 类型是uintptr uintptr占用的字节大小是8  
				- 2.len 指的是切片的长度 类型是int类型  占用字节是8    注释：int类型 uint类型占用字节都是8 并且是在64位操作系统上，32为操作系统是4个字节
       	- 3.cap 指的是切片的容量  类型为int 占用字节为8
	19、定义一个字符串他占用的字节大小是多少？
      在golang里面string在运行时是基于stringHead的  stringHead的数据结构是data 类型为unitptr 表示底层的内存地址   还有len 表示字符串的长度   所以string为16个字节
	20、strct结构体占用多少字节？
    	具体要看结构体里面的类型 如果空结构体占用字节为0
	21、atomic.value 自行百度 很重要

*/

// 2.Go面试题链接
/*
	https://github.com/KeKe-Li/data-structures-questions/blob/master/src/chapter05/golang.01.md#Golang%E4%B8%AD%E9%99%A4%E4%BA%86%E5%8A%A0Mutex%E9%94%81%E4%BB%A5%E5%A4%96%E8%BF%98%E6%9C%89%E5%93%AA%E4%BA%9B%E6%96%B9%E5%BC%8F%E5%AE%89%E5%85%A8%E8%AF%BB%E5%86%99%E5%85%B1%E4%BA%AB%E5%8F%98%E9%87%8F
	https://studygolang.com/search?q=%E9%9D%A2%E8%AF%95
	https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-netpoller/
*/

// 3.理论题
/*
	http://wen.topgoer.com/docs/gozhuanjia/gozhuanjiaxiecheng2
	https://blog.csdn.net/qq_44756792/article/details/103851500
	https://studygolang.com/search?q=%E9%9D%A2%E8%AF%95
*/

// 4.面试准备
/*
### golang语言基础
​	http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/
### golang数据结构底层原理学习资料
​	https://golang.design/go-questions/
### golang修养之路 
Ps: gc、已经golang gpm调度底层原理s
​	http://www.topgoer.cn/docs/golangxiuyang/golangxiuyang-1cmedc59gtgpi
### golang context底层原理解析资料
[由浅入深聊聊Golang的context_咖啡色的羊驼-CSDN博客](https://blog.csdn.net/u011957758/article/details/82948750)
### golang gpm调度底层原理解析视频
[Golang深入理解GPM模型_哔哩哔哩 (゜-゜)つロ 干杯~-bilibili](https://www.bilibili.com/video/BV19r4y1w7Nx)
### golang gc原理视频
[Golang中GC回收机制三色标记与混合写屏障_哔哩哔哩 (゜-゜)つロ 干杯~-bilibili](https://www.bilibili.com/video/BV1wz4y1y7Kd)
### golang grpc底层源码分析视频
[#17 grpc 开发及 grpcp 的源码分析 【 Go 夜读 】_哔哩哔哩 (゜-゜)つロ 干杯~-bilibili](https://www.bilibili.com/video/BV1ht41187Wh)
### golang channel 底层原理解析视频
[#56 channel & select 源码分析 【 Go 夜读 】_哔哩哔哩 (゜-゜)つロ 干杯~-bilibili](https://www.bilibili.com/video/BV1g4411R7p5)



golang面试篇

​```
推荐面试之前多刷刷牛客网的golang面试题
https://www.nowcoder.com/discuss/experience?tagId=2656
​```





*/
```



## 4、Python面试题

```python
# Python面试题 (一)
'''
1) 用Python读取50万行内容的 excel 文件入库，怎么才能提高导入数据库速度？
3) django中间件从上到下执行了哪些内容？
4) 上有没有搭建过mysql或redis读写分离，说说搭建过程。
5) 三次握手，四次握手实现原理。
6) docker制作镜像流程，docker-compose编排制作完整项目一键启动
7) linux常用命令
8) uwsgi，WSGI，ASGI，如何拉起服务器的。
9) 在项目中哪里使用到一般框架，定时任务。
10) 讲讲redis的持久化方式rdb和aof。
11) 讲讲session，cookies，token的区别
13) redis单线程下为什么那么快
14) mysql事物
15) 如何解耦合开发api
16) MySQL行锁是什么？
17) 介绍下IPC进程间实现通信。
18) Django中哪里用到了元类？
19) 如何反向输出单链表？
20) 如何实现二叉树的层次遍历？
21) 协程是如何实现的？
22) 为什么要有GIL锁？
23) 事物的隔离机制有哪几个？
24) 你们现在服务器有几台？功能服务器有几台？如果服务器挂挂了怎么办？
25) 公网用的是ip还是域名？
26) 域名用的是http还是https？
27) 是直接把域名的流量转到nginx上面吗？
28) 域名对应的ip有几个？为什么是多个？
29) 域名到nginx的链入机制？
30) nginx有几台服务
'''

# 纳纳科技面试题（二）
'''
1.数据库的话，在之前的公司数据量是多少？有多少行数据？
2.有没有做过数据库索引的优化？
3.如果有很大的数据量，50万行左右的数据，其中有一个字段是标记他是否是删除的，有一个场景是要搜索被删除的人数是多少？为了提高查询速度，是否要对判断的这个字段添加索引。
4.索引有单个索引、联合索引、那个单个索引比较好，还是多个索引在一起比较好？什么样的应用场景用联合索引好一点？
5.有一个联合索引设置的数据是abc，有一个字段的查询是a=1 and c=2，这种情况下能不能用到索引？
6.redis常见的数据类型
	set、zset、list、string、hashmap
7.工作中用redis的哪种数据类型比较多？
8.zset有没有用到过？
9.字典类型哪种场景下会用到？为什么不用string类型？
10.项目中单例模式的最常见的场景
11.你做的项目中觉得最成功的，比较突出你能力的是哪个？
'''

# 面试题（三）
'''
1. 并发量情况写你们项目用的是什么锁
2. 索引优化
3. 分布式事务你们怎么处理的
4. 订单量很大时，我要做一个模糊匹配，该怎么优化
5. 我现在定义一个联合索引（A，B，C），我如果先查C，再查B，再查询A，会先走哪个，对效率有没有影响？什么是最左前缀匹配原则？
6. 做过主从搭建吗？主从同步的过程中，主从一致性怎么保证？
	https://www.jianshu.com/p/328ad87bde5e
7.布隆过滤器，什么原理
8.讲讲ES的倒排索引
9.项目中那部分模块做的最久，最熟练，有没有哪些核心模块
10.你们用的表结构是怎么设计的
11.项目数据库框架，有没有缓存，数据库有没有主从
12.Redis有没有做集群，还是单点
13.冒泡排序、快速排序、堆排 各是什么原理
14.如果倒置列表，你将使用什么方法
15.Python是动态的还是静态语言，怎么理解的
16.怎么优化sql语句
17.项目的日活、并发
18.Mysql和redis的版本
19.分布式id生成，分布式锁，rabbitmq，双写一致性，负载均衡
20.设计模式
'''
```



### Python项目

```python
'''
妙小程项目
> miaocode 
涉及技术：Django rest framwork、vue、mysql、redis、nginx、uwsgi 
项目描述： 该项目是结合django中drf框架以及前端vue框架实现的企业官网。实现了集成学员开发环境、学员作品集锦、课程销售市场、个人课程管理等一系列功能。 
1、登陆注册模块：使用短信验证码方式实现账号注册，jwt签发校验用户token，并使用认证频率双限制，避免恶意攻击，同时集成第三方平台的登录注册。
2、首页轮播图模块：celery定时更新redis，减少对数据库压力 
3、作品展览馆模块：使用缓存处理推荐作品，实现页面静态化
4、搜索模块：结合haystack + ES进行全文检索，搜索学员对应的作品信息 
5、订单模块：采用rabbitmq消息队列+celery异步处理子订单，实现购物车、接口幂等一系列并发场景的开发，以及发送短信等功能 
6、支付接口：调用支付宝、微信支付
7、接口文档：使用drf的core api库自动生成接口文档 

> kk课程管理系统 
涉及技术：Django、MySQL、admin、celery、nginx、uwsgi
方便公司内部工作的开展，使用Django框架结合自带的admin组件重构一个后台管理组件。实现了用户管理、作品管理、销售管理、流量课程管理、教务管理等一系列功能。
1、登录模块：用户账号由员工入职后统一分配，通过图形验证码实现用户登录 
2、公告栏模块：将每次版本更新，通过公共栏告知所有员工 
3、权限管理：将不同用户与权限进行关联，实现基于角色的访问控制 
4、教务管理模块： 
- 课程信息管理 - 课程表管理 - 课次管理 - 教务配置

> 妙小程pc端官网 
涉及技术： Flask、MySQL、redis、nginx、uwsgi
项目描述：用于宣传公司企业文化，以及引流。
1、后端采用flask框架，提高运行效率及体验 
2、使用SQLAlchemy进行数据库表的设计 
3、前端结合bootstrap框架及自定义内容进行搭建 
4、通过ajax向后台提交数据 
5、采用flask-admin管理后台数据

> 日志分析报表
根据公司业务需求开发报表系统 
涉及技术：python、flask、MySQL、hightcharts 
1、分析nginx访问日志、自定义生成日志文件
2、统计日志数据存入mysql 
3、结合flask生成路由、渲染图表页面 
4、利用highcharts生成实时监控数据图表

> 老师考核统计 
顺应公司发展，为教学部门结合数据分析开发的一个简易考核平台
涉及技术：python 、Django、bootstrap、pandas、numpy、matplotlib 
1、结合Django和bootstarp搭建一个简易数据收集平台 
2、利用pandas和numpy统计分析，每位老师的周转化率、月转化率 
3、将分析结果通过matplotlib和seaborn展示成图表
'''
```



## 5、ELK

### ES

```cmd
# elasticsearch：分布式全文检索引擎，倒排索引，集群搭建，脑裂，打分机制，数据分片
```



## 6、MySQL

```mysql
-- A小课直播 《海量数据高并发MySQL应用实战》
Day1 基础篇：深入理解MySQL存储引擎工作原理和优化思路
第一天直播链接：https://appRZ8ZZTzY8571.h5.xeknow.com/st/4GjBU3qcP

Day2 进阶篇：掌握MySQL事务管理和锁机制
第二天直播链接:  https://appRZ8ZZTzY8571.h5.xeknow.com/st/7mcqQ7TW2

Day3 实战篇：MySQL分库分表综合实战
第三天直播链接：https://appRZ8ZZTzY8571.h5.xeknow.com/st/17Svfxppg

专栏回放链接   https://appRZ8ZZTzY8571.h5.xeknow.com/st/76aJZoT1q

```



## 你为什么离开你上一家公司

```python
# 严禁回答: 工资低, 与同事有矛盾,           
# 整体思路: 遇到瓶颈, 寻求个人更好的发展。 但是这句话不能直接讲出来
# 关键点: 绝对不能说上一家公司的坏话、不能说上一任老板的坏话、不能说和同事不和、 谁也不能得罪
# 回答模板: 我负责的那个项目, 到现在已经比较稳定了, 公司近期在项目方面就没有太多新的开发的、推进的计划。我个人就是觉得说 短期内公司的业务, 对我来讲, 在我的职业生涯上的技术来讲 可能帮助不是特别大, 并且我又深刻的认识到, 虽然我之前公司的环境啊氛围啊特别好, 但是呢感觉公司的开发呢 因为公司的规模毕竟比较小嘛, 好多的互联网上的一些新的东西 没有机会用, 我个人又是对这些新技术特别的热爱, 我就想能有一个机会到更专业的一个环境里, 能再提高我自己。
```







![l14dd81908b191f4deac02c774ead3730-s-mca796c3b35fd08a8eaf3de6e477ece30](/Users/tank/Library/Containers/com.tencent.xinWeChat/Data/Library/Application Support/com.tencent.xinWeChat/2.0b4.0.9/e096ae970eb4892333f77d2ad91b8f70/Favorites/data/l14dd81908b191f4deac02c774ead3730-s-mca796c3b35fd08a8eaf3de6e477ece30.jpg)



