package main

import (
	"fmt"
	"math"
	"sort"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var N int = 4
	fmt.Scan(&N)
	var T []int

	diff := 10000000

	for i := 0; i < N; i++ {
		var pi int
		fmt.Scan(&pi)

		T = append(T, pi)

	}

	sort.IntSlice.Sort(T)

	for i := 1; i < len(T)-1; i++ {

		
			if T[i]-T[i+1] == 0 {
				diff = 0
				break
			}
			if diff > int(math.Abs(float64(T[i]-T[i+1]))) {
				diff = int(math.Abs(float64(T[i] - T[i+1])))
			}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(diff) // Write answer to stdout
}