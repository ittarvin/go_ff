package main

import "fmt"

func ChanCap(){
ch := make(chan int,10)
ch <-1
ch <-2
fmt.Println(len(ch))
fmt.Println(cap(ch))
}

func main(){
ChanCap()
}
