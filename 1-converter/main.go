package main

import "fmt"

func main() {

	const usdToEur = 0.92
	const usdToRub = 98.5
	eurToRub := usdToRub / usdToEur
	fmt.Print(eurToRub)
}