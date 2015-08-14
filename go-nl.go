package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		color := ""
		if count%2 == 0 {
			color = fmt.Sprintf("\x1b[48;5;233m")
		}
		fmt.Printf("%s%d - %s\x1b[0m\n", color, count, line)
		count += 1
	}

}
