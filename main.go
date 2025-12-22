package main

import (
	"errors"
	"fmt"
)

const (
	RUB = "RUB"
	EUR = "EUR"
	USD = "USD"
)

const (
	usdToEurRate float64 = 0.85
	usdToRubRate float64 = 79.51
	eurToRubRate float64 = usdToRubRate / usdToEurRate
	eurToUsdRate float64 = 1.0 / usdToEurRate
	rubToUsdRate float64 = 1.0 / usdToRubRate
	rubToEurRate float64 = rubToUsdRate / eurToUsdRate
)

func main() {
	for {
		firstCurrency, err := getUserFirstCurrency()
		if errorCheck(err) {
			continue
		}

		quantityCurrency, err := getUserQuantityCurrency()
		if errorCheck(err) {
			continue
		}

		lastCurrency, err := getUserLastCurrency(firstCurrency)
		if errorCheck(err) {
			continue
		}

		result := calculateRate(quantityCurrency, firstCurrency, lastCurrency)
		fmt.Printf("%.0f %s = %.2f %s\n", quantityCurrency, firstCurrency, result, lastCurrency)
		if !checkRepeat() {
			break
		}
	}
}

func checkRepeat() bool {
	fmt.Print("Продолжить конвертацию? (Y/n): ")
	var checkRepeat string
	fmt.Scan(&checkRepeat)
	if checkRepeat == "y" || checkRepeat == "Y" {
		return true
	}
	return false
}

func errorCheck(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func getUserQuantityCurrency() (float64, error) {
	fmt.Print("Введите количество валюты: ")
	var result float64
	_, err := fmt.Scan(&result)
	if err != nil {
		return 0, errors.New("ERROR! Введите число")
	}
	return result, nil
}

func getUserFirstCurrency() (input string, errInput error) {
	fmt.Print("Введите исходную валюту для конвертации (USD/EUR/RUB): ")
	_, err := fmt.Scan(&input)
	if err != nil {
		return input, err
	}
	if input != RUB && input != EUR && input != USD {
		return input, errors.New("ERROR! Не доступная валюта")
	}
	return input, nil
}

func getUserLastCurrency(currency string) (input string, errInput error) {
	fmt.Print("Введите целевую валюту для конвертации, доступная валюта: ")
	switch currency {
	case RUB:
		fmt.Print("(USD EUR): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			return input, err
		}
		if input != EUR && input != USD {
			return input, errors.New("ERROR! Не доступная валюта")
		}
	case EUR:
		fmt.Print("(RUB USD): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			return input, err
		}
		if input != USD && input != RUB {
			return input, errors.New("ERROR! Не доступная валюта")
		}
	case USD:
		fmt.Print("(RUB EUR): ")
		_, err := fmt.Scan(&input)
		if err != nil {
			return input, err
		}
		if input != RUB && input != EUR {
			return input, errors.New("ERROR! Не доступная валюта")
		}
	}
	return input, nil
}

func calculateRate(number float64, firstCurrency string, secondCurrency string) float64 {
	switch {
	case firstCurrency == USD && secondCurrency == RUB:
		return number * usdToRubRate
	case firstCurrency == USD && secondCurrency == EUR:
		return number * usdToEurRate
	case firstCurrency == EUR && secondCurrency == RUB:
		return number * eurToRubRate
	case firstCurrency == EUR && secondCurrency == USD:
		return number * eurToUsdRate
	case firstCurrency == RUB && secondCurrency == EUR:
		return number * rubToEurRate
	case firstCurrency == RUB && secondCurrency == USD:
		return number * rubToUsdRate
	}
	return 0
}
