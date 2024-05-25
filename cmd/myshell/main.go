package main

import (
	"bufio"
	"fmt"
	"os"
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

		cmd = strings.TrimSpace(cmd)
		args := strings.SplitAfterN(cmd, " ", 2)

		switch args[0] {
		case "":
			continue
		case "echo ":
			fmt.Println(args[1])
		case "type ":
			isValidCmd(args[1])
		case "exit ":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func isValidCmd(cmd string) {
	valid_cmds := []string{
		"echo", "exit", "type",
	}

	for _, v := range valid_cmds {
		if cmd == v {
			fmt.Printf("%s is a shell builtin\n", cmd)
			return
		}
	}

	fmt.Printf("%s not found\n", cmd)
}
