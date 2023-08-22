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
		if arg == "-h" || arg == "--help" {
			fmt.Printf("Usage: %s [options...] <url>\nWill spam 'curl' command until receive SIGINT\nBelow is curl help\n==============================================\n", args[0])
			args[0] = "curl"
			cmd := exec.Command(path, args...)
			stdout, err := cmd.Output()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(stdout))
			os.Exit(0)
		}
	}
	args[0] = "curl"
	for {
		cmd := exec.Command(path, args...)
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
	}
}
