package main

import "fmt"

func main() {
    age:=17
    if age<18{//条件表达式无需()
        fmt.Println("no you can't see this movie")
    }else{
        fmt.Println("oh you can see this AV")
    }
    if age+=1;age<18{//加特效呀。。。。duang if 表达式前可执行一条简单的语句,如果是声明变量则在else块中依然可用
        fmt.Println("no you can't see this movie")
    }else{
        fmt.Println("oh you can see this AV")
    }
}