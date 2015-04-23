package main

import(
    "fmt"
    "time"
)

func say(s string) {
    for i:=0;i<5;i++{
        time.Sleep(100*time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("Go")//开启一个由go运行时环境管理的轻量级线程
    say("Hello")
}