package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		cmdLine := strings.TrimSpace(cmd)
		args := strings.SplitAfterN(cmdLine, " ", 2)
		cmd = strings.TrimSpace(args[0])

		switch cmd {
		case "":
		case "echo":
			if len(args) > 1 {
				fmt.Println(args[1])
			}
		case "type":
			if len(args) > 1 {
				printType(args[1])
			}
		case "exit":
			os.Exit(0)
		default:
			runExecutable(args)
		}
	}
}

func runExecutable(args []string) {
	cmd := strings.TrimSpace(args[0])
	var ex *exec.Cmd
	if len(args) < 2 {
		ex = exec.Command(cmd)
	} else {
		ex = exec.Command(cmd, args[1])
	}
	ex.Stdin = os.Stdin
	ex.Stdout = os.Stdout
	err := ex.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func printType(cmd string) {
	builtin_cmds := []string{
		"echo", "exit", "type",
	}

	for _, v := range builtin_cmds {
		if cmd == v {
			fmt.Printf("%s is a shell builtin\n", cmd)
			return
		}
	}

	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)

		for _, file := range files {
			if cmd == file.Name() {
				if !file.IsDir() {
					fmt.Printf("%s is %s/%s\n", cmd, dir, file.Name())
					return
				}
			}
		}
	}

	fmt.Printf("%s not found\n", cmd)
}
