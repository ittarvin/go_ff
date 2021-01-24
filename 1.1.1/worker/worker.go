package main

import "fmt"

//方式一
var counter int =0
var ch = make(chan int,1)

func worker(){

     ch <-1
     counter ++
     <-ch

     fmt.Printf("ch = %d",ch)
}

func main(){
   worker()
}
