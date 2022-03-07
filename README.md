## myGo

### 基础

|类型| 范围                       | 区别           |
|---|--------------------------|--------------|
|值类型| int,string,arr,float,bool | 栈stack       |
|引用类型| *,切片,chan,interface,map  | 堆heap,初始值为nil |

1. [结构体相关](base/struct.go)
  空结构体的三种应用场景(方法接受者,集合类型,空通道)
  作用：节省空间， 每个值都是有一定宽度的
2. string和bytes的转换[string&base](base/string.go) 
3. sync
4. 
### go协程相关

M,N 模型 GMP G:协程，P处理器，M线程 内存：几KB 调度器的设计策略 复用线程 work slealing机制，hand offj机制 利用并行 抢占 最多10ms 全局G队列

runtime.Goexit() 退出当前goroutine

1. [打印素数](goroutine/prime.go)
2. []

### 算法

[树]

### grpc
https://grpc.io/docs/languages/go/quickstart/
protobuf
保留使用 reversed
枚举 enum option allow_alias=true 允许相同值 第一个要为0 

oneOf 指针类型

定义服务
流式传入
流式返回

### gin
bind 还可以验证为空 自定义规则
中间件原理，自定义中间件
gorm 预加载reload 
jwt json web token
headers cliams signature


### leetcode
滑动窗口
1.二分法查找数
2.最长不重复字符串
3.最长回文字符串，双指针, eg 翻转字符串
4.字符串包含strStr， 滑动窗口解决, hello => ll
5.回文数 15551=> 150, 15,怎么求出15 /= 10, for i < j { i = i*10+ j%10} 
6.是否为有效字符串，{}[], 没有考虑一开始就错误的情况eg.]{, 仔细冷静思考
7.两数之和
8.利用缓存求斐波那契数列
9.动态规划求能爬楼梯的最小费用， 学会状态转移
10.杨辉三角, dp[i][0] = 1, dp[i][l] = 1
11.邻家打劫， 邻家打劫II, dp[1:], dp[:len(dp)]
12.衍生出来的，去除求最大数
13.颜色排序

排序：sort.Intn()
crtl+U 插入图片
ctrl+e 最近文件

##知识点
什么时候传一个结构体
传结构体节省内存, 涉及宽度， ch := make(ch chan struct{})


## 明日计划
grpc 入门学习
出现导入包错误
go mod init
go mod tidy
go mod download 解决
leetcode 10道题
使用goroutine 实现文件下载
Accept-Ranges 范围请求

## 未解决问题
1. 环形数组最大和
2. 子数组最大乘积