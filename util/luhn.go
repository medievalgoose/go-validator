package util

import (
	"regexp"
	"strconv"
)

func ValidateCreditCard(number string) (bool, string) {
	var numberArray []int
	var numberArrayProcessed []int
	indexTracker := 0
	totalNum := 0

	for _, num := range number {
		numConv, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}

		numberArray = append(numberArray, numConv)
	}

	for i := len(numberArray) - 1; i >= 0; i-- {
		currentNumber := 0
		indexTracker++
		if indexTracker%2 == 0 {
			currentNumber = numberArray[i] * 2
			if currentNumber > 9 {
				numberString := strconv.Itoa(currentNumber)
				var tempArray []int
				for _, num := range numberString {
					convNumber, err := strconv.Atoi(string(num))
					if err != nil {
						panic(err)
					}
					tempArray = append(tempArray, convNumber)
				}
				currentNumber = tempArray[0] + tempArray[1]
			}
		} else {
			currentNumber = numberArray[i]
		}
		numberArrayProcessed = append(numberArrayProcessed, currentNumber)
	}

	for _, num := range numberArrayProcessed {
		totalNum += num
	}

	if totalNum%10 == 0 {
		return true, classifyCreditCard(number)
	} else {
		return false, classifyCreditCard(number)
	}
}

func classifyCreditCard(number string) string {
	visaCardExp := "\\b4\\d{15}\\b"

	// Visa
	visaMatch, _ := regexp.MatchString(visaCardExp, number)
	if visaMatch {
		return "Visa"
	}

	// Mastercard
	masterCardExp := "\\b5[1-5]\\d{14}\\b"
	masterMatchv1, _ := regexp.MatchString(masterCardExp, number)
	if masterMatchv1 {
		return "Mastercard"
	}

	masterCardExpV2 := "\\b2[2-7]\\d{14}\\b"
	masterMatchv2, _ := regexp.MatchString(masterCardExpV2, number)
	if masterMatchv2 {
		return "Mastercard"
	}

	// American Express
	amexCardExp := "\\b3[47]\\d{13}\\b"
	amexMatch, _ := regexp.MatchString(amexCardExp, number)
	if amexMatch {
		return "American Express"
	}

	return "-"
}
