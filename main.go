package main

import "fmt"

const usdToEurRate = 0.85
const usdToRubRate = 79.51
const eurToRubRate = usdToRubRate / usdToEurRate

func main() {

}

func getUserInput() (input string) {
	fmt.Scan(&input)
	return
}

func calculateRate(number float64, firstRate string, secondRate string) {

}
