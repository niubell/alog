package main

import (
	"fmt"
	"github.com/bigpyer/alog"
	"time"
)

var logger = alog.NewLogger("./", "example.log")

func main() {
	logger.SetLogLevel(alog.INFO)
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
