package main

import "fmt"

func sum(a []int,c chan int) {
    sum:=0
    for _,v:= range a{
        sum+=v
    }
    c<-sum//向c中写入值
}

func main() {
    a:=[]int{7,2,8,-9,4,0}
    c:=make(chan int)//channeld 必须先make 
    go sum(a[:len(a)/2],c)//两个线程
    go sum(a[len(a)/2:],c)//同时执行
    x,y:=<-c,<-c//从c中获取值
    fmt.Println(x,y,x+y)//处理执行结果
}