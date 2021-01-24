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
  ch2 := make(chan int)

  /*
    总结：
  * 管道实现了一种 FIFO（先入先出）的队列，数据总是按照写入的顺序流出管道。
    协程读取管道时，阻塞的条件有：管道无缓冲区，管道缓冲区中无数据，管道的值为nil
    协程写入管道时，阻塞的条件有：管道无缓冲区，管道缓冲区已满，管道的值为nil
  */

}
