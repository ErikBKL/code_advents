package advent2

import (

	"bufio"
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
	if next - curr >= 1 && next - curr <= 3{
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

			isValid := TryMakeValidReport(numSlice, idx)

			if isValid != true {
				return false, nil
			}
		}
	}
	return true, nil
}

func TryMakeValidReport(numSlice []int, idx int) bool {

	numSlice = append(numSlice[:idx], numSlice[idx+1:]...)
		
	return IsPureValidReport(numSlice, idx)
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