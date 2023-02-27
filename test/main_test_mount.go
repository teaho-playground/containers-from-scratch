package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	//must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("/proc", 0))
	//must(syscall.Unmount("/mytemp", 0))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
