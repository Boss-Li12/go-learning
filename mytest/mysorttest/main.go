package main

import (
	"fmt"
	"mySort"
)

type myArray []int

//implement the interface of Sort to sort the nodelist
func (s myArray) Len() int { return len(s) }
func (s myArray) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s myArray) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	var a myArray = []int{3, 4, 2, 5, 1, 0}
	mySort.Sort(a)
	fmt.Println(a)
}
