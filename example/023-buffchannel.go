package main

import "fmt"

func main() {
    ch:=make(chan int,5)//缓冲channel 只有缓冲区满的时候 才会阻塞 第二个参数 5 是缓冲区的长度
    //缓冲区填满的时候 填充阻塞  缓冲区清空的时候接收阻塞
    ch<-4
    ch<-5
    ch<-6
    ch<-3
    ch<-2
    close(ch)//只有发送者才可以关闭一个channel以告诉接受者channel已经关闭
    for i:=range ch {//从通道不断读取数据
        fmt.Println(i)
    }
}