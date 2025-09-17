package advent7

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)


func SumValidEquations (pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := 0
	for scanner.Scan() {
		line := scanner.Text()
		
		target, operands, err := SplitTargetOperands(line)
		if err != nil {
			return 0, err
		}

		if RecIsValidEquation(target, operands[0], operands[1 : ]) == true {
			ret += target
		}
	}

	return ret, nil
}

func RecIsValidEquation(target int, cumulativeRes int, operands []int) bool {
	operators := []string{"+", "*", "||"}

	if cumulativeRes == target && len(operands) == 0 { 
		return true 
	} else if len(operands) == 0  && cumulativeRes != target { 
		return false
	}

	cumulative := 0
	for _, op := range operators {
		
		switch op {
		case "+":
			cumulative = cumulativeRes + operands[0]
		case "*":
			cumulative = cumulativeRes * operands[0]
		default:
			lenToAppend := len([]rune(strconv.Itoa(operands[0])))
			cumulative = cumulativeRes * int(math.Pow10(lenToAppend)) + operands[0]
		}

		isValid := RecIsValidEquation(target, cumulative, operands[1 : ])
		if isValid == true {
			return true
		}
	} 
	return false
}

func SplitTargetOperands(input string) (int, []int, error) {
	expression := strings.Split(input, ":")	
	target := expression[ : 1]
	
	intTarget, err := strconv.Atoi(strings.TrimSpace(target[0]))
	if err != nil {
		return 0, nil, err
	}

	operands := expression[1 : ]
	sliceOperands := strings.Split(strings.TrimSpace(operands[0]), " ")
	intOperands := []int{}

	for _,v := range sliceOperands {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, nil, err
		}
		
		intOperands = append(intOperands, num)
	}

	return intTarget, intOperands, nil
}

