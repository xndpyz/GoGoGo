## myGo

### 基础

|类型| 范围                       | 区别           |
|---|--------------------------|--------------|
|值类型| int,string,arr,float,bool | 栈stack       |
|引用类型| *,切片,chan,interface,map  | 堆heap,初始值为nil |

1. [结构体相关](base/struct.go)

### go协程相关

M,N 模型 GMP G:协程，P处理器，M线程 内存：几KB 调度器的设计策略 复用线程 work slealing机制，hand offj机制 利用并行 抢占 最多10ms 全局G队列

runtime.Goexit() 退出当前goroutine

1. [打印素数](goroutine/prime.go)
2. []

### 算法

[树]