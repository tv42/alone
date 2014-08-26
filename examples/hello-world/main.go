package main

import (
	"fmt"
	"runtime"

	"github.com/tv42/alone/vm"
)

func main() {
	fmt.Println("Hello, world! I have", runtime.NumCPU(), "CPUs")
	vm.Exit()
}
