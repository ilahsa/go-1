package main

import "fmt"

func main() {
    p:=[]int{2,4,5,8,9,0}//[]T 会生成一个类型为T 的数组切片 slice，slice 零值 是nil
    s:=p[2:4]//slice可以重新切片生成一个新的slice，slice[m:]从m 到最后 slice[:n] 从0到n slice[m:n] 从m到n 不包含n
    fmt.Println(p)
    fmt.Println(s)
    v:=make([]int,5)//make 可以构造slice
    fmt.Println(v)
    //可传递第三个参数到make 以指定容量
    c:=make([]int,0,5)//第二个参数为数组长度，第三个参数为数组容量默认为第二个参数，容量用完会重新划分 原容量*2 的容量
    fmt.Println(len(c),cap(c))
    c=append(c,0,2)//添加一个元素第一个参数为slice其余参数添加到slice中
    fmt.Println(c)
}