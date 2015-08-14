package main

import (
	"fmt"
	"github.com/bigpyer/ymtlog"
	"time"
)

func main() {
	fmt.Println("example start...")
	for {
		fmt.Println(" ----------batch start----------")
		ymtlog.Info("%v\n", "info...")
		ymtlog.Debug("%v\n", "debug...")
		ymtlog.Warn("%v\n", "warn...")
		ymtlog.Error("%v\n", "error...")
		time.Sleep(10 * time.Second)
		fmt.Println(" ----------batch end----------")
	}
}
