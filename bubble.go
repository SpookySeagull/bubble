// This is my first GitHub commithion
// It is just a bubble sort writed from memory from 2014 when I was in college
// In this itherathion I added a fancy visualization of sorting process
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This is ANSI escape code for clearing line in terminal
// I'm now sure if it will work on Windows
const ClearLine = "\033[2K"

func bubble(array [10]int) [10]int {
	var hold int
	for i := 0; i < len(array); i++ {
		for x := 0; x < (len(array) - 1); x++ {
			fmt.Print(array)
			time.Sleep(time.Second / 10)
			if array[x] > array[x+1] {
				hold = array[x]
				array[x] = array[x+1]
				array[x+1] = hold
			}
			fmt.Print(ClearLine)
			fmt.Printf("\r")
		}
	}
	return array
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10)
		// PC nowadays so fast so I need to add Sleep
		// Just to avoid duplicate numbers in a row
		// Feel free to add this to cursed programming compilation
		time.Sleep(time.Second / 10)
	}
	fmt.Println(array)
	fmt.Println(bubble(array))
}
