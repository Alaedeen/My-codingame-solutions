package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	MESSAGE := scanner.Text()
	BinaryMessage := ""
	for _, m := range MESSAGE {
		binary := fmt.Sprintf("%b", m)
		if len(binary) < 7 {
			for i := 0; i < 7-len(binary); i++ {
				binary = "0" + binary
			}
		}
		BinaryMessage += binary
	}
	UnaryMessage := ""
	var current rune
	for i, b := range BinaryMessage {
		if current == b {
			UnaryMessage += "0"
		} else {
			if i != 0 {
				UnaryMessage += " "
			}
			if b == '1' {
				UnaryMessage += "0 0"
			} else {
				UnaryMessage += "00 0"
			}
		}
		current = b
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(UnaryMessage) // Write answer to stdout
}
