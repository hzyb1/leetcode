package main

import (
	"fmt"
	"leetcode/dfs"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func consumer(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	arr := dfs.CombinationSum2([]int{10,1,2,7,6,1,5}, 8)
	for i:=0;i<len(arr);i++ {
		fmt.Printf("%v\n", arr[i])
	}
}
func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}

