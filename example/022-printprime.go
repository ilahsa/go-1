package main

import "fmt"

func main() {//打印1-100 的素数10个一行
    times:=0//输出次数
    for i:=2;i<100;i++{//遍历
        isprime:=true//是否素数
        for j:=2;j<i;j++{//循环判断因子
            if i%j==0{//有符合因子，则不是素数
                isprime=false
                break//跳出循环
            }
        }
        if isprime {//是否打印
            fmt.Print("\t",i)
            times+=1
        }
        if times%10==0{//是否换行
            fmt.Print("\r\n")
        }
    }
    fmt.Print("\r\n")
}