package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	operations := map[string]func([]float64) float64{
		"AVG": OperationAVG,
		"SUM": OperationSUM,
		"MED": OperationMED,
	}

	operation := getInputOperation(operations)
	numbers := getInputNumber()

	result := operations[operation](numbers)

	fmt.Printf("The Result of your %v operation = %v\n", operation, result)
}

func getInputOperation(operations map[string]func([]float64) float64) string {
	var operation string
	for {
		fmt.Print("Select the operation to be calculated (AVG, SUM, MED):")
		fmt.Scan(&operation)
		if _, exists := operations[operation]; exists {
			break
		}
		fmt.Println("Invalid operation. Please enter AVG, SUM, MED")
	}
	return operation
}

func getInputNumber() []float64 {
	var listNumbers string
	var sliceNumbers []float64

	fmt.Print("Input numbers for operation in commas:")
	fmt.Scan(&listNumbers)

	stringNumbers := strings.Split(listNumbers, ",")

	for _, string := range stringNumbers {
		num, err := strconv.ParseFloat(string, 64)

		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			continue
		}

		sliceNumbers = append(sliceNumbers, num)
	}

	return sliceNumbers
}

func OperationAVG(numbers []float64) float64 {
	var sum float64
	lengthSlice := len(numbers)

	if lengthSlice == 0 {
		return 0
	}

	for _, num := range numbers {
		sum += num
	}
	AVG := sum / float64(lengthSlice)
	return AVG
}

func OperationSUM(numbers []float64) float64 {
	var SUM float64
	lengthSlice := len(numbers)

	if lengthSlice == 0 {
		return 0
	}

	for _, num := range numbers {
		SUM += num
	}
	return SUM
}

func OperationMED(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Float64s(numbers)
	middleNumber := len(numbers) / 2

	if len(numbers)%2 != 0 {
		return numbers[middleNumber]
	}
	return (numbers[middleNumber-1] + numbers[middleNumber]) / 2
}
