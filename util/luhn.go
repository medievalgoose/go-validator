package util

import "strconv"

func ValidateCreditCard(number string) bool {
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
		return true
	} else {
		return false
	}
}
