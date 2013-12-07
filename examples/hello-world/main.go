package main

import (
	"fmt"
	"github.com/tv42/alone/vm"
	"runtime"
)

func main() {
	fmt.Println("Hello, world! I have", runtime.NumCPU(), "CPUs")
	vm.Exit()
}
