package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var cgroups = "/sys/fs/cgroup"
var custom_cgroup = filepath.Join(cgroups, "teaho")

// go run main_test_cgroup.go run /bin/bash
// 执行20次 sleep 1000 & 会发现线程限制
func main() {
	switch os.Args[1] {
	case "run":
		run()
		// Removes the instantiated cgroup after container exit
		cgCleanup()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])
	cg()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}

func cg() {
	os.Mkdir(custom_cgroup, 0755)

	must(os.WriteFile(filepath.Join(custom_cgroup, "pids.max"), []byte("20"), 0644))
	must(os.WriteFile(filepath.Join(custom_cgroup, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0644))
}

func cgCleanup() error {
	alive, err := os.ReadFile(filepath.Join(custom_cgroup, "pids.current"))
	if err != nil { // or must(err).. but then it'll look weird..
		panic(err)
	}

	if alive[0] != uint8(48) {
		must(os.WriteFile(filepath.Join(custom_cgroup, "cgroup.kill"), []byte("1"), 0644))
	}
	must(os.Remove(custom_cgroup))

	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
