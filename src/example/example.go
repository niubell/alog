package main

import (
	"fmt"
	"logger"
	"time"
)

func main() {
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
