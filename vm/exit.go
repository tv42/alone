package vm

import "syscall"
import "os"
import "log"

// Exit causes the virtual machine to exit. Do not run outside an
// alone app. Exit never returns to the caller.
func Exit() {
	var err error
	err = syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
	if err != nil {
		log.Printf("power off failed: %v", err)
	}
	err = syscall.Reboot(syscall.LINUX_REBOOT_CMD_HALT)
	if err != nil {
		log.Printf("halt failed: %v", err)
	}
	os.Exit(1)
	select {}
}
