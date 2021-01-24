package main

import "fmt"

func main(){

  // 声明管道, 这种方式声明的管道，值为 nil ，每个管道只能存储一种类型的数据。
  var nilV chan int

  fmt.Println("nilV := ",nilV)
  //写 --- 对于 nil 的管道，无论是读写都会阻塞，而且是永久阻塞
  //nilV <- 1

  //读
 /*  var readV <- nilV

  fmt.Println("readV <- nilV ",readV) */

  // 带缓冲管道
  ch := make(chan int ,10)

  ch <-1

  d := <-ch

  fmt.Println("带缓冲管道输出 ",d)


  // 无缓冲管道
  //ch2 := make(chan int)

  //管道最多可以复制给两个变量
  ch3 := make(chan int ,10)

  ch3 <-1

  v1 := <-ch3

  fmt.Println("复制给一个变量",v1)

  ch3 <-2

  //关闭管道
  close(ch3)

  //取出管道数据
  v2 := <-ch3

  fmt.Println("读出缓冲区数据",v2)

  //两个变量赋值
  x,ok := <- ch3

  // ok 表示是否读取成功，不指示管道的关闭状态
  fmt.Println("复制给两个变量",x,ok)

  /*
    总结：内置函数 len() 数据个数， cap() 缓冲区大型，close() 关闭管道。
  * 管道实现了一种 FIFO（先入先出）的队列，数据总是按照写入的顺序流出管道。
    协程读取管道时，阻塞的条件有：管道无缓冲区，管道缓冲区中无数据，管道的值为nil
    协程写入管道时，阻塞的条件有：管道无缓冲区，管道缓冲区已满，管道的值为nil
  */

}
