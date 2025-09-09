package advent1

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)




func PrepareData(pathToFile string) ([]int, []int, error) {
	file, err := os.Open(pathToFile)
	defer file.Close()
	 
	if err != nil {
		return nil, nil, err
	}
	
	// foreach line in file
	scanner := bufio.NewScanner(file)
	var numbers1 []int
	var numbers2 []int

	for scanner.Scan() {
		// split by space/tab delimiter and store to num1 num2
		line := scanner.Text()

		lineSlice := strings.Fields(line)
		// append num1 to slice1
		num1, err := strconv.Atoi(lineSlice[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			return nil, nil, err
		}

		numbers1 = append(numbers1, num1 )
		// append num2 to slice2
		numbers2 = append(numbers2, num2)
	}

	return numbers1, numbers2, nil
}



func AppendNumber(slice []int, toAppend []rune) ([]int, []rune, error) {
	numToInsert, err := strconv.Atoi(string(toAppend))
	if err != nil {
		return slice, toAppend, err
	}

	slice = append(slice, numToInsert)
	toAppend = toAppend[:0]	

	return slice, toAppend, nil
}


func TotalDistance(numbers1, numbers2 []int) int {


	slices.Sort(numbers1)
	slices.Sort(numbers2)

	sum := 0
	for i := 0 ; i < len(numbers1) ; i++ {
		sum += int(math.Abs(float64(numbers1[i] - numbers2[i])))
	}

	return sum
}