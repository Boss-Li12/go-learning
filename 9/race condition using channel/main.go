package main

import (
	"fmt"
	"sync"
)

var x = 0 //全局变量

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1 //访问全局变量
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	//必须带大小为1的缓冲区(一个时间只有一个人），因为所有的协程都是以写入开始
	ch := make(chan bool, 1000)
	for i := 0; i < 100000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x) //结果不确定
}
