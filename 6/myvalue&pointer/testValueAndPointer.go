package main

import (
	"fmt"
)

type Alien struct {
	color string
	sex string
}

//方法接受者是value
func (a Alien) changeColor(newColor string) {
	a.color = newColor
}

//方法接受者是address
func (a *Alien) changeSex(newSex string) {
	a.sex = newSex
}

func main() {
	a := Alien{
		color: "white",
		sex: "unknown",
	}

	fmt.Printf("Alien color before change: %s\n", a.color)
	a.changeColor("Black")
	fmt.Printf("Alien color after change: %s\n", a.color)

	fmt.Printf("Alien sex before change: %s\n", a.sex)
	a.changeSex("doublesex")
	fmt.Printf("Alien sex after change: %s\n", a.sex)
}
