package main

import(
    "fmt"
    "runtime"
    "time"
)

func main() {
    fmt.Println("Go runs on ")
    switch os:=runtime.GOOS; os{//类似if for switch 表达式前也可执行语句
    case "darwin"://除非以 fallthrough 语句结束，否则分支会自动终止。
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux.")
    default://默认情况
        fmt.Println("other os.")
    }
    t:=time.Now()
    switch {//没有条件的switch语句，可用来更清晰的书写 if then else
    case t.Hour()<12:
        fmt.Println("Good morning")
    case t.Hour()<17:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
}