package main
//通过使用chanel 实现主协程等待其他协程执行结束。
import (
	"fmt"
)

func main()  {
	//sign := make(chan string,1)
	done := make(chan bool,1)
	//sign <- "nihao"

	//signal.Notify(sign,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		//sig := <- sign
		fmt.Println()
		fmt.Println("正在执行go协程")
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
