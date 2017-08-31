package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func run(name string, args ...string) {
	fmt.Println("Running %v", name)
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	switch os.Args[1] {
	case "run":
		run("/proc/self/exe", append([]string{"child"}, "/bin/bash")...)
	case "child":
		err := syscall.Chroot("files")
		if err != nil {
			fmt.Println(err)
		}
		
		err = os.Chdir("/")
		if err != nil {
			fmt.Println(err)
		}
		
		err = syscall.Mount("proc", "proc", "proc", 0, "")
		if err != nil {
			fmt.Println(err)
		}
		
		run(os.Args[2])
	default:
		panic("What???")
	}
}
