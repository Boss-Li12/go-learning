package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓冲区的读取和写入永远不死锁
	ch1 := make(chan int, 5)
	go pump(ch1) // pump hangs
	go pull(ch1)
	time.Sleep(1e9)
}
func pump(ch chan int) {
	for i := 0; ; i++ {
		//for i := 0; i < 7; i++ {
		ch <- i
		fmt.Println("pump i", i)
	}
}
func pull(ch chan int) {
	//time.Sleep(1e9)
	for i := 0; ; i++ {
		//for i := 0; i < 7; i++ {
		fmt.Println(<-ch)
	}
}
