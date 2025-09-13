package advent3

import (
	"os"
	"regexp"
	"strconv"
)

var patternRX = `mul\(\d{1,3},\d{1,3}\)`


func SumOfAllMul(pathToFile string) (int, error) {
	bytesData, err := os.ReadFile(pathToFile)
	if err != nil {
		return 0, err
	}

	stringData := string(bytesData)
	rx := regexp.MustCompile(patternRX)

	matches := rx.FindAllString(stringData, -1)
	if matches == nil {
		return 0, nil
	}

	ret := 0
	res := 0
	for i := 0 ; i < len(matches) ; i++ {
		res , err = Calculate(matches[i])
		if err != nil {
			return 0, err
		}

		ret += res
	}

	return ret, nil
}

func Calculate (expr string) (int,error) {
	runeExpr := []rune(expr)

	i := 4
	for ; i < 7 ; i++ {
		if runeExpr[i] == ',' {
			break
		}
	}

	numLeft, err := strconv.Atoi(string(runeExpr[4:i]))
	if err != nil {
		return 0, nil
	}

	i++
	firstDigit := i
	for ; runeExpr[i] != ')' ; i++ {
		
	}

	numRight, err := strconv.Atoi(string(runeExpr[firstDigit : i]))
	if err != nil {
		return 0, err
	}

	return numRight * numLeft, nil
}