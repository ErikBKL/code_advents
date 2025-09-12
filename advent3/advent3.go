package advent3

import (
	"fmt"
	"bufio"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func SumOfAllMul(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}

	idx := 0
	scanner := bufio.NewScanner(file)
	
	ret := 0

	for scanner.Scan() {
		line := scanner.Text()
		runeLine := []rune(line)
		for  idx = 0 ; idx < len(runeLine) ; idx++ {

			// find mul return idxM
			pos, isFound := FindMul(runeLine, idx )
			if isFound != true {
				break
			}
			idx = pos 

			if IsParenthesis(runeLine, idx+3) != true {
				idx += 2
				continue 
			}
			
			
			numLeft, idxAfterNum, err := ValidLeftNumber(runeLine, idx+4, idx+7)
			if err != nil {
				return 0, err
			} else if idxAfterNum == -1 {
				idx += 2
				continue
			}
				
				
			if IsComma(runeLine, idxAfterNum) != true {
				idx += 2
				continue
			}
			
			numRight, idxAfterNum, err := ValidRightNumber(runeLine, idxAfterNum+1, idxAfterNum + 4)
			if err != nil {
				return 0, err
			} else if idxAfterNum == -1 {
				idx += 2
				continue
			}
						
			ret += numLeft * numRight
			idx += idxAfterNum + 1000
			fmt.Printf("ret is %d\n", ret)
		}
	}

	return ret, nil
}

func ValidLeftNumber (runeLine []rune, start, end int ) (int, int, error) {
	i := start
	for ; i < end ; i++ {
		if runeLine[i] == ',' && i > start { //happy path if found comma and digit is at least 1 long
			break
		} else if unicode.IsDigit(runeLine[i]) != true { //sad path if it's here and not a digit it's also not a comma
			return 0, -1, nil
		} 
	}

	ret, err := strconv.Atoi(string(runeLine[start:i]))
	if err != nil {
		return 0, 0, err
	}

	return ret, i , nil
}


func ValidRightNumber (runeLine []rune, start, end int ) (int, int, error) {
	i := start
	for ; i < end ; i++ {
		if runeLine[i] == ')' && i > start { //happy path if found ) and digit is at least 1 long
			break
		} else if unicode.IsDigit(runeLine[i]) != true { //sad path if it's here and not a digit it's also not a )
			return 0, -1, nil
		} 
	}

	ret, err := strconv.Atoi(string(runeLine[start:i]))
	if err != nil {
		return 0, 0, err
	}

	return ret, i , nil
}


func FindTargetInSrc(target []rune, src []rune, startIdx int) (int, bool) {
	// assuming no whitespace

	for i := startIdx; i < len(src)-(len(target)-1); i++ {
		toInspect := src[i : i+len(target)]

		if slices.Equal(toInspect, target) {
			return i, true
		}
	}

	return 0, false
}

func FindMul(src []rune, idxStart int) (int, bool) {
	ret, isFound := FindTargetInSrc([]rune("mul"), src, idxStart)
	return ret, isFound
}

func IsParenthesis(runeLine []rune, idx int) bool {
	if runeLine[idx] != '(' {
		return false
	}

	return true
}

func IsComma(runeLine []rune, idx int) bool {
	if runeLine[idx] != ',' {
		return false
	}

	return true
}