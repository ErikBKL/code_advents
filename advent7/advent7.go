package advent7

import (
	"bufio"
	"os"
	"strings"
	"strconv"
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

		if RecIsValidEquation(target, 0, operands) == true {
			ret += target
		}
	}

	return ret, nil
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