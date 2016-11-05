package main

import "fmt"
import "math"

func main() {
	for i := 1; i <= 100; i++ {
		if math.Mod(float64(i), 15) == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if math.Mod(float64(i), 3) == 0 {
			fmt.Printf("Fizz\n")
		} else if math.Mod(float64(i), 5) == 0 {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
