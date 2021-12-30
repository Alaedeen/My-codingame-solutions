package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")

	if n == 0 {
		fmt.Println(0)
	} else {
		var temperatures []int
		for i := 0; i < n; i++ {
			// t: a temperature expressed as an integer ranging from -273 to 5526
			t, _ := strconv.ParseInt(inputs[i], 10, 32)
			temperatures = append(temperatures, int(t))
		}

		sort.IntSlice.Sort(temperatures)

		for len(temperatures) > 1 {
			first := math.Abs(float64(temperatures[(len(temperatures)/2)-1]))
			second := math.Abs(float64(temperatures[(len(temperatures) / 2)]))

			if first < second {
				temperatures = temperatures[0 : len(temperatures)/2]
			} else {
				temperatures = temperatures[(len(temperatures) / 2):]
			}
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		fmt.Println(temperatures[0]) // Write answer to stdout
	}
}
