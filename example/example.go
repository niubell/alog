package main

import (
	"fmt"
	"time"
	"ymtlog"
)

var logger = ymtlog.NewLogger("./", "example.log")

func main() {
	logger.SetLogLevel(ymtlog.INFO)
	fmt.Println("example start...")
	for {
		fmt.Println(" ----------batch start----------")
		logger.Info("%v\n", "info...")
		logger.Debug("%v\n", "debug...")
		logger.Warn("%v\n", "warn...")
		logger.Error("%v\n", "error...")
		time.Sleep(10 * time.Second)
		fmt.Println(" ----------batch end----------")
	}
}
