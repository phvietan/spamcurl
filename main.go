package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	path, err := exec.LookPath("curl")
	if err != nil {
		panic("Command 'curl' not found")
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received SIGINT: stopping process")
		os.Exit(0)
	}()

	args := os.Args
	for _, arg := range args {
		if arg == "-h" {
			fmt.Printf("Usage: %s [options...] <url>\nWill spam 'curl' command until receive SIGINT\nBelow is curl help\n==============================================\n", args[0])
			args[0] = "curl"
			syscall.Exec(path, args, os.Environ())
			break
		}
	}
	args[0] = "curl"
	for {
		pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if pid == 0 {
			syscall.Exec(path, args, os.Environ())
		} else {
			proc, err := os.FindProcess(int(pid))
			if err != nil {
				panic("Something is wrong in spawning processes")
			}
			_, err = proc.Wait()
			if err != nil {
				panic(err.Error())
			}
		}
	}

}
