package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //goroutine执行结束后将信号量减1
}

func main() {

	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1) //增加信号量
		go process(i, &wg)
	}
	//等所有协程做完
	wg.Wait() //主goroutine阻塞调用该方法的协程，直到等待信号量为0
	fmt.Println("All go routines finished executing time:", (time.Now().UnixNano() - start) / 1e9)
}
