package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				os.Exit(1)
			}
			fmt.Println(err)
		}

		fmt.Printf("%s: command not found\n", strings.TrimSpace(cmd))
	}
}
