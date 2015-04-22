package main

import "fmt"

func main() {
    sli:=[]int{234,22,45,5}
    for i,v:=range sli{//range可配合for 遍历 array slice map
        fmt.Println(i,"-",v)
    }
    for _,v:=range sli{//可通过_忽略索引 或者 值
        fmt.Println(v)
    }
}