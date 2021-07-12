package main

import (
	"fmt"
)


func main() {

	num := 10
	sign := make(chan struct{},num)
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}

	for j:=0;j<num;j++ {
		<-sign
	}
}
