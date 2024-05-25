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
		if cmd == "exit 0" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", cmd)
	}
}
