package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。") // 2021/11/10 10:51:41.674901 F:/code/GoCodes/awesomeProject/golangBasic/log/log.go:17: 这是一条很普通的日志。
	log.SetPrefix("[小王子]") // 设置前缀，方便检索
	log.Println("这是一条很普通的日志。") // [小王子]2021/11/10 10:51:41.689831 F:/code/GoCodes/awesomeProject/golangBasic/log/log.go:19: 这是一条很普通的日志。
}