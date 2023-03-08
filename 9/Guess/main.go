package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//建立消息通道
	ch1 := make(chan int)
	ch2 := make(chan int)
	//拉起随机数协程
	go generate(ch1, ch2)
	//拉起聪明的寻找数协程
	go smart(ch1)
	//拉起愚蠢的寻找数协程
	go idiot(ch2)
	time.Sleep(1e9)
}

func generate(ch1, ch2 chan int) {
	//设置随机种子
	fmt.Println("generate function start")
	rand.Seed(time.Now().UnixNano())
	//生成随机数
	rdn := rand.Intn(10000)
	fmt.Println("generate function generates", rdn)
	//设置退出标志
	flag1 := 0
	flag2 := 0
	for {
		//注意！千万不能写成<-ch1 == rdn 不然直接凉凉
		//先给聪明的函数判断
		if flag1 == 0 {
			v1, _ := <-ch1
			if v1 == rdn {
				fmt.Println("Oh Lord! Smart function has found it!")
				//猜中了返回0到隧道中
				ch1 <- 0
				flag1 = 1
			} else if v1 > rdn {
				//猜大了返回1到隧道中
				//fmt.Println("bigger")
				ch1 <- 1
			} else {
				//猜小了返回-1到隧道中
				//fmt.Println("smaller")
				ch1 <- -1
			}
		}

		//再给愚蠢的函数判断
		if flag2 == 0 {
			v2, _ := <-ch2
			if v2 == rdn {
				fmt.Println("Oh Lord! Idiot function has found it!")
				//猜中了返回0到隧道中
				ch2 <- 0
				flag2 = 1
			} else if v2 > rdn {
				//猜大了返回1到隧道中
				//fmt.Println("bigger")
				ch2 <- 1
			} else {
				//猜小了返回-1到隧道中
				//fmt.Println("smaller")
				ch2 <- -1
			}
		}
		if flag1 == 1 && flag2 == 1 {
			break
		}
	}
	fmt.Println("generate function end, both function have found it!")
}

func smart(ch1 chan int) {
	fmt.Println("smart function start")
	start := time.Now().UnixNano()//获取当前时间戳
	left := 0
	right := 10000
	mid := 0
	//利用二分找数
	for {
		mid = (left + right) / 2
		//fmt.Println("I am before ch1 <- mid")
		ch1 <- mid
		//fmt.Println("I am after ch1 <- mid")
		//注意！千万不能写成<-ch1 == x 不然直接凉凉
		v, _ := <-ch1
		if v == 0 {
			fmt.Println("Smart function: Oh Yeah! I have Found it! It is", mid)
			break
		} else if v == 1 {
			//fmt.Println("try smaller")
			right = mid
		} else if v == -1 {
			//fmt.Println("try bigger")
			left = mid
		}
	}
	fmt.Println("I am smart function, I spent", time.Now().UnixNano() - start, "Nanoseconds(纳秒) to found the answer")
}

func idiot(ch2 chan int) {
	fmt.Println("idiot function start")
	start := time.Now().UnixNano()//获取当前时间戳
	left := 0
	right := 10000
	cur := (left + right) / 2
	//利用二分找数
	for {
		//fmt.Println("I am before ch1 <- mid")
		ch2 <- cur
		//fmt.Println("I am after ch1 <- mid")
		//注意！千万不能写成<-ch1 == x 不然直接凉凉
		v, _ := <-ch2
		if v == 0 {
			fmt.Println("idiot fucntion: Oh Yeah! I have Found it! It is", cur)
			break
		} else if v == 1 {
			//fmt.Println("try smaller")
			cur -= 1
		} else if v == -1 {
			//fmt.Println("try bigger")
			cur += 1
		}
	}
	fmt.Println("I am idiot function, I spent", time.Now().UnixNano() - start, "Nanoseconds(纳秒) to found the answer")
}