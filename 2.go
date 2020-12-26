package main

import (
	"fmt"
)

func main() {
	chs := make([] chan int, 123456)
	for i:=0; i<1; i++ {
		chs[i] = make(chan int)
		go perfectNum(123456)
	}

	for _, ch := range(chs) {
		<-ch
	}
}
func perfectNum(a int) {
	for i := 1; i <=a; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if sum == i {
			for j := 1; j < i; j++ {
				if i%j == 0 {
					fmt.Print(j,"+\t")
				}
			}
			fmt.Println("=",i)
		}
	}
}
//完美数算法来自VSDN