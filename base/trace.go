package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	//启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	//go run trace.go
	//go tool trace .\trace.go
	trace.Stop()
}
