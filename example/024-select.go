package main

import "fmt"
import "time"

func fibonacci(c,quit chan int) {
    x,y:=0,1
    for{
        select {//select 语句可以阻塞多个goroutine 直到其中一个满足条件，当多个都准备好时候会随机选择一个
        case c<-x:
            x,y=y,x+y
        case <-quit:
            fmt.Println("quit")
            return
        default://其他分支都没准备好时候会执行 default 分支 为了非阻塞的发送或者接收可以使用default分支
            time.Sleep(50*time.Millisecond)
        }
    }
}

func main() {
    c:=make(chan int)
    quit:=make(chan int)
    go func () {//goroutine 运行一个匿名函数
        for i:=0;i<10;i++{
            fmt.Println(<-c)
        }
        quit<-0
    }()
    fibonacci(c,quit)
}