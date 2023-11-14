package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath string = "./input.txt"

	// Read the entire file content
	allElfCalories, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// total calories per elf split by two consecutive line breaks
	elfCaloriesArr := strings.Split(string(allElfCalories), "\n\n")

	var mostCalories int = 0 // total calories of elf with most calories
	var mostCalories2 int = 0 // total calories of elf with second-most calories
	var mostCalories3 int = 0 // total calories of elf with third-most calories

	for _, elfCalories := range elfCaloriesArr {
		calorieCount := 0
		// total elf calories per item, split by line break
		elfCaloriesArray := strings.Split(elfCalories, "\n")

		for _, itemCalories := range elfCaloriesArray {
			itemCalories, err := strconv.Atoi(itemCalories)
			if err == nil {
				calorieCount += itemCalories
			}
		}

		if calorieCount > mostCalories {
			mostCalories = calorieCount
		} else if calorieCount > mostCalories2{
			mostCalories2 = calorieCount
		} else if calorieCount > mostCalories3 {
			mostCalories3 = calorieCount
		}
	}

	// Total calories of top three elves with most calories
	fmt.Println("total : ", mostCalories + mostCalories2 + mostCalories3)
}
