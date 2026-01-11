package main

import (
	"fmt"
	"lasagna"
)

func main() {
	fmt.Println("Lasagna Cooking Calculator")
	fmt.Println("==========================")
	
	layers := 3
	minutesInOven := 20
	
	result := lasagna.ElapsedTime(layers, minutesInOven)
	fmt.Printf("Elapsed time for %d layers with %d minutes in oven: %d minutes\n", layers, minutesInOven, result)
	
	fmt.Printf("Expected oven time: %d minutes\n", lasagna.OvenTime)
	fmt.Printf("Preparation time for %d layers: %d minutes\n", layers, lasagna.PreparationTime(layers))
	fmt.Printf("Remaining oven time: %d minutes\n", lasagna.RemainingOvenTime(minutesInOven))
}
