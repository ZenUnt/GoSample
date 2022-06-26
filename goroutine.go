package main

import (
	"fmt"
	"time"
)

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go funcA() // 並列処理でfuncAを実行
	for i :=0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(20 * time.Millisecond)
	}
}