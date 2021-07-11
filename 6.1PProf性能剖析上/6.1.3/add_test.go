package main

import (
	"testing"

	"learn-pprof-part6.1.3/add"
)

func TestAdd(t *testing.T) {
	_ = add.Add("go-programming-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add.Add("go-programming-tour-book")
	}
}

// PowerShell 下 . 符号要加双引号，cmd 下不用，试了 macOS 的 zsh，也不用
// . 是个运算符，而我们只需要字面量，让 go 去解释
// go test -bench="." -cpuprofile="cpu.profile"
// 如果写成 go test -bench=.
// PowerShell 认为此命令是 go test -bench= . 传给 go test
// go test 解释为 -bench 没有任何参数，因此不运行基准，只运行测试

// 也可以写成 go test -bench. -cpuprofile="cpu.profile"
// go test -bench . -cpuprofile="cpu.profile"
// go test "-bench=." -cpuprofile="cpu.profile"
// go test "-bench=." "-cpuprofile=cpu.profile"

// 但是 go test "-bench=. -cpuprofile=cpu.profile" 是不可以的
// 因为 go test 收到的是一个参数而不是两个，解释为 -bench 的参数为 ". -cpuprofile=cpu.profile"
// 没有基准函数符合 ". -cpuprofile=cpu.profile"，因此不运行基准，只运行测试

// 生成 cpu.profile 后，运行
// go tool pprof "cpu.profile"
