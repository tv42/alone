package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/tv42/alone/vm"
)

func showNets() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("error listing network interfaces: ", err)
		return
	}

	for _, iface := range ifaces {
		fmt.Printf("%s %v %v\n", iface.Name, iface.Flags, iface.HardwareAddr)
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("  %s %s\n", addr.Network(), addr.String())
		}
	}
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, world!\n")
}

func main() {
	showNets()

	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("error serving HTTP: %v", err)
		return
	}

	vm.Exit()
}
