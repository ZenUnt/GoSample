package main

import (
	"fmt"
	"time"
)

func funcA(chA chan <- string) {
	time.Sleep(1 * time.Second)
	chA <- "funcA Finished"
}

func funcB(chB chan <- string) {
	time.Sleep(2 * time.Second)
	chB <- "funcB Finished"
}

func main() {
	chA := make(chan string)
	chB := make(chan string)
	defer close(chA)
	defer close(chB)
	finflagA := false
	finflagB := false
	go funcA(chA)
	go funcB(chB)
	for {
		select {
		case msg := <- chA:
			finflagA = true
			fmt.Println(msg)
		case msg := <- chB:
			finflagB = true
			fmt.Println(msg)
		}
		if finflagA && finflagB { // funcAとfuncBを待ち合わせる
			fmt.Println("funcA & funcB Finished!")
			break
		}
	}
}