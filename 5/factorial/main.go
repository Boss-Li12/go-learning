package main

import (
	"fmt"
	"time"
)

func fac1(n int) int {
	if n == 1 {
		return n
	}
	return n * fac1(n-1)
}

func fac2(n int) (a int) {
	a = 1
	for i := 1; i <= n; i++ {
		a = a * i
	}
	return
}

func fac3(n int) int {
	arr := make([]int, n+1, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = i * arr[i-1]
	}
	return arr[n]
}


func main() {

	//var a int
	start := time.Now().UnixNano() //记录当前时间的纳秒数
	for i := 0; i < 1000000; i++ {
		fac1(10)
	}
	fmt.Println(fac1(10))
	fmt.Printf("Fac1's time was: %v\n", (time.Now().UnixNano() - start))

	start = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		fac2(10)
	}
	fmt.Println(fac2(10))
	fmt.Printf("Fac2's time was: %v\n", (time.Now().UnixNano() - start))

	start = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		fac3(10)
	}
	fmt.Println(fac3(10))
	fmt.Printf("Fac3's time was: %v\n", (time.Now().UnixNano() - start))


}

