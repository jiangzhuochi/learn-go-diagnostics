package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}

// go build main.go
// ./main

// 通过浏览器访问
// http://127.0.0.1:6060/debug/pprof/

// 通过交互式终端使用
// PowerShell 下运行命令

// （1）CPU Profiling
// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

// 交互式终端输入 help 查看帮助

// 下同
// （2）Heap
// （3）Goroutine
