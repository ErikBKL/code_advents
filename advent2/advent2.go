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
			
		isValid, err := IsValidReport(lineSlice)
		if err != nil {
			return 0, err
		}
		if  isValid== true {
			ret++	
		}
	}
	
	return ret, nil
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

func GetCurrNext(lineSlice []string, idx int) (int, int, error) {
	curr, err := strconv.Atoi(lineSlice[idx - 1])
	if err != nil {
		return 0,0, err
	}
	next, err := strconv.Atoi(lineSlice[idx]) 
	if err != nil {
		return 0,0, err
	}

	return curr, next, nil
}

func IsPairValid(curr, next int) bool {
	if next - curr >= 1 && next - curr <= 3{
		return true		
	}
	return false
}

func IsValidReport(lineSlice []string )(bool, error) { 
	dampenerWasActivated := false

	for nextIdx := 1 ; nextIdx < len(lineSlice)  ; nextIdx++ {
		curr, next, err := GetCurrNext(lineSlice, nextIdx)
		if err != nil {
			return false, err
		}

		isValid := IsPairValid(curr, next)
		if isValid == false && dampenerWasActivated == false{
			dampenerWasActivated = true
			isValid, err = ActivateDampener(lineSlice, nextIdx)
			if err != nil {
				return false, err
			}
			
			if isValid == false {
				return false, nil
			} 	else if isValid == false && dampenerWasActivated == true {
				return false, nil
			}
		}
	}
	return true, nil
}

func ActivateDampener(lineSlice []string, nextIdx int) (bool, error){

	curr,err := strconv.Atoi(lineSlice[nextIdx - 1])
	if err != nil {
		return false, err
	}
	next,err := strconv.Atoi(lineSlice[nextIdx])
	if err != nil {
		return false, err
	}
	
	if nextIdx + 1 == len(lineSlice) {
		return true, nil
	}

	afterNext, err := strconv.Atoi(lineSlice[nextIdx + 1])
	if err != nil {
		return false, err
	}
	
	if IsPairValid(curr, afterNext) == true || 	IsPairValid(next, afterNext) == true {
		return true, nil
	}

	return false, nil
}