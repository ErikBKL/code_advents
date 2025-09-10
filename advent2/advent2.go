package advent2

import (

	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

func AmountSafeReports(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := 0
	for scanner.Scan() {
		// 	put line in slice
		lineSlice := strings.Fields(scanner.Text())
		lineSlice, err := MakeAscendingOrder(lineSlice)
		if err != nil {
			return 0, err
		}

		numSlice, err := StrSliceToInts(lineSlice)
		if err != nil {
			return 0, err
		}
		isValid, err := IsValidReport(numSlice)
		if err != nil {
			return 0, err
		}
		if  isValid== true {
			ret++	
		}
	}
	
	return ret, nil
}

func StrSliceToInts(lineSlice []string) ([]int, error) {
	numSlice := make([]int, 0, len(lineSlice))

	for i,_ := range lineSlice {
		toAppend, err := GetIntAtIdx(lineSlice, i)
		if err != nil {
			return nil, err
		}

		numSlice = append(numSlice, toAppend )
	}

	return numSlice, nil
}

func MakeAscendingOrder(lineSlice []string) ([]string, error) {
	curr, err := strconv.Atoi(lineSlice[0])
	if err != nil {
		return nil, err
	}
	next, err := strconv.Atoi(lineSlice[1])
	if err != nil {
		return nil, err
	}

	if curr > next {
		slices.Reverse(lineSlice)
	}

	return lineSlice, nil
}

func GetIntAtIdx(lineSlice []string, idx int) (int, error) {
	ret, err := strconv.Atoi(lineSlice[idx])
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func IsPairValid(curr, next int) bool {
	diff := next - curr
	if diff < 0 {
		diff = -1*diff
	}

	if diff >= 1 && diff <= 3{
		return true		
	}
	return false
}

func IsValidReport(numSlice []int )(bool, error) { 

	for idx := 0 ; idx < len(numSlice) - 1; idx++ {
		curr := numSlice[idx]
		next := numSlice[idx + 1]
		
		isValid := IsPairValid(curr, next)
		if isValid == false {
			// if becomes valid by popping curr or next
			if TryMakeValidReport(numSlice, idx) || TryMakeValidReport(numSlice, idx+1) {
				return true ,nil
			} else {
				return false, nil
			}
		}
	}
	return true, nil
}

func TryMakeValidReport(numSlice []int, idx int) bool {
	
	cpy := make([]int, len(numSlice))
	
	copy(cpy, numSlice)
	
	cpy = append(cpy[:idx], cpy[idx+1:]...)

	fmt.Printf("cpy: %v\n", cpy)
	isValid := IsPureValidReport(cpy, idx-1)

	return isValid
}

func IsPureValidReport(numSlice []int, idx int, ) bool {
	i := idx - 1
	if i < 0 {
		i = 0
	}
	for ; i < len(numSlice) - 1 ; i++ {
		// foreach idx, tryremove from slice and check if report is pure valid
		isValid := IsPairValid(numSlice[i], numSlice[i+1])
		if isValid == false {
			return false
		}
	}

	return true
}