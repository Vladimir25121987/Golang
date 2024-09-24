package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("КАЛЬКУЛЯТОР")
	results := []float64{}
	for {
	number := scanNumber()
	results = append(results, number)
	}
	operation := scanOperation()
	switch operation {
	case "AVG":
		fmt. Printf("Результат вычислений: %.f", calculateAvg(results))
	case "SUM":
		fmt. Printf("Результат вычислений: %.f", calculateSum(results))
	case "MED":
		fmt. Printf("Результат вычислений: %.f", calculateMed(results))
	default:
		fmt. Println("Нет операции. Попробуйте заново")
	 }	
   }


func scanOperation() string {
	var operation string
	fmt.Println("Введите нужную операцию(AVG, SUM, MED): ")
	fmt.Scan(&operation)
	return operation
}

func scanNumber() float64 {
	var number float64
	fmt.Println("Введите числа через запятую: ")
	fmt.Scan(&number)
	return number
}

func calculateAvg(results []float64) float64 {
	var avg = 0.0
	var number float64
	for _, number = range results {
		avg = (avg + number) / float64(len(results))
	}
	return avg
}

func calculateSum(results []float64) float64 {
	var sum = 0.0
	var namber float64
	for _, namber = range results {
		sum += namber
	}
	return sum
}

func calculateMed(results []float64) float64 {
	resultsCopy := results
	sort.Float64s(resultsCopy)
	if  len(resultsCopy) % 2 != 0 {
		return resultsCopy[(len(resultsCopy) + 1) / 2]
	} else {
        return resultsCopy[(len(resultsCopy) + 1) / 2 + len(resultsCopy)] / 2 
	}
}