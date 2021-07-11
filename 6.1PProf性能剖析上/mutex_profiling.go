package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func init() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
}

func main() {
	var m sync.Mutex
	var datas = make(map[int]struct{})
	for i := 0; i < 999; i++ {
		go func(i int) {
			m.Lock()
			defer m.Unlock()
			datas[i] = struct{}{}
		}(i)
	}
	_ = http.ListenAndServe("0.0.0.0:6061", nil)
}

// go build mutex_profiling.go
// go tool pprof http://localhost:6061/debug/pprof/mutex
