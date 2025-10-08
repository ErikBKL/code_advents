package advent11

import (
	"os"
	"strconv"
	"strings"
	"fmt"
	// "slices"
)

func CountStones(pathToFile string) (int, error) {
	bytesData, err := os.ReadFile(pathToFile)
	if err != nil {
		return 0, err
	}

	stringData := string(bytesData)
	input := strings.Split(stringData, " ")
	
	for blinks := 0 ; blinks < 25 ; blinks++ {

		nextInput := []string{}

		for _,v := range input {
			stoneNumber, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
	
		
			if stoneNumber == 0 {
				nextInput = append(nextInput, "1")
			} else if len(v) % 2 == 0 {
				left,right, err := SplitNumber(v)
				if err != nil {
					return 0, err
				}
				nextInput = append(nextInput, left)
				nextInput = append(nextInput, right)
			} else {
				toInsert := strconv.Itoa(stoneNumber * 2024)
				nextInput = append(nextInput, toInsert)
			}
			
		}
		input = nextInput
		fmt.Printf("%+v", input)
	}
	return len(input), nil
}

func SplitNumber(original string)(string, string, error) {

	length := len(original)
	left := original[0 : length/2]
	right := original[length/2 : length]

	numRight, err := strconv.Atoi(right)
	if err != nil {
		return "","", err
	}

	right = strconv.Itoa(numRight)

	return left, right, nil
}