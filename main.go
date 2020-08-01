package main

import (
	"log"
	"os"
	"time"
)

func init() {
	file := "./" + "log" + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[oss-helper]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	go func() {
		for i := 0; i < 10000; i++ {
			time.Sleep(time.Second)
		}
	}()

	println("启动了~")

	select {}
}
